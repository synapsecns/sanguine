package p2p

import "github.com/libp2p/go-libp2p/core/peer"

func MakePeers(peers []string) ([]peer.AddrInfo, error) {
	// nolint: wrapcheck
	return makePeers(peers)
}
