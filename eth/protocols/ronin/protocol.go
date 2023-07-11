package ronin

import (
	"errors"

	"github.com/ethereum/go-ethereum/core/types"
)

// Constants to match up protocol versions and messages
const (
	Ronin1 = 1
)

// ProtocolName is the official short name of the `ronin` protocol used during
// devp2p capability negotiation.
const ProtocolName = "ronin"

// ProtocolVersions are the supported versions of the `ronin` protocol
var ProtocolVersions = []uint{Ronin1}

// protocolLengths are the number of implemented message corresponding to
// different protocol versions.
var protocolLengths = map[uint]uint64{Ronin1: 1}

// maxMessageSize is the maximum cap on the size of a protocol message.
const maxMessageSize = 10 * 1024 * 1024

const (
	NewVoteMsg = 0x00
)

var (
	errMsgTooLarge    = errors.New("message too long")
	errDecode         = errors.New("invalid message")
	errInvalidMsgCode = errors.New("invalid message code")
)

// Packet represents a p2p message in the `ronin` protocol.
type Packet interface {
	Name() string // Name returns a string corresponding to the message type.
	Kind() byte   // Kind returns the message type.
}

type NewVotePacket struct {
	Vote []*types.RawVoteEnvelope
}

func (*NewVotePacket) Name() string { return "NewVote" }
func (*NewVotePacket) Kind() byte   { return NewVoteMsg }
