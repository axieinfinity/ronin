package bls

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	keystorev4 "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"
)

// Keystore json file representation as a Go struct.
type Keystore struct {
	Crypto  map[string]interface{} `json:"crypto"`
	ID      string                 `json:"uuid"`
	Pubkey  string                 `json:"pubkey"`
	Version uint                   `json:"version"`
	Name    string                 `json:"name"`
	Path    string                 `json:"path"`
}

// Defines a struct containing 1-to-1 corresponding
// private keys and public keys for Ethereum validators.
type AccountStore struct {
	PrivateKeys [][]byte `json:"private_keys"`
	PublicKeys  [][]byte `json:"public_keys"`
}

// Copy creates a deep copy of accountStore
func (a *AccountStore) Copy() *AccountStore {
	storeCopy := &AccountStore{}
	storeCopy.PrivateKeys = common.Copy2dBytes(a.PrivateKeys)
	storeCopy.PublicKeys = common.Copy2dBytes(a.PublicKeys)
	return storeCopy
}

// AccountsKeystoreRepresentation defines an internal Prysm representation
// of validator accounts, encrypted according to the EIP-2334 standard.
type AccountsKeystoreRepresentation struct {
	Crypto  map[string]interface{} `json:"crypto"`
	ID      string                 `json:"uuid"`
	Version uint                   `json:"version"`
	Name    string                 `json:"name"`
}

// CreateAccountsKeystoreRepresentation is a pure function that takes an accountStore and wallet password and returns the encrypted formatted json version for local writing.
func CreateAccountsKeystoreRepresentation(
	store *AccountStore,
	walletPW string,
) (*AccountsKeystoreRepresentation, error) {
	encryptor := keystorev4.New()
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	encodedStore, err := json.MarshalIndent(store, "", "\t")
	if err != nil {
		return nil, err
	}
	cryptoFields, err := encryptor.Encrypt(encodedStore, walletPW)
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
