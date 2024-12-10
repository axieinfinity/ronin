package snap

import "github.com/ethereum/go-ethereum/metrics"

var (
	ingressRegistrationErrorName = "eth/protocols/snap/ingress/registration/error"
	egressRegistrationErrorName  = "eth/protocols/snap/egress/registration/error"

	IngressRegistrationErrorMeter = metrics.NewRegisteredMeter(ingressRegistrationErrorName, nil)
	EgressRegistrationErrorMeter  = metrics.NewRegisteredMeter(egressRegistrationErrorName, nil)

	// deletionGauge is the metric to track how many trie node deletions
	// are performed in total during the sync process.
	deletionGauge = metrics.NewRegisteredGauge("eth/protocols/snap/sync/delete", nil)

	// lookupGauge is the metric to track how many trie node lookups are
	// performed to determine if node needs to be deleted.
	lookupGauge = metrics.NewRegisteredGauge("eth/protocols/snap/sync/lookup", nil)

	// boundaryAccountNodesGauge is the metric to track how many boundary trie
	// nodes in account trie are met.
	boundaryAccountNodesGauge = metrics.NewRegisteredGauge("eth/protocols/snap/sync/boundary/account", nil)

	// boundaryAccountNodesGauge is the metric to track how many boundary trie
	// nodes in storage tries are met.
	boundaryStorageNodesGauge = metrics.NewRegisteredGauge("eth/protocols/snap/sync/boundary/storage", nil)

	// smallStorageGauge is the metric to track how many storages are small enough
	// to retrieved in one or two request.
	smallStorageGauge = metrics.NewRegisteredGauge("eth/protocols/snap/sync/storage/small", nil)

	// largeStorageGauge is the metric to track how many storages are large enough
	// to retrieved concurrently.
	largeStorageGauge = metrics.NewRegisteredGauge("eth/protocols/snap/sync/storage/large", nil)

	// skipStorageHealingGauge is the metric to track how many storages are retrieved
	// in multiple requests but healing is not necessary.
	skipStorageHealingGauge = metrics.NewRegisteredGauge("eth/protocols/snap/sync/storage/noheal", nil)
)
