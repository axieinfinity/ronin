package discover

import (
	"fmt"

	"github.com/ethereum/go-ethereum/metrics"
)

const (
	moduleName = "discover"

	// ingressMeterName is the prefix of the per-packet inbound metrics
	ingressMeterName = moduleName + "/ingress"

	// egressMeterName is the prefix of the per-packet outbound metrics
	egressMeterName = moduleName + "/egress"
)

var (
	bucketGauge         []metrics.Gauge
	ingressTrafficMeter = metrics.NewRegisteredMeter(ingressMeterName, nil)
	egressTrafficMeter  = metrics.NewRegisteredMeter(egressMeterName, nil)
)

func init() {
	for i := 0; i < nBuckets; i++ {
		bucketGauge = append(bucketGauge, metrics.NewRegisteredGauge(fmt.Sprintf("%s/bucket/%d/count", moduleName, i), nil))
	}
}

type meteredUdpConn struct {
	UDPConn
}

func newMeteredConn(conn UDPConn) UDPConn {
	// Short circuit if metrics are disabled
	if !metrics.Enabled {
		return conn
	}
	return &meteredUdpConn{UDPConn: conn}
}

// Read delegates a network read to the underlying connection, bumpding the
