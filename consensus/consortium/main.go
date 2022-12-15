package consortium

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	consortiumCommon "github.com/ethereum/go-ethereum/consensus/consortium/common"
	v1 "github.com/ethereum/go-ethereum/consensus/consortium/v1"
	v2 "github.com/ethereum/go-ethereum/consensus/consortium/v2"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

// Consortium is a proxy that decides the consensus version will be called
// based on params.ConsortiumV2Block
type Consortium struct {
	chainConfig *params.ChainConfig
	v1          *v1.Consortium
	v2          *v2.Consortium
}

// New creates a Consortium proxy that decides what Consortium version will be called
func New(chainConfig *params.ChainConfig, db ethdb.Database, ee *ethapi.PublicBlockChainAPI, genesisHash common.Hash) *Consortium {
	// Set any missing consensus parameters to their defaults
	consortiumV1 := v1.New(chainConfig, db, ee)
	consortiumV2 := v2.New(chainConfig, db, ee, genesisHash, consortiumV1)

	return &Consortium{
		chainConfig: chainConfig,
		v1:          consortiumV1,
		v2:          consortiumV2,
	}
}

// Author implements consensus.Engine, returning the coinbase directly
func (c *Consortium) Author(header *types.Header) (common.Address, error) {
	if c.chainConfig.IsConsortiumV2(header.Number) {
		return c.v2.Author(header)
	}

	return c.v1.Author(header)
}

// VerifyHeader checks whether a header conforms to the consensus rules.
func (c *Consortium) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header, seal bool) error {
	if c.chainConfig.IsConsortiumV2(header.Number) {
		return c.v2.VerifyHeader(chain, header, seal)
	}

	return c.v1.VerifyHeader(chain, header, seal)
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers. The
// method returns a quit channel to abort the operations and a results channel to
// retrieve the async verifications (the order is that of the input slice).
func (c *Consortium) VerifyHeaders(chain consensus.ChainHeaderReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	go func() {
		for i, header := range headers {
			var err error
			if c.chainConfig.IsConsortiumV2(header.Number) {
				err = c.v2.VerifyHeaderAndParents(chain, header, headers[:i])
			} else {
				err = c.v1.VerifyHeaderAndParents(chain, header, headers[:i])
			}

			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()

	return abort, results
}

// VerifyUncles implements consensus.Engine, always returning an error for any
// uncles as this consensus mechanism doesn't permit uncles.
func (c *Consortium) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	if c.chainConfig.IsConsortiumV2(block.Header().Number) {
		return c.v2.VerifyUncles(chain, block)
	}

	return c.v1.VerifyUncles(chain, block)
}

// Prepare implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (c *Consortium) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
	if c.chainConfig.IsConsortiumV2(header.Number) {
		return c.v2.Prepare(chain, header)
	}

	return c.v1.Prepare(chain, header)
}

// Finalize implements consensus.Engine as a proxy
func (c *Consortium) Finalize(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs *[]*types.Transaction,
	uncles []*types.Header, receipts *[]*types.Receipt, systemTxs *[]*types.Transaction, internalTxs *[]*types.InternalTransaction, usedGas *uint64) error {
	if c.chainConfig.IsConsortiumV2(header.Number) {
		return c.v2.Finalize(chain, header, state, txs, uncles, receipts, systemTxs, internalTxs, usedGas)
	}

	return c.v1.Finalize(chain, header, state, txs, uncles, receipts, systemTxs, internalTxs, usedGas)
}

// FinalizeAndAssemble implements consensus.Engine as a proxy
func (c *Consortium) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB,
	txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, []*types.Receipt, error) {
	if c.chainConfig.IsConsortiumV2(header.Number) {
		return c.v2.FinalizeAndAssemble(chain, header, state, txs, uncles, receipts)
	}

	return c.v1.FinalizeAndAssemble(chain, header, state, txs, uncles, receipts)
}

// Seal implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (c *Consortium) Seal(chain consensus.ChainHeaderReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	if c.chainConfig.IsConsortiumV2(block.Header().Number) {
		return c.v2.Seal(chain, block, results, stop)
	}

	return c.v1.Seal(chain, block, results, stop)
}

// SealHash returns the hash of a block prior to it being sealed.
func (c *Consortium) SealHash(header *types.Header) common.Hash {
	if c.chainConfig.IsConsortiumV2(header.Number) {
		return c.v2.SealHash(header)
	}

	return c.v1.SealHash(header)
}

// Close implements consensus.Engine. It's a noop for Consortium as there are no background threads.
func (c *Consortium) Close() error {
	return nil
}

// APIs doesn't need to check whether the current block is v1 or v2 because in proxy/server.go create empty
// params.ChainConfig{} so we can't use it.
func (c *Consortium) APIs(chain consensus.ChainHeaderReader) []rpc.API {
	var apis []rpc.API
	apisV1 := c.v1.APIs(chain)
	apisV2 := c.v2.APIs(chain)
	apis = append(apis, apisV1...)
	apis = append(apis, apisV2...)

	return apis
}

// CalcDifficulty is the difficulty adjustment algorithm
func (c *Consortium) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	if c.chainConfig.IsConsortiumV2(parent.Number) {
		return c.v2.CalcDifficulty(chain, time, parent)
	}

	return c.v1.CalcDifficulty(chain, time, parent)
}

// Authorize injects a private key into the consensus engine to mint new blocks with
func (c *Consortium) Authorize(signer common.Address, signFn consortiumCommon.SignerFn, signTxFn consortiumCommon.SignerTxFn) {
	c.v1.Authorize(signer, signFn, signTxFn)
	c.v2.Authorize(signer, signFn, signTxFn)
}

// SetGetSCValidatorsFn backward compatible for consortium v1
func (c *Consortium) SetGetSCValidatorsFn(fn func() ([]common.Address, error)) {
	c.v1.SetGetSCValidatorsFn(fn)
}

// SetGetFenixValidators backward compatible for consortium v1
func (c *Consortium) SetGetFenixValidators(fn func() ([]common.Address, error)) {
	c.v1.SetGetFenixValidators(fn)
}

// IsSystemTransaction implements consensus.PoSA. It is only available on v2 since v1 doesn't have system contract
func (c *Consortium) IsSystemTransaction(tx *types.Transaction, header *types.Header) (bool, error) {
	return c.v2.IsSystemTransaction(tx, header)
}

// IsSystemContract implements consensus.PoSA. It is only available on v2 since v1 doesn't have system contract
func (c *Consortium) IsSystemContract(to *common.Address) bool {
	return c.v2.IsSystemContract(to)
}
