package node_test

import (
	"math/big"
	"sync"
	"testing"

	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/committee/config"
	"github.com/synapsecns/sanguine/committee/contracts/synapsemodule"
	"github.com/synapsecns/sanguine/committee/node"
	"github.com/synapsecns/sanguine/committee/testutil"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
	"github.com/synapsecns/sanguine/services/omnirpc/testhelper"
)

type NodeSuite struct {
	*testsuite.TestSuite
	metrics       metrics.Handler
	originChain   backends.SimulatedTestBackend
	destChain     backends.SimulatedTestBackend
	deployManager *testutil.DeployManager
	originModule  *synapsemodule.SynapseModuleRef
	destModule    *synapsemodule.SynapseModuleRef
	nodes         []*node.Node
	omnirpcURL    string
}

func NewNodeSuite(tb testing.TB) *NodeSuite {
	tb.Helper()
	return &NodeSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		metrics:   metrics.NewNullHandler(),
	}
}

func (n *NodeSuite) SetupTest() {
	n.TestSuite.SetupTest()
	n.deployManager = testutil.NewDeployManager(n.T())

	var originInfo, destInfo contracts.DeployedContract
	// create origin and destination chains
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		n.originChain = geth.NewEmbeddedBackendForChainID(n.GetTestContext(), n.T(), big.NewInt(1))
		originInfo, n.originModule = n.deployManager.GetSynapseModule(n.GetTestContext(), n.originChain)
	}()
	go func() {
		defer wg.Done()
		n.destChain = geth.NewEmbeddedBackendForChainID(n.GetTestContext(), n.T(), big.NewInt(2))
		destInfo, n.destModule = n.deployManager.GetSynapseModule(n.GetTestContext(), n.destChain)
	}()
	wg.Wait()

	n.omnirpcURL = testhelper.NewOmnirpcServer(n.GetTestContext(), n.T(), n.originChain, n.destChain)

	for i := 0; i < 3; i++ {
		n.makeNode()
	}

	var interchainValidators []common.Address
	for _, otherNode := range n.nodes {
		interchainValidators = append(interchainValidators, otherNode.Address())
	}

	n.setValidators(interchainValidators, n.originChain, originInfo, n.originModule)
	n.setValidators(interchainValidators, n.destChain, destInfo, n.destModule)

	for _, on := range n.nodes {
		on := on
		go func() {
			err := on.Start(n.GetTestContext())
			n.NoError(err)
		}()
	}
}

func (n *NodeSuite) setValidators(validators []common.Address, backend backends.SimulatedTestBackend, info contracts.DeployedContract, contract *synapsemodule.SynapseModuleRef) {
	transactOpts := backend.GetTxContext(n.GetTestContext(), info.OwnerPtr())

	for _, validator := range validators {
		tx, err := contract.AddVerifier(transactOpts.TransactOpts, validator)
		n.NoError(err)
		backend.WaitForConfirmation(n.GetTestContext(), tx)
	}
}

func (n *NodeSuite) makeNode() {
	testWallet, err := wallet.FromRandom()
	n.NoError(err)

	n.originChain.FundAccount(n.GetTestContext(), testWallet.Address(), *new(big.Int).SetUint64(params.Ether))
	n.destChain.FundAccount(n.GetTestContext(), testWallet.Address(), *new(big.Int).SetUint64(params.Ether))

	shouldRelay := false
	if len(n.nodes) == 0 {
		shouldRelay = true
	}

	var bootstrapPeers []string
	for _, prevNode := range n.nodes {
		bootstrapPeers = append(bootstrapPeers, prevNode.IPFSAddress()...)
	}

	cfg := config.Config{
		Chains: map[int]string{
			1: n.originModule.Address().String(),
			2: n.destModule.Address().String(),
		},
		OmnirpcURL: n.omnirpcURL,
		Database: config.DatabaseConfig{
			Type: dbcommon.Sqlite.String(),
			DSN:  filet.TmpDir(n.T(), ""),
		},
		Signer: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(n.T(), "", testWallet.PrivateKeyHex()).Name(),
		},
		SubmitterConfig: submitterConfig.Config{},
		ShouldRelay:     shouldRelay,
		BootstrapPeers:  bootstrapPeers,
		P2PPort:         freeport.GetPort(),
	}

	myNode, err := node.NewNode(n.GetTestContext(), n.metrics, cfg)
	n.NoError(err)

	n.nodes = append(n.nodes, myNode)
}

func TestNodeSuite(t *testing.T) {
	suite.Run(t, NewNodeSuite(t))
}
