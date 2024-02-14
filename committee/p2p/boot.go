package p2p

import (
	"context"
	"fmt"
	realtimeDB "github.com/dTelecom/p2p-realtime-database"
	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
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
	DoSomething()
	DoSomethingElse()
}

type libP2PManagerImpl struct {
	host              host.Host
	dht               *dual.DHT
	announcementTopic *pubsub.Topic
	connectionManager *realtimeDB.ConnectionManager
	pubsub            *pubsub.PubSub
	pubSubBroadcaster *crdt.PubSubBroadcaster
	globalDS          datastore.Batching
	datastoreDs       *crdt.Datastore
}

const dbTopic = "crdt_db"

var RebroadcastingInterval = time.Millisecond * 10

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

func (l *libP2PManagerImpl) DoSomething() {
	err := l.datastoreDs.Put(context.Background(), datastore.NewKey("test"), []byte("test"))
	if err != nil {
		fmt.Println("error: ", err)
	}

	err = l.datastoreDs.Sync(context.Background(), datastore.NewKey("/"))
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func (l *libP2PManagerImpl) DoSomethingElse() {
	l.datastoreDs.Sync(context.Background(), datastore.NewKey("/"))

	val, err := l.datastoreDs.Get(context.Background(), datastore.NewKey("test"))
	if err != nil {
		fmt.Println("error: ", err)
	}
	_ = val

	fmt.Println(len(l.host.Network().Peers()))
	r, err := l.datastoreDs.Query(context.TODO(), query.Query{KeysOnly: true})
	if err != nil {
		fmt.Println(errors.Wrap(err, "crdt list query"))
	}

	l.datastoreDs.InternalStats()

	var keys []string
	for k := range r.Next() {
		keys = append(keys, k.Key)
	}

}

func (l *libP2PManagerImpl) setupHost(ctx context.Context, privKeyWrapper crypto.PrivKey) (host.Host, error) {
	port, _ := freeport.GetFreePort()
	// Create a new libp2p host
	sourceMultiAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))
	if err != nil {
		return nil, errors.Wrap(err, "create multi addr")
	}

	ds := ipfs_datastore.MutexWrap(datastore.NewMapDatastore())

	opts := ipfslite.Libp2pOptionsExtra
	// todo: setup datastore
	// TODO: add eth connection gater: https://github.com/dTelecom/p2p-realtime-database/blob/main/gater.go
	l.host, l.dht, err = ipfslite.SetupLibp2p(ctx, privKeyWrapper, nil, []multiaddr.Multiaddr{sourceMultiAddr}, ds, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not create libp2p host: %w", err)
	}

	l.connectionManager = realtimeDB.NewConnectionManager(l.host)

	traceme, err := pubsub.NewJSONTracer(fmt.Sprintf("/tmp/%d.json", i))
	if err != nil {
		return nil, fmt.Errorf("could not create tracer: %w", err)
	}
	i++

	l.pubsub, err = pubsub.NewGossipSub(context.Background(), l.host, pubsub.WithEventTracer(traceme))
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

	crdtOpts := crdt.DefaultOptions()
	crdtOpts.Logger = logging.Logger("p2p_logger")

	crdtOpts.RebroadcastInterval = RebroadcastingInterval
	crdtOpts.PutHook = func(k datastore.Key, v []byte) {
		fmt.Printf("[%s] Added: [%s] -> %s\n", time.Now().Format(time.RFC3339), k, string(v))
		// TODO: some validation goes here
	}
	// TODO: this probably never gets called
	crdtOpts.DeleteHook = func(k datastore.Key) {
		fmt.Printf("[%s] Removed: [%s]\n", time.Now().Format(time.RFC3339), k)
	}
	crdtOpts.RebroadcastInterval = time.Second

	l.datastoreDs, err = crdt.New(l.globalDS, datastore.NewKey(dbTopic), ipfs, l.pubSubBroadcaster, crdtOpts)
	if err != nil {
		return err
	}

	err = l.datastoreDs.Sync(ctx, datastore.NewKey("/"))

	return nil
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
