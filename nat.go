// Deprecated: This package has moved into go-libp2p as a sub-package: github.com/libp2p/go-libp2p/p2p/net/nat.
package nat

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p/p2p/net/nat"
)

// ErrNoMapping signals no mapping exists for an address
// Deprecated: use github.com/libp2p/go-libp2p/p2p/net/nat.ErrNoMapping instead.
var ErrNoMapping = nat.ErrNoMapping

// MappingDuration is a default port mapping duration.
// Port mappings are renewed every (MappingDuration / 3)
// Deprecated: use github.com/libp2p/go-libp2p/p2p/net/nat.MappingDuration instead.
const MappingDuration = time.Second * 60

// CacheTime is the time a mapping will cache an external address for
// Deprecated: use github.com/libp2p/go-libp2p/p2p/net/nat.CacheTime instead.
const CacheTime = time.Second * 15

// DiscoverNAT looks for a NAT device in the network and
// returns an object that can manage port mappings.
// Deprecated: use github.com/libp2p/go-libp2p/p2p/net/nat.DiscoverNat instead.
func DiscoverNAT(ctx context.Context) (*NAT, error) {
	return nat.DiscoverNAT(ctx)
}

// NAT is an object that manages address port mappings in
// NATs (Network Address Translators). It is a long-running
// service that will periodically renew port mappings,
// and keep an up-to-date list of all the external addresses.
// Deprecated: use github.com/libp2p/go-libp2p/p2p/net/nat.Nat instead.
type NAT = nat.NAT
