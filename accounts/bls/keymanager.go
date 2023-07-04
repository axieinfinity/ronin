package bls

import (
	"context"
	"encoding/json"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/bls"
	"github.com/ethereum/go-ethereum/crypto/bls/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	keystorev4 "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"
	"strings"
	"sync"
)

// IncorrectPasswordErrMsg defines a common error string representing an EIP-2335
// keystore password was incorrect.
const IncorrectPasswordErrMsg = "invalid checksum"

type SignRequest struct {
	PublicKey       []byte `json:"public_key,omitempty"`
	SigningRoot     []byte `json:"signing_root,omitempty"`
	SignatureDomain []byte `json:"signature_domain,omitempty"`
}

type KeyManager struct {
	lock sync.RWMutex

	pubKeys [][params.BLSPubkeyLength]byte
	secKeys map[[params.BLSPubkeyLength]byte]common.SecretKey

	wallet        *Wallet
	accountsStore *accountStore
}

// NewKeyManager instantiates a new local keymanager .
func NewKeyManager(ctx context.Context, wallet *Wallet) (*KeyManager, error) {
	k := &KeyManager{
		wallet:        wallet,
		accountsStore: &accountStore{},
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

	store := &accountStore{}
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
