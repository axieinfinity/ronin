package bls

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

// AccountsKeystoreFileName exposes the name of the keystore file.
const AccountsKeystoreFileName = "all-accounts.keystore.json"

type Wallet struct {
	walletDir      string
	walletPassword string
}

func New(walletDir, walletPassword string) *Wallet {
	return &Wallet{walletDir: walletDir, walletPassword: walletPassword}
}

// SaveWallet persists the wallet's directories to disk.
func (w *Wallet) SaveWallet() error {
	if err := os.MkdirAll(w.walletDir, 0700); err != nil {
		return errors.Wrap(err, "could not create wallet directory")
	}
	return nil
}

func (w *Wallet) ReadFile(ctx context.Context, filename string) ([]byte, error) {
	existDir, err := hasDir(w.walletDir)
	if err != nil {
		return nil, err
	}
	if !existDir {
		if err = w.SaveWallet(); err != nil {
			return nil, err
		}
	}
	fullPath := filepath.Join(w.walletDir, filename)
	matches, err := filepath.Glob(fullPath)
	if err != nil {
		return []byte{}, errors.Wrap(err, "could not find file")
	}
	if len(matches) == 0 {
		return []byte{}, fmt.Errorf("no files found in path: %s", fullPath)
	}
	rawData, err := os.ReadFile(matches[0])
	if err != nil {
		return nil, errors.Wrapf(err, "could not read path: %s", fullPath)
	}
	return rawData, nil
}

func (w *Wallet) WriteFile(ctx context.Context, filename string, data []byte) error {
	existDir, err := hasDir(w.walletDir)
	if err != nil {
		return err
	}
	if !existDir {
		if err = w.SaveWallet(); err != nil {
			return err
		}
	}
	fullPath := filepath.Join(w.walletDir, filename)
	if err := os.WriteFile(fullPath, data, 0700); err != nil {
		return errors.Wrapf(err, "could not write %s", fullPath)
	}
	log.Debug("Wrote new file at path", "path", fullPath, "filename", filename)
	return nil
}

func hasDir(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if info == nil {
		return false, err
	}
	return info.IsDir(), err
}
