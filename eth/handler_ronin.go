package eth

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/protocols/ronin"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

type roninHandler handler

func (r *roninHandler) RunPeer(peer *ronin.Peer, hand ronin.Handler) error {
	return (*handler)(r).runRoninExtension(peer, hand)
}

func (r *roninHandler) PeerInfo(id enode.ID) interface{} {
	ethPeer := r.peers.peer(id.String())
	if ethPeer.roninExt != nil {
		return ethPeer.roninExt.Version()
	} else {
		return nil
	}
}

func (r *roninHandler) Handle(peer *ronin.Peer, packet ronin.Packet) error {
	switch packet.Kind() {
	case ronin.NewVoteMsg:
		if r.votePool != nil {
			votePacket := packet.(*ronin.NewVotePacket)
			for _, rawVote := range votePacket.Vote {
				vote := &types.VoteEnvelope{
					RawVoteEnvelope: *rawVote,
				}
				r.votePool.PutVote(peer.ID(), vote)
			}
		} else {
			peer.Log().Debug("Local node does not enable fast finality, drop new vote msg")
		}
	}
	return nil
}
