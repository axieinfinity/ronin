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
	"time"
)

type Consortium struct {
	chainConfig *params.ChainConfig
	v1          *v1.Consortium
	v2          *v2.Consortium
}

// New creates a Consortium proof-of-stake consensus engine with the initial
// signers set to the ones provided by the user.
func New(chainConfig *params.ChainConfig, db ethdb.Database, ee *ethapi.PublicBlockChainAPI, genesisHash common.Hash) *Consortium {
	// Set any missing consensus parameters to their defaults
	consortiumV1 := v1.New(chainConfig.Consortium, db)
	consortiumV2 := v2.New(chainConfig, db, ee, genesisHash)

	return &Consortium{
		chainConfig: chainConfig,
		v1:          consortiumV1,
		v2:          consortiumV2,
	}
}

// Author since v1 and v2 are implemented the same logic, so we don't need to check whether the current block is version 1
// or version 2
func (c *Consortium) Author(header *types.Header) (common.Address, error) {
	return c.v1.Author(header)
}

func (c *Consortium) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header, seal bool) error {
	if c.chainConfig.IsConsortiumV2(header.Number) {
		return c.v2.VerifyHeader(chain, header, seal)
	}

	return c.v1.VerifyHeader(chain, header, seal)
}

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

func (c *Consortium) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	if c.chainConfig.IsConsortiumV2(block.Header().Number) {
		return c.v2.VerifyUncles(chain, block)
	}

	return c.v1.VerifyUncles(chain, block)
}

func (c *Consortium) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
	if c.chainConfig.IsConsortiumV2(header.Number) {
		return c.v2.Prepare(chain, header)
	}

	return c.v1.Prepare(chain, header)
}

func (c *Consortium) Finalize(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs *[]*types.Transaction,
	uncles []*types.Header, receipts *[]*types.Receipt, systemTxs *[]*types.Transaction, usedGas *uint64) error {
	if c.chainConfig.IsConsortiumV2(header.Number) {
		return c.v2.Finalize(chain, header, state, txs, uncles, receipts, systemTxs, usedGas)
	}

	return c.v1.Finalize(chain, header, state, txs, uncles, receipts, systemTxs, usedGas)
}

func (c *Consortium) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB,
	txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	if c.chainConfig.IsConsortiumV2(header.Number) {
		return c.v2.FinalizeAndAssemble(chain, header, state, txs, uncles, receipts)
	}

	return c.v1.FinalizeAndAssemble(chain, header, state, txs, uncles, receipts)
}

func (c *Consortium) Delay(chain consensus.ChainReader, header *types.Header) *time.Duration {
	return nil
}

func (c *Consortium) Seal(chain consensus.ChainHeaderReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	if c.chainConfig.IsConsortiumV2(block.Header().Number) {
		return c.v2.Seal(chain, block, results, stop)
	}

	return c.v1.Seal(chain, block, results, stop)
}

func (c *Consortium) SealHash(header *types.Header) common.Hash {
	if c.chainConfig.IsConsortiumV2(header.Number) {
		return c.v2.SealHash(header)
	}

	return c.v1.SealHash(header)
}

// Close since v1 and v2 are implemented the same logic, so we don't need to check whether the current block is version 1
// or version 2
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

func (c *Consortium) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	if c.chainConfig.IsConsortiumV2(parent.Number) {
		return c.v2.CalcDifficulty(chain, time, parent)
	}

	return c.v1.CalcDifficulty(chain, time, parent)
}

// Authorize backward compatible for consortium v1
func (c *Consortium) Authorize(signer common.Address, signFn consortiumCommon.SignerFn) {
	c.v1.Authorize(signer, signFn)
}

// SetGetSCValidatorsFn backward compatible for consortium v1
func (c *Consortium) SetGetSCValidatorsFn(fn func() ([]common.Address, error)) {
	c.v1.SetGetFenixValidators(fn)
}

// SetGetFenixValidators backward compatible for consortium v1
func (c *Consortium) SetGetFenixValidators(fn func() ([]common.Address, error)) {
	c.v1.SetGetFenixValidators(fn)
}
