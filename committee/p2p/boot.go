package p2p

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	realtimeDB "github.com/dTelecom/p2p-realtime-database"
	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/ipfs/go-datastore"
	ipfs_datastore "github.com/ipfs/go-datastore/sync"
	crdt "github.com/ipfs/go-ds-crdt"
	logging "github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p-kad-dht/dual"
	"github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"github.com/phayes/freeport"
	"github.com/pkg/errors"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"time"
)

type LibP2PManager interface {
	Host() host.Host // Expose host from manager
	// Start starts the libp2p manager.
	Start(ctx context.Context, bootstrapPeers []string) error
}

type libP2PManagerImpl struct {
	host              host.Host
	dht               *dual.DHT
	announcementTopic *pubsub.Topic
	connectionManager *realtimeDB.ConnectionManager
	pubsub            *pubsub.PubSub
	pubSubBroadcaster *crdt.PubSubBroadcaster
	globalDS          datastore.Batching
	datastoreDs       datastore.Batching
}

const dbTopic = "crdt_db"

var RebroadcastingInterval = time.Second * 30

// NewLibP2PManager creates a new libp2p manager.
// listenHost is the host to listen on.
//
// TODO: we need to figure out how this works across multiple nodes
func NewLibP2PManager(ctx context.Context, auth signer.Signer) (LibP2PManager, error) {
	l := &libP2PManagerImpl{}
	_, err := l.setupHost(ctx, auth.PrivKey()) // call createHost function
	if err != nil {
		return nil, err
	}

	l.pubSubBroadcaster, err = crdt.NewPubSubBroadcaster(ctx, l.pubsub, dbTopic)
	if err != nil {
		return nil, err
	}

	l.globalDS = ipfs_datastore.MutexWrap(datastore.NewMapDatastore())

	return l, nil
}

// Host returns the host from the manager.
// TODO: consider not exposing the host directly
func (l *libP2PManagerImpl) Host() host.Host {
	return l.host
}

func (l *libP2PManagerImpl) setupHost(ctx context.Context, privKeyWrapper crypto.PrivKey) (host.Host, error) {
	port, _ := freeport.GetFreePort()
	// Create a new libp2p host
	sourceMultiAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))
	if err != nil {
		return nil, errors.Wrap(err, "create multi addr")
	}

	ds := ipfs_datastore.MutexWrap(datastore.NewMapDatastore())

	// todo: setup datastore
	// TODO: add eth connection gater: https://github.com/dTelecom/p2p-realtime-database/blob/main/gater.go
	l.host, l.dht, err = ipfslite.SetupLibp2p(ctx, privKeyWrapper, nil, []multiaddr.Multiaddr{sourceMultiAddr}, ds, ipfslite.Libp2pOptionsExtra...)
	if err != nil {
		return nil, fmt.Errorf("could not create libp2p host: %w", err)
	}

	l.connectionManager = realtimeDB.NewConnectionManager(l.host)

	l.pubsub, err = pubsub.NewGossipSub(context.Background(), l.host)
	if err != nil {
		return nil, fmt.Errorf("could not create pubsub: %w", err)
	}

	return l.host, nil
}

var i = 0

func (l *libP2PManagerImpl) Start(ctx context.Context, bootstrapPeers []string) error {
	// setup ipfs
	peers, err := makePeers(bootstrapPeers)
	if err != nil {
		return err
	}

	ipfs, err := ipfslite.New(ctx, l.globalDS, nil, l.host, l.dht, &ipfslite.Config{})
	ipfs.Bootstrap(peers)

	crtdOpts := crdt.DefaultOptions()
	crtdOpts.Logger = logging.Logger("p2p_logger")
	crtdOpts.RebroadcastInterval = RebroadcastingInterval
	crtdOpts.PutHook = func(k datastore.Key, v []byte) {
		fmt.Printf("[%s] Added: [%s] -> %s\n", time.Now().Format(time.RFC3339), k, string(v))
		// TODO: some validation goes here
	}
	// TODO: this probably never gets called
	crtdOpts.DeleteHook = func(k datastore.Key) {
		fmt.Printf("[%s] Removed: [%s]\n", time.Now().Format(time.RFC3339), k)
	}
	crtdOpts.RebroadcastInterval = time.Second

	l.datastoreDs, err = crdt.New(l.globalDS, datastore.NewKey(dbTopic), ipfs, l.pubSubBroadcaster, crtdOpts)
	if err != nil {
		return err
	}

	err = l.datastoreDs.Sync(ctx, datastore.NewKey("/"))

	return nil
}

func HashString(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

func (l *libP2PManagerImpl) Select(key string, values [][]byte) (int, error) {
	// TODO: implement me
	return 0, nil
}

func makePeers(peers []string) ([]peer.AddrInfo, error) {
	var p []peer.AddrInfo
	for _, addr := range peers {
		maddr, err := multiaddr.NewMultiaddr(addr)
		if err != nil {
			return nil, fmt.Errorf("could not create multiaddr: %w", err)
		}
		info, err := peer.AddrInfoFromP2pAddr(maddr)
		if err != nil {
			return nil, fmt.Errorf("could not create peer info: %w", err)
		}
		p = append(p, *info)
	}
	return p, nil

}
