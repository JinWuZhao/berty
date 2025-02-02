package ipfsutil

import (
	"context"
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"math/rand"
	"testing"

	rendezvous "github.com/berty/go-libp2p-rendezvous"
	p2p_rpdb "github.com/berty/go-libp2p-rendezvous/db/sqlcipher"
	ds "github.com/ipfs/go-datastore"
	dsync "github.com/ipfs/go-datastore/sync"
	ipfs_cfg "github.com/ipfs/kubo/config"
	ipfs_core "github.com/ipfs/kubo/core"
	ipfs_mock "github.com/ipfs/kubo/core/mock"
	ipfs_repo "github.com/ipfs/kubo/repo"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	p2p_ci "github.com/libp2p/go-libp2p/core/crypto"
	host "github.com/libp2p/go-libp2p/core/host"
	p2pnetwork "github.com/libp2p/go-libp2p/core/network"
	p2p_peer "github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/core/routing"
	p2p_mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	ipfs_mobile "berty.tech/berty/v2/go/internal/ipfsutil/mobile"
	tinder "berty.tech/berty/v2/go/internal/tinder"
)

// CoreAPIMock implements ipfs.CoreAPI and adds some debugging helpers
type CoreAPIMock interface {
	API() ExtendedCoreAPI

	PubSub() *pubsub.PubSub
	Tinder() *tinder.Service
	MockNetwork() p2p_mocknet.Mocknet
	MockNode() *ipfs_core.IpfsNode
	Close()
}

func getOrCreatePrivateKeyFromDatastore(t testing.TB, ctx context.Context, datastore ds.Datastore) p2p_ci.PrivKey {
	const datastoreKeyForPrivateKey = "p2p_private_key"

	privkeyb, err := datastore.Get(ctx, ds.NewKey("private_key"))
	if err == ds.ErrNotFound {
		priv, _, err := p2p_ci.GenerateKeyPairWithReader(p2p_ci.RSA, 2048, crand.Reader)
		if err != nil {
			t.Fatalf("failed to generate pair key: %v", err)
		}

		privkeyb, err := p2p_ci.MarshalPrivateKey(priv)
		if err != nil {
			t.Fatalf("failed to get raw priv key: %v", err)
		}

		if err := datastore.Put(ctx, ds.NewKey(datastoreKeyForPrivateKey), privkeyb); err != nil {
			t.Fatalf("failed to save priv key: %v", err)
		}

		return priv
	} else if err != nil {
		t.Fatalf("failed to get value from datastore: %v", err)
	}

	priv, err := p2p_ci.UnmarshalPrivateKey(privkeyb)
	if err != nil {
		t.Fatalf("failed to unmarshal priv key: %v", err)
	}

	return priv
}

func TestingRepo(t testing.TB, ctx context.Context, datastore ds.Datastore) ipfs_repo.Repo {
	t.Helper()

	c := ipfs_cfg.Config{}
	priv := getOrCreatePrivateKeyFromDatastore(t, ctx, datastore)

	pid, err := p2p_peer.IDFromPublicKey(priv.GetPublic())
	if err != nil {
		t.Fatalf("failed to get pid from pub key: %v", err)
	}

	privkeyb, err := p2p_ci.MarshalPrivateKey(priv)
	if err != nil {
		t.Fatalf("failed to get raw priv key: %v", err)
	}

	c.Bootstrap = []string{}
	c.Addresses.Swarm = []string{"/ip6/::/tcp/0"}
	c.Identity.PeerID = pid.Pretty()
	c.Identity.PrivKey = base64.StdEncoding.EncodeToString(privkeyb)

	if datastore == nil {
		datastore = ds.NewMapDatastore()
	}
	dstore := dsync.MutexWrap(datastore)

	return &ipfs_repo.Mock{
		D: dstore,
		C: c,
	}
}

type TestingAPIOpts struct {
	Logger    *zap.Logger
	Mocknet   p2p_mocknet.Mocknet
	RDVPeer   p2p_peer.AddrInfo
	Datastore ds.Batching
}

// TestingCoreAPIUsingMockNet returns a fully initialized mocked Core API with the given mocknet
func TestingCoreAPIUsingMockNet(ctx context.Context, t testing.TB, opts *TestingAPIOpts) (CoreAPIMock, func()) {
	if opts.Logger == nil {
		opts.Logger = zap.NewNop()
	}

	datastore := opts.Datastore
	if datastore == nil {
		datastore = dsync.MutexWrap(ds.NewMapDatastore())
	}

	repo := TestingRepo(t, ctx, datastore)

	var ps *pubsub.PubSub
	var stinder *tinder.Service
	var discAdaptater *tinder.DiscoveryAdaptater

	configureRouting := func(h host.Host, r routing.Routing) error {
		var err error
		drivers := []tinder.IDriver{}
		if opts.RDVPeer.ID != "" {
			h.Peerstore().AddAddrs(opts.RDVPeer.ID, opts.RDVPeer.Addrs, peerstore.PermanentAddrTTL)
			cl := rendezvous.NewSyncInMemClient(ctx, h)
			ms := tinder.NewRendezvousDiscovery(opts.Logger.Named("rdvp"), h, opts.RDVPeer.ID, rendezvous.DefaultAddrFactory, rand.New(rand.NewSource(rand.Int63())), cl)
			if _, err = opts.Mocknet.LinkPeers(h.ID(), opts.RDVPeer.ID); err != nil {
				return err
			}

			drivers = append(drivers, ms)
		}

		if r != nil {
			driver := tinder.NewRoutingDiscoveryDriver("dht", r)
			drivers = append(drivers, driver)
		}

		// enable discovery monitor
		stinder, err = tinder.NewService(h, opts.Logger, drivers...)
		if err != nil {
			return fmt.Errorf("unable to monitor discovery driver: %w", err)
		}

		discAdaptater = tinder.NewDiscoveryAdaptater(opts.Logger, stinder)
		ps, err = pubsub.NewGossipSub(ctx, h, pubsub.WithDiscovery(discAdaptater))

		return err
	}

	mrepo := ipfs_mobile.NewRepoMobile("", repo)
	mnode, err := NewIPFSMobile(ctx, mrepo, &MobileOptions{
		HostOption:        ipfs_mock.MockHostOption(opts.Mocknet),
		RoutingConfigFunc: configureRouting,
		ExtraOpts: map[string]bool{
			"pubsub": false,
		},
	})

	require.NoError(t, err, "failed to initialize IPFS node mock")
	require.NotNil(t, ps, "pubsub should not be nil")
	require.NotNil(t, stinder, "discovery should not be nil")

	exapi, err := NewExtendedCoreAPIFromNode(mnode.IpfsNode)
	require.NoError(t, err, "unable to extend core api from node")

	psapi := NewPubSubAPI(ctx, opts.Logger, ps)
	exapi = InjectPubSubCoreAPIExtendedAdapter(exapi, psapi)
	EnableConnLogger(ctx, opts.Logger, mnode.PeerHost())

	api := &coreAPIMock{
		coreapi: exapi,
		mocknet: opts.Mocknet,
		pubsub:  ps,
		node:    mnode.IpfsNode,
		tinder:  stinder,
	}

	return api, func() {
		if discAdaptater != nil {
			discAdaptater.Close()
		}

		_ = mnode.Close()
		_ = mnode.PeerHost().Close()
		_ = repo.Close()
	}
}

// TestingCoreAPI returns a fully initialized mocked Core API.
// If you want to do some tests involving multiple peers you should use
// `TestingCoreAPIUsingMockNet` with the same mocknet instead.
func TestingCoreAPI(ctx context.Context, t testing.TB) (CoreAPIMock, func()) {
	t.Helper()

	m := p2p_mocknet.New()
	defer m.Close()

	rdvPeer, err := m.GenPeer()
	require.NoError(t, err)

	_, cleanrdvp := TestingRDVP(ctx, t, rdvPeer)
	api, cleanapi := TestingCoreAPIUsingMockNet(ctx, t, &TestingAPIOpts{
		Mocknet: m,
		RDVPeer: rdvPeer.Network().Peerstore().PeerInfo(rdvPeer.ID()),
	})

	cleanup := func() {
		cleanapi()
		cleanrdvp()
		_ = rdvPeer.Close()
	}
	return api, cleanup
}

func TestingRDVP(ctx context.Context, t testing.TB, h host.Host) (*rendezvous.RendezvousService, func()) {
	db, err := p2p_rpdb.OpenDB(ctx, ":memory:")
	require.NoError(t, err)

	provider, err := rendezvous.NewSyncInMemProvider(h)
	require.NoError(t, err)

	svc := rendezvous.NewRendezvousService(h, db, provider)
	cleanup := func() {
		_ = db.Close() // dont use me for now as db is open in_memory
	}
	return svc, cleanup
}

type coreAPIMock struct {
	coreapi ExtendedCoreAPI

	pubsub  *pubsub.PubSub
	mocknet p2p_mocknet.Mocknet
	node    *ipfs_core.IpfsNode
	tinder  *tinder.Service
}

func (m *coreAPIMock) ConnMgr() ConnMgr {
	return m.node.PeerHost.ConnManager()
}

func (m *coreAPIMock) NewStream(ctx context.Context, p p2p_peer.ID, pids ...protocol.ID) (p2pnetwork.Stream, error) {
	return m.node.PeerHost.NewStream(ctx, p, pids...)
}

func (m *coreAPIMock) SetStreamHandler(pid protocol.ID, handler p2pnetwork.StreamHandler) {
	m.node.PeerHost.SetStreamHandler(pid, handler)
}

func (m *coreAPIMock) RemoveStreamHandler(pid protocol.ID) {
	m.node.PeerHost.RemoveStreamHandler(pid)
}

func (m *coreAPIMock) API() ExtendedCoreAPI {
	return m.coreapi
}

func (m *coreAPIMock) MockNetwork() p2p_mocknet.Mocknet {
	return m.mocknet
}

func (m *coreAPIMock) MockNode() *ipfs_core.IpfsNode {
	return m.node
}

func (m *coreAPIMock) PubSub() *pubsub.PubSub {
	return m.pubsub
}

func (m *coreAPIMock) Tinder() *tinder.Service {
	return m.tinder
}

func (m *coreAPIMock) Close() {
	m.node.Close()
}
