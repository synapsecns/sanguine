package p2p

import (
	"context"
	"encoding/json"
	"fmt"
	realtimeDB "github.com/dTelecom/p2p-realtime-database"
	"github.com/google/uuid"
	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/ipfs/go-datastore"
	ipfs_datastore "github.com/ipfs/go-datastore/sync"
	crdt "github.com/ipfs/go-ds-crdt"
	"github.com/ipfs/go-log"
	logging "github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-kad-dht/dual"
	"github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	"github.com/libp2p/go-libp2p/p2p/discovery/util"
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
	DoSomethingElse() bool
}

type libP2PManagerImpl struct {
	host              host.Host
	dht               *dual.DHT
	connectionManager *realtimeDB.ConnectionManager
	pubsub            *pubsub.PubSub
	pubSubBroadcaster *crdt.PubSubBroadcaster
	globalDS          datastore.Batching
	datastoreDs       *crdt.Datastore
	discoveryTopic    *pubsub.Topic
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

func (l *libP2PManagerImpl) DoSomethingElse() bool {
	for f := 0; f < 400; f++ {
		l.datastoreDs.Sync(context.Background(), datastore.NewKey("/"))
	}
	time.Sleep(time.Second)

	val, err := l.datastoreDs.Get(context.Background(), datastore.NewKey("test"))
	if err != nil {
		fmt.Println("error: ", err)
	}

	return val != nil
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
	opts = append(opts, libp2p.EnableRelay(), libp2p.Ping(true))
	// todo: setup datastore
	// TODO: add eth connection gater: https://github.com/dTelecom/p2p-realtime-database/blob/main/gater.go
	l.host, l.dht, err = ipfslite.SetupLibp2p(ctx, privKeyWrapper, nil, []multiaddr.Multiaddr{sourceMultiAddr}, ds, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not create libp2p host: %w", err)
	}

	l.connectionManager = realtimeDB.NewConnectionManager(l.host)

	for _, myPeer := range l.host.Peerstore().Peers() {
		l.host.ConnManager().TagPeer(myPeer, "keep", 100)
	}

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

	ipfs, err := ipfslite.New(ctx, l.globalDS, nil, l.host, l.dht, &ipfslite.Config{
		ReprovideInterval: time.Second * 5,
		//UncachedBlockstore: true,
	})
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

	l.discoveryTopic, err = l.pubsub.Join(realtimeDB.DiscoveryTag)

	l.startDiscovery(ctx)
	l.netPingPeers(ctx, "")

	return nil
}

var logger = log.Logger("p2p_logger")

func (l *libP2PManagerImpl) startDiscovery(ctx context.Context) {
	rendezvous := realtimeDB.DiscoveryTag
	routingDiscovery := routing.NewRoutingDiscovery(l.dht)
	util.Advertise(ctx, routingDiscovery, rendezvous)

	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	go func() {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			peers, err := util.FindPeers(ctx, routingDiscovery, rendezvous)
			if err != nil {
				logger.Errorf("discrovery find peers error %s", err)
				return
			}
			for _, p := range peers {
				logger.Errorf("found peer %s", p.String())

				if p.ID == l.host.ID() {
					continue
				}

				if l.host.Network().Connectedness(p.ID) != network.Connected {
					_, err = l.host.Network().DialPeer(ctx, p.ID)
					if err != nil {
						logger.Errorf("discrovery connected to peer error %s: %s", p.ID.String(), err)
						continue
					}
					logger.Infof("discrovery connected to peer %s\n", p.ID.String())
				}
			}
		}
	}()
}

func (l *libP2PManagerImpl) netPingPeers(ctx context.Context, netTopic string) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				_, err := l.Publish(ctx, []byte(realtimeDB.NetSubscriptionPublishValue))
				if err != nil {
					logger.Errorf("try publish message to net ps topic: %s", err)
					if errors.Is(err, pubsub.ErrTopicClosed) {
						return
					}
				}
				time.Sleep(20 * time.Second)
			}
		}
	}()
}

func (l *libP2PManagerImpl) Publish(ctx context.Context, value interface{}, opts ...pubsub.PubOpt) (realtimeDB.Event, error) {
	event := realtimeDB.Event{
		ID:         uuid.New().String(),
		Message:    value,
		FromPeerId: l.host.ID().String(),
	}
	marshaled, err := json.Marshal(event)
	if err != nil {
		return realtimeDB.Event{}, errors.Wrap(err, "try marshal message")
	}

	err = l.discoveryTopic.Publish(ctx, marshaled, opts...)
	if err != nil {
		return realtimeDB.Event{}, errors.Wrap(err, "pub sub publish message")
	}

	return event, nil
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
