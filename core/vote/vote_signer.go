package vote

import (
	"context"
	"fmt"
	wallet "github.com/ethereum/go-ethereum/accounts/bls"
	"io/ioutil"
	"time"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/bls"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
)

const (
	voteSignerTimeout = time.Second * 5
)

var votesSigningErrorCounter = metrics.NewRegisteredCounter("votesSigner/error", nil)

type VoteSigner struct {
	km     *wallet.KeyManager
	pubKey [48]byte
}

func NewVoteSigner(blsPasswordPath, blsWalletPath string) (*VoteSigner, error) {
	dirExists, err := wallet.HasDir(blsWalletPath)
	if err != nil {
		log.Error("Check BLS wallet exists", "err", err)
		return nil, err
	}
	if !dirExists {
		log.Error("BLS wallet did not exists.")
		return nil, fmt.Errorf("BLS wallet did not exists.")
	}

	walletPassword, err := ioutil.ReadFile(blsPasswordPath)
	if err != nil {
		log.Error("Read BLS wallet password", "err", err)
		return nil, err
	}
	log.Info("Read BLS wallet password successfully")

	w := wallet.New(blsWalletPath, string(walletPassword))
	km, err := wallet.NewKeyManager(context.Background(), w)
	if err != nil {
		log.Error("Initialize key manager failed", "err", err)
		return nil, err
	}
	log.Info("Initialized keymanager successfully")

	ctx, cancel := context.WithTimeout(context.Background(), voteSignerTimeout)
	defer cancel()

	pubKeys, err := km.FetchValidatingPublicKeys(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "could not fetch validating public keys")
	}

	return &VoteSigner{
		km:     km,
		pubKey: pubKeys[0],
	}, nil
}

func (signer *VoteSigner) SignVote(vote *types.VoteEnvelope) error {
	// Sign the vote, fetch the first pubKey as validator's bls public key.
	pubKey := signer.pubKey
	blsPubKey, err := bls.PublicKeyFromBytes(pubKey[:])
	if err != nil {
		return errors.Wrap(err, "convert public key from bytes to bls failed")
	}

	voteDataHash := vote.Data.Hash()

	ctx, cancel := context.WithTimeout(context.Background(), voteSignerTimeout)
	defer cancel()

	signature, err := (*signer.km).Sign(ctx, &wallet.SignRequest{
		PublicKey:   pubKey[:],
		SigningRoot: voteDataHash[:],
	})
	if err != nil {
		return err
	}

	copy(vote.VoteAddress[:], blsPubKey.Marshal()[:])
	copy(vote.Signature[:], signature.Marshal()[:])
	return nil
}
