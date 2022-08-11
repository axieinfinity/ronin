package consortium

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	v1 "github.com/ethereum/go-ethereum/consensus/consortium/v1"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

type Consortium struct {
	v1 *v1.Consortium
}

// New creates a Consortium proof-of-authority consensus engine with the initial
// signers set to the ones provided by the user.
func New(config *params.ConsortiumConfig, db ethdb.Database) *Consortium {
	// Set any missing consensus parameters to their defaults
	consortiumV1 := v1.New(config, db)

	return &Consortium{
		v1: consortiumV1,
	}
}

func (c *Consortium) Author(header *types.Header) (common.Address, error) {
	return c.v1.Author(header)
}

func (c *Consortium) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header, seal bool) error {
	return c.v1.VerifyHeader(chain, header, seal)
}

func (c *Consortium) VerifyHeaders(chain consensus.ChainHeaderReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	var headersV1 []*types.Header
	var headersV2 []*types.Header

	for _, header := range headers {
		if chain.Config().IsConsortiumV2(header.Number) {
			headersV2 = append(headersV2, header)
		} else {
			headersV1 = append(headersV1, header)
		}
	}

	// TODO: handle headers v2 is WIP

	return c.v1.VerifyHeaders(chain, headersV1, seals)
}

func (c *Consortium) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	return c.v1.VerifyUncles(chain, block)
}

func (c *Consortium) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
	return c.v1.Prepare(chain, header)
}

func (c *Consortium) Finalize(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	c.v1.Finalize(chain, header, state, txs, uncles)
}

func (c *Consortium) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	return c.v1.FinalizeAndAssemble(chain, header, state, txs, uncles, receipts)
}

func (c *Consortium) Seal(chain consensus.ChainHeaderReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	return c.v1.Seal(chain, block, results, stop)
}

func (c *Consortium) SealHash(header *types.Header) common.Hash {
	return c.v1.SealHash(header)
}

func (c *Consortium) Close() error {
	return nil
}

func (c *Consortium) APIs(chain consensus.ChainHeaderReader) []rpc.API {
	return c.v1.APIs(chain)
}

func (c *Consortium) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	return c.v1.CalcDifficulty(chain, time, parent)
}

func (c *Consortium) SetGetSCValidatorsFn(fn func() ([]common.Address, error)) {
	c.v1.SetGetFenixValidators(fn)
}

func (c *Consortium) SetGetFenixValidators(fn func() ([]common.Address, error)) {
	c.v1.SetGetFenixValidators(fn)
}
