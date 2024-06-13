package eth

import (
	"github.com/ethereum/go-ethereum/metrics"
)

// meters stores ingress and egress handshake meters.
var meters bidirectionalMeters

// bidirectionalMeters stores ingress and egress handshake meters.

type bidirectionalMeters struct {
	ingress *hsMeters
	egress  *hsMeters
}

// get returns the corresponding meter depending if ingress or egress is desired
func (h *bidirectionalMeters) get(ingress bool) *hsMeters {
	if ingress {
		return h.ingress
	}
	return h.egress
}

// hsMeters is a collection of meters which track metrics related to the eth subprotocol handshake.
type hsMeters struct {
	// peerError measures the number of errorrs related to incorrect peer behaviour, such as invalid message code, size, encoding, etc.
	peerError metrics.Meter
	// timeoutError measures the number of timeouts.
	timeoutError metrics.Meter
	// networkIDMismatch measures the number of network ID mismatches.
	networkIDMismatch metrics.Meter
	// protocolVersionMismatch measures the number of differing protocol versions.
	protocolVersionMismatch metrics.Meter
	// genesisMismatch measures the number of differing genesies.
	genesisMismatch metrics.Meter

	// forkidRejected measures the number of rejected fork IDs.
	forkidRejected metrics.Meter
}

// newHandshakeMeters registers and returns handshake meters for the given base.
func newHandshakeMeters(base string) *hsMeters {
	return &hsMeters{
		peerError:               metrics.NewRegisteredMeter(base+"error/peer", nil),
		timeoutError:            metrics.NewRegisteredMeter(base+"error/timeout", nil),
		networkIDMismatch:       metrics.NewRegisteredMeter(base+"error/network", nil),
		protocolVersionMismatch: metrics.NewRegisteredMeter(base+"error/version", nil),
		genesisMismatch:         metrics.NewRegisteredMeter(base+"error/genesis", nil),
		forkidRejected:          metrics.NewRegisteredMeter(base+"error/forkid", nil),
	}
}

func init() {
	// Init meters for eth handshakeMeters.
	meters = bidirectionalMeters{
		ingress: newHandshakeMeters("eth/protocols/eth/ingress/handshake/"),
		egress:  newHandshakeMeters("eth/protocols/eth/egress/handshake/"),
	}
}
