package p2p

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/ipfs/go-datastore"
	ipfs_datastore "github.com/ipfs/go-datastore/sync"
	crdt "github.com/ipfs/go-ds-crdt"
	logging "github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-kad-dht/dual"
	"github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/discovery"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	"github.com/libp2p/go-libp2p/p2p/discovery/util"
	"github.com/multiformats/go-multiaddr"
	"github.com/phayes/freeport"
	"github.com/pkg/errors"
	"github.com/synapsecns/sanguine/committee/db"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"log"
	"sync"
	"time"
)

type LibP2PManager interface {
	Host() host.Host // Expose host from manager
	// Start starts the libp2p manager.
	Start(ctx context.Context, bootstrapPeers []string) error
	DoSomething()
	DoSomethingElse() bool
	AddValidators(ctx context.Context, addr ...common.Address) error
	Address() common.Address
	GetSignature(ctx context.Context, address common.Address, chainID, nonce int) ([]byte, error)
	PutSignature(ctx context.Context, chainID, nonce int, signature []byte) error
}

type libP2PManagerImpl struct {
	host              host.Host
	dht               *dual.DHT
	announcementTopic *pubsub.Topic
	pubsub            *pubsub.PubSub
	pubSubBroadcaster crdt.Broadcaster
	globalDS          datastore.Batching
	// ipfs is the ipfs lite peer
	ipfs *ipfslite.Peer
	// datastoreDs is the datastore for the crdt
	// TODO: remove or something, not sure.
	datastoreDs *crdt.Datastore
	// discovery is used to discover peers
	discovery discovery.Discovery
	// datastoreMux is used to lock the datastores map
	datastoreMux sync.RWMutex
	// datastores stores the underlying datastore for each peer
	datastores map[common.Address]datastore.Batching
	// datastoreFactory is used to create new datastores
	datastoreFactory db.Datstores
	// address is the address of the node
	address common.Address
}

const dbTopic = "crdt_db"

var RebroadcastingInterval = time.Minute

// NewLibP2PManager creates a new libp2p manager.
// listenHost is the host to listen on.
//
// validators should be a list of addresses that are allowed to connect to the host. This should include the address of the
// node itself.
func NewLibP2PManager(ctx context.Context, auth signer.Signer, store db.Datstores) (LibP2PManager, error) {
	l := &libP2PManagerImpl{}
	_, err := l.setupHost(ctx, auth.PrivKey()) // call createHost function
	if err != nil {
		return nil, err
	}

	l.globalDS, err = store.GlobalDatastore()
	if err != nil {
		return nil, err
	}
	l.address = auth.Address()
	l.datastoreFactory = store
	l.datastores = make(map[common.Address]datastore.Batching)

	return l, nil
}

// Host returns the host from the manager.
// TODO: consider not exposing the host directly.
func (l *libP2PManagerImpl) Host() host.Host {
	return l.host
}

func (l *libP2PManagerImpl) Address() common.Address {
	return l.address
}

func (l *libP2PManagerImpl) DoSomething() {
	var err error
	for i := 500 - 1; i >= 0; i-- {
		err = l.datastoreDs.Put(context.Background(), datastore.NewKey(gofakeit.Word()), []byte("test"))
	}

	err = l.datastoreDs.Put(context.Background(), datastore.NewKey("test"), []byte("test"))
	if err != nil {
		fmt.Println("error: ", err)
	}

	err = l.datastoreDs.Sync(context.Background(), datastore.NewKey("/"))
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func (l *libP2PManagerImpl) DoSomethingElse() bool {
	fmt.Println(len(l.pubsub.ListPeers(dbTopic)))
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
	opts = append(opts, libp2p.Ping(true))
	// todo: setup datastore
	// TODO: add eth connection gater: https://github.com/dTelecom/p2p-realtime-database/blob/main/gater.go
	l.host, l.dht, err = ipfslite.SetupLibp2p(ctx, privKeyWrapper, nil, []multiaddr.Multiaddr{sourceMultiAddr}, ds, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not create libp2p host: %w", err)
	}

	return l.host, nil
}

func (l *libP2PManagerImpl) Start(ctx context.Context, bootstrapPeers []string) error {
	// setup ipfs
	peers, err := makePeers(bootstrapPeers)
	if err != nil {
		return err
	}

	l.ipfs, err = ipfslite.New(ctx, l.globalDS, nil, l.host, l.dht, &ipfslite.Config{})
	go l.Discover(ctx, l.host, l.dht, dbTopic)

	l.pubsub, err = pubsub.NewGossipSub(ctx, l.host)

	l.ipfs.Bootstrap(peers)
	for _, p := range peers {
		l.host.ConnManager().TagPeer(p.ID, "keep", 100)
	}

	if err != nil {
		return fmt.Errorf("could not create pubsub: %w", err)
	}

	go func() {
		for {
			time.Sleep(time.Second * 1)
			fmt.Println("pubsub peers: ", len(l.pubsub.ListPeers(dbTopic)))
			fmt.Println("global peers: ", len(l.host.Peerstore().Peers()))
		}
	}()

	l.pubSubBroadcaster, err = crdt.NewPubSubBroadcaster(ctx, l.pubsub, dbTopic)
	if err != nil {
		return err
	}

	crdtOpts := crdt.DefaultOptions()
	crdtOpts.Logger = logging.Logger("p2p_logger")

	crdtOpts.RebroadcastInterval = RebroadcastingInterval
	crdtOpts.MaxBatchDeltaSize = 1

	crdtOpts.PutHook = func(k datastore.Key, v []byte) {
		fmt.Printf("[%s] Added: [%s] -> %s\n", time.Now().Format(time.RFC3339), k, string(v))
		// TODO: some validation goes here
	}
	// TODO: this probably never gets called
	crdtOpts.DeleteHook = func(k datastore.Key) {
		fmt.Printf("[%s] Removed: [%s]\n", time.Now().Format(time.RFC3339), k)
	}
	crdtOpts.RebroadcastInterval = time.Second

	l.datastoreDs, err = crdt.New(l.globalDS, datastore.NewKey(dbTopic), l.ipfs, l.pubSubBroadcaster, crdtOpts)
	if err != nil {
		return err
	}

	err = l.datastoreDs.Sync(ctx, datastore.NewKey("/"))
	if err != nil {
		return err
	}

	return nil
}

func (l *libP2PManagerImpl) GetSignature(ctx context.Context, address common.Address, chainID, nonce int) ([]byte, error) {
	theirStore, ok := l.datastores[address]
	if !ok {
		return nil, fmt.Errorf("could not find datastore for address: %s", l.address.String())
	}

	// get from my store
	return theirStore.Get(ctx, datastore.NewKey(fmt.Sprintf("sig_%d_%d", chainID, nonce)))
}

// PutSignature puts a signature into the datastore.
func (l *libP2PManagerImpl) PutSignature(ctx context.Context, chainID, nonce int, signature []byte) error {
	myStore, ok := l.datastores[l.address]
	if !ok {
		return fmt.Errorf("could not find datastore for address: %s", l.address.String())
	}

	// add to my store
	err := myStore.Put(ctx, datastore.NewKey(fmt.Sprintf("sig_%d_%d", chainID, nonce)), signature)
	if err != nil {
		return fmt.Errorf("could not put signature: %w", err)
	}

	err = myStore.Sync(ctx, datastore.NewKey("/"))
	if err != nil {
		return fmt.Errorf("could not sync: %w", err)
	}
	return nil
}

func (l *libP2PManagerImpl) AddValidators(ctx context.Context, addr ...common.Address) error {
	// no point parallelizing this, it's all muxed.
	for _, a := range addr {
		err := l.addValidator(ctx, a)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *libP2PManagerImpl) addValidator(ctx context.Context, addr common.Address) error {
	l.datastoreMux.Lock()
	defer l.datastoreMux.Unlock()

	if l.datastores[addr] != nil {
		// validator already exists
		return nil
	}

	// create new datastore for validator
	ds, err := l.datastoreFactory.DatastoreForSigner(addr)
	if err != nil {
		return err
	}

	topic := fmt.Sprintf("crdt_db_%s", addr.String())

	err = l.pubsub.RegisterTopicValidator(topic, func(ctx context.Context, from peer.ID, msg *pubsub.Message) bool {
		// TODO: see p2p-realtime-database, convert the from peer id to an address
		// and then make sure it matches the validator

		return true
	})
	if err != nil {
		return fmt.Errorf("could not register topic validator: %w", err)
	}

	// TODO: there's an edge case where if we error after this line is succesful, newpubsubbroadcaster will not be able ot be created again
	// this should be saved in the same manner as datastores
	pubSubBroadcaster, err := crdt.NewPubSubBroadcaster(ctx, l.pubsub, topic)
	if err != nil {
		return fmt.Errorf("could not create pubsub broadcaster: %w", err)
	}

	crdtOpts := crdt.DefaultOptions()
	crdtOpts.Logger = logging.Logger(fmt.Sprintf("%s_logger", addr.String()))
	crdtOpts.RebroadcastInterval = RebroadcastingInterval
	crdtOpts.MaxBatchDeltaSize = 1

	crdtOpts.PutHook = func(k datastore.Key, v []byte) {
		fmt.Printf("[%s] Added: [%s] -> %s\n", time.Now().Format(time.RFC3339), k, string(v))
		// TODO: some validation goes here
	}
	// TODO: this probably never gets called
	crdtOpts.DeleteHook = func(k datastore.Key) {
		fmt.Printf("[%s] Removed: [%s]\n", time.Now().Format(time.RFC3339), k)
	}
	crdtOpts.RebroadcastInterval = time.Second

	l.datastores[addr], err = crdt.New(ds, datastore.NewKey(topic), l.ipfs, pubSubBroadcaster, crdtOpts)
	if err != nil {
		return err
	}

	err = l.datastores[addr].Sync(ctx, datastore.NewKey("/"))
	if err != nil {
		return err
	}

	return nil
}

func (l *libP2PManagerImpl) Discover(ctx context.Context, h host.Host, dht *dual.DHT, rendezvous string) {
	routingDiscovery := routing.NewRoutingDiscovery(dht)
	util.Advertise(ctx, routingDiscovery, rendezvous)

	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:

			peers, err := util.FindPeers(ctx, routingDiscovery, rendezvous)
			if err != nil {
				log.Fatal(err)
			}

			for _, p := range peers {
				if p.ID == h.ID() {
					continue
				}
				if h.Network().Connectedness(p.ID) != network.Connected {
					_, err = h.Network().DialPeer(ctx, p.ID)
					fmt.Printf("Connected to peer %s\n", p.ID)
					if err != nil {
						continue
					}
				}
			}
		}
	}
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
