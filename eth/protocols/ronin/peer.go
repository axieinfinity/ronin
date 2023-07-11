package ronin

import (
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
)

const (
	voteChannelSize = 50
	batchInterval   = 100 * time.Millisecond
)

// Peer is a collection of relevant information we have about a `ronin` peer.
type Peer struct {
	id string // Unique ID for the peer, cached

	*p2p.Peer                          // The embedded P2P package peer
	rw        p2p.MsgReadWriter        // Input/output streams for snap
	version   uint                     // Protocol version negotiated
	term      chan struct{}            // Terminate the batch vote loop
	voteCh    chan *types.VoteEnvelope // Put vote into pool for batching

	logger log.Logger // Contextual logger with the peer id injected
}

// NewPeer create a wrapper for a network connection and negotiated  protocol
// version.
func NewPeer(version uint, p *p2p.Peer, rw p2p.MsgReadWriter) *Peer {
	id := p.ID().String()
	peer := &Peer{
		id:      id,
		Peer:    p,
		rw:      rw,
		version: version,
		voteCh:  make(chan *types.VoteEnvelope, voteChannelSize),
		term:    make(chan struct{}),
		logger:  log.New("peer", id[:8]),
	}
	go peer.batchVote()

	return peer
}

// Close terminates the vote batch goroutine.
func (p *Peer) Close() {
	close(p.term)
}

// ID retrieves the peer's unique identifier.
func (p *Peer) ID() string {
	return p.id
}

// Version retrieves the peer's negoatiated `ronin` protocol version.
func (p *Peer) Version() uint {
	return p.version
}

// Log overrides the P2P logget with the higher level one containing only the id.
func (p *Peer) Log() log.Logger {
	return p.logger
}

// sendNewVote sends votes to the peer.
func (p *Peer) sendNewVote(votes []*types.VoteEnvelope) error {
	var rawVote []*types.RawVoteEnvelope
	for _, vote := range votes {
		rawVote = append(rawVote, vote.Raw())
	}
	return p2p.Send(p.rw, NewVoteMsg, NewVotePacket{
		Vote: rawVote,
	})
}

// AsyncSendNewVote puts the vote into the batch vote goroutine.
func (p *Peer) AsyncSendNewVote(vote *types.VoteEnvelope) {
	p.voteCh <- vote
}

// batchVote batches multiple votes and sends to the peer.
func (p *Peer) batchVote() {
	var pendingVote []*types.VoteEnvelope
	ticker := time.NewTicker(batchInterval)

	for {
		select {
		case vote := <-p.voteCh:
			pendingVote = append(pendingVote, vote)
		case <-ticker.C:
			if len(pendingVote) > 0 {
				if err := p.sendNewVote(pendingVote); err != nil {
					p.Log().Debug("Failed to send vote", "err", err)
					return
				}
				pendingVote = nil
			}
		case <-p.term:
			ticker.Stop()
			return
		}
	}
}
