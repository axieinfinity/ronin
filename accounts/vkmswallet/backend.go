package vkmswallet

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	vkms "github.com/ethereum/go-ethereum/accounts/vkmswallet/message"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum/go-ethereum/rlp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

//go:generate protoc --go_out=message --go-grpc_out=message message/message.proto

const SuccessCode = 1
const InternalErrorCode = 2

type WalletConfig struct {
	VKMSAddress        string
	KeyUsageTokenPath  string
	SourceAddress      string
	SslCertificatePath string
}

type Backend struct {
	wallets []accounts.Wallet
}

func NewBackend(configs []*WalletConfig) (*Backend, error) {
	wallets := make([]accounts.Wallet, 0, len(configs))
	for _, config := range configs {
		wallet, err := NewWallet(config)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet)
	}
	return &Backend{wallets: wallets}, nil
}

func (b *Backend) Wallets() []accounts.Wallet {
	return b.wallets
}

func (b *Backend) Subscribe(sink chan<- accounts.WalletEvent) event.Subscription {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		<-quit
		return nil
	})
}

type Wallet struct {
	account       accounts.Account
	connection    *grpc.ClientConn
	keyUsageToken []byte
	config        *WalletConfig
	status        string
}

func NewWallet(config *WalletConfig) (*Wallet, error) {
	// parse source address
	sourceAddr, err := net.ResolveTCPAddr("tcp", config.SourceAddress)
	if err != nil {
		return nil, err
	}
	var dialer = net.Dialer{
		LocalAddr: sourceAddr,
	}

	// load VKMS certificate
	certString, err := ioutil.ReadFile(config.SslCertificatePath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM([]byte(certString))

	// prepare a GRPC client
	conn, err := grpc.Dial(config.VKMSAddress,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{RootCAs: certPool})),
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return dialer.DialContext(ctx, "tcp", s)
		}),
	)
	if err != nil {
		return nil, err
	}
	client := vkms.NewUserClient(conn)

	// load key usage token
	token, err := ioutil.ReadFile(config.KeyUsageTokenPath)
	if err != nil {
		return nil, err
	}

	// get a testing signature
	resp, err := client.Sign(
		metadata.AppendToOutgoingContext(context.Background(), "vkms_data_type", "non-ether"),
		&vkms.SignRequest{
			KeyUsageToken: token,
			Data:          []byte{},
		})
	if err != nil {
		return nil, err
	}
	if resp.Code != SuccessCode {
		return nil, fmt.Errorf("internal server error")
	}

	// recover public key and create corresponding account
	publicKey, err := crypto.SigToPub(crypto.Keccak256([]byte{}), resp.Signature)
	if err != nil {
		utils.Fatalf("can not recover public key")
	}
	account := accounts.Account{
		Address: crypto.PubkeyToAddress(*publicKey),
	}
	log.Info("KMS wallet account", "address", account.Address)

	// everything is fine
	wallet := &Wallet{
		account:       account,
		connection:    conn,
		keyUsageToken: token,
		config:        config,
		status:        "OK",
	}
	return wallet, nil
}

func (w *Wallet) URL() accounts.URL {
	return accounts.URL{}
}

func (w *Wallet) Status() (string, error) {
	return w.status, nil
}

func (w *Wallet) Open(passphrase string) error {
	return nil
}

func (w *Wallet) Close() error {
	_ = w.connection.Close()
	return nil
}

func (w *Wallet) Accounts() []accounts.Account {

	return []accounts.Account{w.account}
}

func (w *Wallet) Contains(account accounts.Account) bool {
	return w.account.Address == account.Address
}

func (w *Wallet) Derive(path accounts.DerivationPath, pin bool) (accounts.Account, error) {
	return accounts.Account{}, accounts.ErrNotSupported
}

func (w *Wallet) SelfDerive(bases []accounts.DerivationPath, chain ethereum.ChainStateReader) {
	// nothing to do
	log.Warn("KMS wallet does not support SelfDerive method")
}

func (w *Wallet) SignHash(account accounts.Account, hash []byte) ([]byte, error) {
	// VKMS never signs input hash. It always requires the original data, usually transaction data, to check if the
	// transaction meets application-level policies.
	return nil, accounts.ErrNotSupported
}

func (w *Wallet) SignData(account accounts.Account, mimeType string, data []byte) ([]byte, error) {
	if !w.Contains(account) {
		return nil, accounts.ErrUnknownAccount
	}

	client := vkms.NewUserClient(w.connection)
	resp, err := client.Sign(
		metadata.AppendToOutgoingContext(context.Background(), "vkms_data_type", "non-ether"),
		&vkms.SignRequest{
			KeyUsageToken: w.keyUsageToken,
			Data:          data,
		})
	if err != nil {
		return nil, err
	}
	if resp.Code != SuccessCode {
		return nil, fmt.Errorf("access denied")
	}
	return resp.Signature, nil
}

func (w *Wallet) SignText(account accounts.Account, text []byte) ([]byte, error) {
	wrapped := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(text), string(text))
	return w.SignData(account, "", []byte(wrapped))
}

func (w *Wallet) SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	// try to serialize the transaction based on `chainID` and transaction type
	data := new(bytes.Buffer)
	var txData interface{}

	if chainID == nil { // homestead
		txData = []interface{}{
			tx.Nonce(),
			tx.GasPrice(),
			tx.Gas(),
			tx.To(),
			tx.Value(),
			tx.Data(),
		}
	} else { // london
		if tx.Type() == types.LegacyTxType {
			txData = []interface{}{
				tx.Nonce(),
				tx.GasPrice(),
				tx.Gas(),
				tx.To(),
				tx.Value(),
				tx.Data(),
				chainID, uint(0), uint(0),
			}
		} else if tx.Type() == types.AccessListTxType {
			data.Write([]byte{tx.Type()})
			txData = []interface{}{
				chainID,
				tx.Nonce(),
				tx.GasPrice(),
				tx.Gas(),
				tx.To(),
				tx.Value(),
				tx.Data(),
				tx.AccessList(),
			}
		} else { // types.DynamicFeeTxType
			data.Write([]byte{tx.Type()})
			txData = []interface{}{
				chainID,
				tx.Nonce(),
				tx.GasTipCap(),
				tx.GasFeeCap(),
				tx.Gas(),
				tx.To(),
				tx.Value(),
				tx.Data(),
				tx.AccessList(),
			}
		}
	}
	if err := rlp.Encode(data, txData); err != nil {
		return nil, err
	}
	sig, err := w.SignData(account, "", data.Bytes())
	if err != nil {
		return nil, err
	}
	signer := types.LatestSignerForChainID(chainID)
	return tx.WithSignature(signer, sig)
}

func (w *Wallet) SignDataWithPassphrase(account accounts.Account, passphrase, mimeType string, data []byte) ([]byte, error) {
	return w.SignData(account, mimeType, data)
}

func (w *Wallet) SignTextWithPassphrase(account accounts.Account, passphrase string, text []byte) ([]byte, error) {
	return w.SignText(account, text)
}

func (w *Wallet) SignTxWithPassphrase(account accounts.Account, passphrase string, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	return w.SignTx(account, tx, chainID)
}
