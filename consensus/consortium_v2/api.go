package consortium_v2

import "github.com/ethereum/go-ethereum/consensus"

// API is a user facing RPC API to allow controlling the signer and voting
// mechanisms of the delegated proof-of-stake scheme.
type API struct {
	chain      consensus.ChainHeaderReader
	consortium *ConsortiumV2
}
