// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Contains the meters and timers used by the networking layer.

package p2p

import (
	"errors"
	"fmt"
	"net"

	"github.com/ethereum/go-ethereum/metrics"
)

const (
	// ingressMeterName is the prefix of the per-packet inbound metrics.
	ingressMeterName = "p2p/ingress"

	// egressMeterName is the prefix of the per-packet outbound metrics.
	egressMeterName = "p2p/egress"

	// HandleHistName is the prefix of the per-packet serving time histograms.
	HandleHistName = "p2p/handle"

	// dialErrorMeterName is the prefix of the dial error metrics.
	dialErrorMeterName = "p2p/dials/error"
)

var (
	activePeerGauge     metrics.Gauge = metrics.NilGauge{}
	ingressTrafficMeter               = metrics.NewRegisteredMeter(ingressMeterName, nil)
	egressTrafficMeter                = metrics.NewRegisteredMeter(egressMeterName, nil)

	// general ingress/egress connection meters
	serveMeter          metrics.Meter = metrics.NilMeter{}
	serveSuccessMeter   metrics.Meter = metrics.NilMeter{}
	dialMeter           metrics.Meter = metrics.NilMeter{}
	dialSuccessMeter    metrics.Meter = metrics.NilMeter{}
	dialConnectionError metrics.Meter = metrics.NilMeter{}

	// Handshake error meters
	dialTooManyPeers        = metrics.NewRegisteredMeter(fmt.Sprintf("%s/saturated", dialErrorMeterName), nil)
	dialAlreadyConnected    = metrics.NewRegisteredMeter(fmt.Sprintf("%s/known", dialErrorMeterName), nil)
	dialSelf                = metrics.NewRegisteredMeter(fmt.Sprintf("%s/self", dialErrorMeterName), nil)
	dialUselessPeer         = metrics.NewRegisteredMeter(fmt.Sprintf("%s/useless", dialErrorMeterName), nil)
	dialUnexpectedIdentity  = metrics.NewRegisteredMeter(fmt.Sprintf("%s/id/unexpected", dialErrorMeterName), nil)
	dialEncHandshakeError   = metrics.NewRegisteredMeter(fmt.Sprintf("%s/rlpx/enc", dialErrorMeterName), nil)
	dialProtoHandshakeError = metrics.NewRegisteredMeter(fmt.Sprintf("%s/rlpx/proto", dialErrorMeterName), nil)
)

func init() {
	if !metrics.Enabled {
		return
	}

	activePeerGauge = metrics.NewRegisteredGauge("p2p/peers", nil)
	serveMeter = metrics.NewRegisteredMeter("p2p/serves", nil)
	serveSuccessMeter = metrics.NewRegisteredMeter("p2p/serves/success", nil)
	dialMeter = metrics.NewRegisteredMeter("p2p/dials", nil)
	dialSuccessMeter = metrics.NewRegisteredMeter("p2p/dials/success", nil)
	dialConnectionError = metrics.NewRegisteredMeter(fmt.Sprintf("%s/connection", dialErrorMeterName), nil)
}

// markDialError matches error that occur while setting up a dial connection
// to the coressponding error meter.
func markDialError(err error) {
	if !metrics.Enabled {
		return
	}

	if err2 := errors.Unwrap(err); err2 != nil {
		err = err2
	}

	switch err {
	case DiscTooManyPeers:
		dialTooManyPeers.Mark(1)
	case DiscAlreadyConnected:
		dialAlreadyConnected.Mark(1)
	case DiscSelf:
		dialSelf.Mark(1)
	case DiscUselessPeer:
		dialUselessPeer.Mark(1)
	case DiscUnexpectedIdentity:
		dialUnexpectedIdentity.Mark(1)
	case errEncHandshakeError:
		dialEncHandshakeError.Mark(1)
	case errProtoHandshakeError:
		dialProtoHandshakeError.Mark(1)
	}
}

// meteredConn is a wrapper around a net.Conn that meters both the
// inbound and outbound network traffic.
type meteredConn struct {
	net.Conn
}

// newMeteredConn creates a new metered connection, bumps the ingress or egress
// connection meter and also increases the metered peer count. If the metrics
// system is disabled, function returns the original connection.
func newMeteredConn(conn net.Conn) net.Conn {
	// Short circuit if metrics are disabled
	if !metrics.Enabled {
		return conn
	}

	return &meteredConn{Conn: conn}
}

// Read delegates a network read to the underlying connection, bumping the common
// and the peer ingress traffic meters along the way.
func (c *meteredConn) Read(b []byte) (n int, err error) {
	n, err = c.Conn.Read(b)
	ingressTrafficMeter.Mark(int64(n))
	return n, err
}

// Write delegates a network write to the underlying connection, bumping the common
// and the peer egress traffic meters along the way.
func (c *meteredConn) Write(b []byte) (n int, err error) {
	n, err = c.Conn.Write(b)
	egressTrafficMeter.Mark(int64(n))
	return n, err
}
