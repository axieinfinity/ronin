package snap

import "github.com/ethereum/go-ethereum/metrics"

var (
	ingressRegistrationErrorName = "eth/protocols/snap/ingress/registration/error"
	egressRegistrationErrorName  = "eth/protocols/snap/egress/registration/error"

	IngressRegistrationErrorMeter = metrics.NewRegisteredMeter(ingressRegistrationErrorName, nil)
	EgressRegistrationErrorMeter  = metrics.NewRegisteredMeter(egressRegistrationErrorName, nil)
)
