package eth

import (
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/forkid"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/protocols/eth"
	"github.com/ethereum/go-ethereum/eth/protocols/ronin"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

// testRoninHandler is a mock event handler to listen for inbound network requests
// on the `eth` protocol and convert them into a more easily testable form.
type testRoninHandler struct {
	voteBroadcasts event.Feed
}

func (h *testRoninHandler) RunPeer(*ronin.Peer, ronin.Handler) error { panic("not used in tests") }
func (h *testRoninHandler) PeerInfo(enode.ID) interface{}            { panic("not used in tests") }

func (h *testRoninHandler) Handle(peer *ronin.Peer, packet ronin.Packet) error {
	switch packet.Kind() {
	case ronin.NewVoteMsg:
		h.voteBroadcasts.Send(packet.Name())
		return nil

	default:
		panic(fmt.Sprintf("unexpected eth packet type in tests: %T", packet))
	}
}

func TestVoteBroadcast(t *testing.T) {
	const peers = 10
	protocols := []p2p.Protocol{
		{
			Name:    eth.ProtocolName,
			Version: eth.ETH66,
		},
		{
			Name:    ronin.ProtocolName,
			Version: ronin.Ronin1,
		},
	}
	caps := []p2p.Cap{
		{
			Name:    eth.ProtocolName,
			Version: eth.ETH66,
		},
		{
			Name:    ronin.ProtocolName,
			Version: ronin.Ronin1,
		},
	}

	// Create a source eth handler
	source := newTestHandler()
	defer source.close()

	sinksEth := make([]*testEthHandler, peers)
	for i := 0; i < len(sinksEth); i++ {
		sinksEth[i] = new(testEthHandler)
	}
	sinksRonin := make([]*testRoninHandler, peers)
	for i := 0; i < len(sinksRonin); i++ {
		sinksRonin[i] = new(testRoninHandler)
	}

	// Interconnect all the sink handlers with the source handler
	var (
		genesis = source.chain.Genesis()
		td      = source.chain.GetTd(genesis.Hash(), genesis.NumberU64())
	)
	for i := range sinksEth {
		sinkEth := sinksEth[i]
		sinkRonin := sinksRonin[i]

		sourceEthPipe, sinkEthPipe := p2p.MsgPipe()
		defer sourceEthPipe.Close()
		defer sinkEthPipe.Close()

		sourceEthPeer := eth.NewPeer(
			eth.ETH66,
			p2p.NewPeerPipeWithProtocol(enode.ID{byte(i + 1)}, "", caps, sourceEthPipe, protocols),
			sourceEthPipe,
			nil,
		)
		sinkEthPeer := eth.NewPeer(
			eth.ETH66,
			p2p.NewPeerPipeWithProtocol(enode.ID{0}, "", caps, sinkEthPipe, protocols),
			sinkEthPipe,
			nil,
		)
		defer sourceEthPeer.Close()
		defer sinkEthPeer.Close()

		sourceRoninPipe, sinkRoninPipe := p2p.MsgPipe()
		defer sourceRoninPipe.Close()
		defer sinkRoninPipe.Close()

		sourceRoninPeer := ronin.NewPeer(
			ronin.Ronin1,
			p2p.NewPeerPipeWithProtocol(enode.ID{byte(i + 1)}, "", caps, sourceRoninPipe, protocols),
			sourceRoninPipe,
		)
		sinkRoninPeer := ronin.NewPeer(
			ronin.Ronin1,
			p2p.NewPeerPipeWithProtocol(enode.ID{0}, "", caps, sinkRoninPipe, protocols),
			sinkRoninPipe,
		)
		defer sourceRoninPeer.Close()
		defer sinkRoninPeer.Close()

		go source.handler.runRoninExtension(sourceRoninPeer, func(peer *ronin.Peer) error {
			return ronin.Handle((*roninHandler)(source.handler), peer)
		})
		go ronin.Handle(sinkRonin, sinkRoninPeer)

		go source.handler.runEthPeer(sourceEthPeer, func(peer *eth.Peer) error {
			return eth.Handle((*ethHandler)(source.handler), peer)
		})

		if err := sinkEthPeer.Handshake(
			1,
			td,
			genesis.Hash(),
			genesis.Hash(),
			forkid.NewIDWithChain(source.chain),
			forkid.NewFilter(source.chain),
		); err != nil {
			t.Fatalf("failed to run protocol handshake, err %s", err)
		}

		go eth.Handle(sinkEth, sinkEthPeer)
	}

	// Subscribe to all the vote sinks
	voteChs := make([]chan string, len(sinksRonin))
	for i := 0; i < len(sinksRonin); i++ {
		voteChs[i] = make(chan string, 1)
		defer close(voteChs[i])

		sub := sinksRonin[i].voteBroadcasts.Subscribe(voteChs[i])
		defer sub.Unsubscribe()
	}

	// Initiate a vote propagation across the peers
	time.Sleep(100 * time.Millisecond)
	source.handler.broadcastVote(&types.VoteEnvelope{
		RawVoteEnvelope: types.RawVoteEnvelope{
			Data: &types.VoteData{
				TargetNumber: 0,
				TargetHash:   common.Hash{},
			},
		},
	})

	// Iterate through all the sinks and ensure the correct number of the votes
	done := make(chan struct{}, peers)
	for _, ch := range voteChs {
		ch := ch
		go func() {
			<-ch
			done <- struct{}{}
		}()
	}
	var received int
	for {
		select {
		case <-done:
			received++

		case <-time.After(200 * time.Millisecond):
			if received != peers {
				t.Errorf("broadcast count mismatch: have %d, want %d", received, peers)
			}
			return
		}
	}
}
