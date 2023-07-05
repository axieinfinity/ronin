package bls

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/bls"
	"github.com/ethereum/go-ethereum/crypto/bls/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	keystorev4 "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"
	"strings"
	"sync"
)

const (
	// IncorrectPasswordErrMsg defines a common error string representing an EIP-2335
	// keystore password was incorrect.
	IncorrectPasswordErrMsg          = "invalid checksum"
	ImportedKeystoreStatus_IMPORTED  = 0
	ImportedKeystoreStatus_DUPLICATE = 1
	ImportedKeystoreStatus_ERROR     = 2
)

type ImportedKeystoreStatus struct {
	Status  int32  `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

var (
	ErrNoPasswords            = errors.New("no passwords provided for keystores")
	ErrMismatchedNumPasswords = errors.New("number of passwords does not match number of keystores")
)

type SignRequest struct {
	PublicKey   []byte `json:"public_key,omitempty"`
	SigningRoot []byte `json:"signing_root,omitempty"`
}

type KeyManager struct {
	lock sync.RWMutex

	pubKeys [][params.BLSPubkeyLength]byte
	secKeys map[[params.BLSPubkeyLength]byte]common.SecretKey

	wallet        *Wallet
	accountsStore *AccountStore
}

// NewKeyManager instantiates a new local keymanager .
func NewKeyManager(ctx context.Context, wallet *Wallet) (*KeyManager, error) {
	k := &KeyManager{
		wallet:        wallet,
		accountsStore: &AccountStore{},
	}

	if err := k.initializeAccountKeystore(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to initialize account store")
	}
	return k, nil
}

func (km *KeyManager) initializeAccountKeystore(ctx context.Context) error {
	encoded, err := km.wallet.ReadFile(ctx, AccountsKeystoreFileName)
	if err != nil && strings.Contains(err.Error(), "no files found") {
		// If there are no keys to initialize at all, just exit.
		return nil
	} else if err != nil {
		return errors.Wrapf(err, "could not read keystore file for accounts %s", AccountsKeystoreFileName)
	}
	keystoreFile := &AccountsKeystoreRepresentation{}
	if err := json.Unmarshal(encoded, keystoreFile); err != nil {
		return errors.Wrapf(err, "could not decode keystore file for accounts %s", AccountsKeystoreFileName)
	}
	// We extract the validator signing private key from the keystore
	// by utilizing the password and initialize a new BLS secret key from
	// its raw bytes.
	password := km.wallet.walletPassword
	decryptor := keystorev4.New()
	enc, err := decryptor.Decrypt(keystoreFile.Crypto, password)
	if err != nil && strings.Contains(err.Error(), IncorrectPasswordErrMsg) {
		return errors.Wrap(err, "wrong password for wallet entered")
	} else if err != nil {
		return errors.Wrap(err, "could not decrypt keystore")
	}

	store := &AccountStore{}
	if err := json.Unmarshal(enc, store); err != nil {
		return err
	}
	if len(store.PublicKeys) != len(store.PrivateKeys) {
		return errors.New("unequal number of public keys and private keys")
	}
	if len(store.PublicKeys) == 0 {
		return nil
	}
	km.accountsStore = store
	err = km.initializeKeysCachesFromKeystore()
	if err != nil {
		return errors.Wrap(err, "failed to initialize keys caches")
	}
	return err
}

// Initialize public and secret key caches that are used to speed up the functions
// FetchValidatingPublicKeys and Sign
func (km *KeyManager) initializeKeysCachesFromKeystore() error {
	km.lock.Lock()
	defer km.lock.Unlock()
	count := len(km.accountsStore.PrivateKeys)
	km.pubKeys = make([][params.BLSPubkeyLength]byte, count)
	km.secKeys = make(map[[params.BLSPubkeyLength]byte]common.SecretKey, count)
	for i, publicKey := range km.accountsStore.PublicKeys {
		publicKey48 := ethCommon.ToBytes48(publicKey)
		km.pubKeys[i] = publicKey48
		secretKey, err := bls.SecretKeyFromBytes(km.accountsStore.PrivateKeys[i])
		if err != nil {
			return errors.Wrap(err, "failed to initialize keys caches from account keystore")
		}
		km.secKeys[publicKey48] = secretKey
	}
	return nil
}

// FetchValidatingPublicKeys fetches the list of active public keys from the local account keystores.
func (km *KeyManager) FetchValidatingPublicKeys(ctx context.Context) ([][params.BLSPubkeyLength]byte, error) {
	km.lock.RLock()
	defer km.lock.RUnlock()
	keys := km.pubKeys
	result := make([][params.BLSPubkeyLength]byte, len(keys))
	copy(result, keys)
	return result, nil
}

// Sign signs a message using a validator key.
func (km *KeyManager) Sign(ctx context.Context, req *SignRequest) (bls.Signature, error) {
	publicKey := req.PublicKey
	if publicKey == nil {
		return nil, errors.New("nil public key in request")
	}
	km.lock.RLock()
	secretKey, ok := km.secKeys[ethCommon.ToBytes48(publicKey)]
	km.lock.RUnlock()
	if !ok {
		return nil, errors.New("no signing key found in keys cache")
	}
	return secretKey.Sign(req.SigningRoot), nil
}

// ImportKeystores into the local keymanager from an external source.
func (km *KeyManager) ImportKeystores(
	ctx context.Context,
	keystores []*Keystore,
	passwords []string,
) ([]*ImportedKeystoreStatus, error) {
	if len(passwords) == 0 {
		return nil, ErrNoPasswords
	}
	if len(passwords) != len(keystores) {
		return nil, ErrMismatchedNumPasswords
	}
	decryptor := keystorev4.New()
	keys := map[string]string{}
	statuses := make([]*ImportedKeystoreStatus, len(keystores))
	var err error

	for i := 0; i < len(keystores); i++ {
		var privKeyBytes []byte
		var pubKeyBytes []byte
		privKeyBytes, pubKeyBytes, _, err = km.attemptDecryptKeystore(decryptor, keystores[i], passwords[i])
		if err != nil {
			statuses[i] = &ImportedKeystoreStatus{
				Status:  ImportedKeystoreStatus_ERROR,
				Message: err.Error(),
			}
			continue
		}
		// if key exists prior to being added then output log that duplicate key was found
		if _, ok := keys[string(pubKeyBytes)]; ok {
			log.Warn(fmt.Sprintf("Duplicate key in import will be ignored: %#x", pubKeyBytes))
			statuses[i] = &ImportedKeystoreStatus{
				Status: ImportedKeystoreStatus_DUPLICATE,
			}
			continue
		}
		keys[string(pubKeyBytes)] = string(privKeyBytes)
		statuses[i] = &ImportedKeystoreStatus{
			Status: ImportedKeystoreStatus_IMPORTED,
		}
	}
	privKeys := make([][]byte, 0)
	pubKeys := make([][]byte, 0)
	for pubKey, privKey := range keys {
		pubKeys = append(pubKeys, []byte(pubKey))
		privKeys = append(privKeys, []byte(privKey))
	}

	// Write the accounts to disk into a single keystore.
	accountsKeystore, err := km.CreateAccountsKeystore(ctx, privKeys, pubKeys)
	if err != nil {
		return nil, err
	}
	encodedAccounts, err := json.MarshalIndent(accountsKeystore, "", "\t")
	if err != nil {
		return nil, err
	}
	if err := km.wallet.WriteFile(ctx, AccountsKeystoreFileName, encodedAccounts); err != nil {
		return nil, err
	}
	return statuses, nil
}

// ImportKeypairs directly into the keyManager.
func (km *KeyManager) ImportKeypairs(ctx context.Context, privKeys, pubKeys [][]byte) error {
	// Write the accounts to disk into a single keystore.
	accountsKeystore, err := km.CreateAccountsKeystore(ctx, privKeys, pubKeys)
	if err != nil {
		return errors.Wrap(err, "could not import account keypairs")
	}
	encodedAccounts, err := json.MarshalIndent(accountsKeystore, "", "\t")
	if err != nil {
		return errors.Wrap(err, "could not marshal accounts keystore into JSON")
	}
	return km.wallet.WriteFile(ctx, AccountsKeystoreFileName, encodedAccounts)
}

// CreateAccountsKeystore creates a new keystore holding the provided keys.
func (km *KeyManager) CreateAccountsKeystore(
	_ context.Context,
	privateKeys, publicKeys [][]byte,
) (*AccountsKeystoreRepresentation, error) {
	encryptor := keystorev4.New()
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	if len(privateKeys) != len(publicKeys) {
		return nil, fmt.Errorf(
			"number of private keys and public keys is not equal: %d != %d", len(privateKeys), len(publicKeys),
		)
	}
	if km.accountsStore == nil {
		km.accountsStore = &AccountStore{
			PrivateKeys: privateKeys,
			PublicKeys:  publicKeys,
		}
	} else {
		existingPubKeys := make(map[string]bool)
		existingPrivKeys := make(map[string]bool)
		for i := 0; i < len(km.accountsStore.PrivateKeys); i++ {
			existingPrivKeys[string(km.accountsStore.PrivateKeys[i])] = true
			existingPubKeys[string(km.accountsStore.PublicKeys[i])] = true
		}
		// We append to the accounts store keys only
		// if the private/secret key do not already exist, to prevent duplicates.
		for i := 0; i < len(privateKeys); i++ {
			sk := privateKeys[i]
			pk := publicKeys[i]
			_, privKeyExists := existingPrivKeys[string(sk)]
			_, pubKeyExists := existingPubKeys[string(pk)]
			if privKeyExists || pubKeyExists {
				continue
			}
			km.accountsStore.PublicKeys = append(km.accountsStore.PublicKeys, pk)
			km.accountsStore.PrivateKeys = append(km.accountsStore.PrivateKeys, sk)
		}
	}
	err = km.initializeKeysCachesFromKeystore()
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize keys caches")
	}
	encodedStore, err := json.MarshalIndent(km.accountsStore, "", "\t")
	if err != nil {
		return nil, err
	}
	cryptoFields, err := encryptor.Encrypt(encodedStore, km.wallet.walletPassword)
	if err != nil {
		return nil, errors.Wrap(err, "could not encrypt accounts")
	}
	return &AccountsKeystoreRepresentation{
		Crypto:  cryptoFields,
		ID:      id.String(),
		Version: encryptor.Version(),
		Name:    encryptor.Name(),
	}, nil
}

// Retrieves the private key and public key from an EIP-2335 keystore file
// by decrypting using a specified password. If the password fails,
// it prompts the user for the correct password until it confirms.
func (_ *KeyManager) attemptDecryptKeystore(
	enc *keystorev4.Encryptor, keystore *Keystore, password string,
) ([]byte, []byte, string, error) {
	// Attempt to decrypt the keystore with the specifies password.
	var privKeyBytes []byte
	var err error
	privKeyBytes, err = enc.Decrypt(keystore.Crypto, password)
	doesNotDecrypt := err != nil && strings.Contains(err.Error(), IncorrectPasswordErrMsg)
	if doesNotDecrypt {
		return nil, nil, "", fmt.Errorf(
			"incorrect password for key 0x%s",
			keystore.Pubkey,
		)
	}
	if err != nil && !strings.Contains(err.Error(), IncorrectPasswordErrMsg) {
		return nil, nil, "", errors.Wrap(err, "could not decrypt keystore")
	}
	var pubKeyBytes []byte
	// Attempt to use the pubkey present in the keystore itself as a field. If unavailable,
	// then utilize the public key directly from the private key.
	if keystore.Pubkey != "" {
		pubKeyBytes, err = hex.DecodeString(keystore.Pubkey)
		if err != nil {
			return nil, nil, "", errors.Wrap(err, "could not decode pubkey from keystore")
		}
	} else {
		privKey, err := bls.SecretKeyFromBytes(privKeyBytes)
		if err != nil {
			return nil, nil, "", errors.Wrap(err, "could not initialize private key from bytes")
		}
		pubKeyBytes = privKey.PublicKey().Marshal()
	}
	return privKeyBytes, pubKeyBytes, password, nil
}
