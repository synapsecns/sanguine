package main

import (
	"context"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/agents/contracts/bondingmanager"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/contracts/test/pingpongclient"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/retry"
	ethergoChain "github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v2"
)

type loadConfig struct {
	SummitDomainID     int                 `yaml:"summit_domain_id"`
	OmniRPCUrl         string              `yaml:"omnirpc_url"`
	ScribeURL          string              `yaml:"scribe_url"`
	BondingManagerAddr string              `yaml:"bonding_manager_addr"`
	Chains             map[int]chainConfig `yaml:"chains"`
}

type chainConfig struct {
	MessageAddr     string            `yaml:"message_contract_addr"`
	OriginAddr      string            `yaml:"origin_addr"`
	DestinationAddr string            `yaml:"destination_addr"`
	Agents          map[string]string `yaml:"agents"`
}

func getLoadConfig(path string) (cfg loadConfig, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &cfg)
	return
}

func getSigner(privateKey string) (signer *localsigner.Signer, err error) {
	localWallet, err := wallet.FromHex(privateKey)
	if err != nil {
		return
	}
	signer = localsigner.NewSigner(localWallet.PrivateKey())
	return signer, nil
}

func getMessageRoutes(chainConfigs map[int]chainConfig, numRoutes int) (routes [][2]int, err error) {
	chainIDs := []int{}
	for chainID := range chainConfigs {
		chainIDs = append(chainIDs, chainID)
	}
	routes = [][2]int{}
	for _, origin := range chainIDs {
		for _, destination := range chainIDs {
			if numRoutes > 0 && len(routes) >= numRoutes {
				return routes, nil
			}
			if origin == destination {
				continue
			}
			routes = append(routes, [2]int{origin, destination})
		}
	}
	return routes, nil
}

func handleLog(log *ethTypes.Log, chainID uint32) (err error) {
	// parse the log and print output
	var event interface{}
	if event, err = pingPongParser.ParsePingSent(*log); err == nil {
		pingSentEvent, ok := event.(*pingpongclient.PingPongClientPingSent)
		if !ok {
			return fmt.Errorf("could not parse ping sent event")
		}
		fmt.Printf("Parsed ping sent on chain %d [ID=%d]\n", chainID, pingSentEvent.PingId.Int64())
	}
	if event, err = pingPongParser.ParsePingReceived(*log); err == nil {
		pingReceivedEvent, ok := event.(*pingpongclient.PingPongClientPingReceived)
		if !ok {
			return fmt.Errorf("could not parse ping received event")
		}
		fmt.Printf("Parsed ping received on chain %d [ID=%d]\n", chainID, pingReceivedEvent.PingId.Int64())
	}
	if event, err = pingPongParser.ParsePongSent(*log); err == nil {
		pongSentEvent, ok := event.(*pingpongclient.PingPongClientPongSent)
		if !ok {
			return fmt.Errorf("could not parse pong sent event")
		}
		fmt.Printf("Parsed pong sent on chain %d [ID=%d]\n", chainID, pongSentEvent.PingId.Int64())
	}
	if event, err = pingPongParser.ParsePongReceived(*log); err == nil {
		pongReceivedEvent, ok := event.(*pingpongclient.PingPongClientPongReceived)
		if !ok {
			return fmt.Errorf("could not parse pong received event")
		}
		fmt.Printf("Parsed pong received on chain %d [ID=%d]\n", chainID, pongReceivedEvent.PingId.Int64())
	}
	if event, ok := originParser.ParseSent(*log); ok {
		message, ok := event.(types.Message)
		if !ok {
			return fmt.Errorf("could not parse message sent event")
		}
		leafBytes, err := message.ToLeaf()
		if err != nil {
			fmt.Printf("Error getting message leaf: %v\n", err)
			return err
		}
		leaf := common.BytesToHash(leafBytes[:])

		// make sure it's a ping that we sent
		_, ok = sentTxes.Load(log.TxHash)
		if ok {
			messages.Store(leaf, message)
			numSent++
			fmt.Printf("Parsed message sent from %d to %d [leaf=%s,num=%d,nonce=%d]\n", message.OriginDomain(), message.DestinationDomain(), leaf, numSent, message.Nonce())
		}
	}
	if event, err = destinationParser.ParseExecuted(*log); err == nil {
		messageExecutedEvent, ok := event.(*destination.DestinationExecuted)
		if !ok {
			return fmt.Errorf("could not parse message executed event")
		}
		leaf := common.BytesToHash(messageExecutedEvent.MessageHash[:])

		// make sure it's a message that we sent
		_, ok = messages.Load(leaf)
		if ok {
			fmt.Printf("\u2713 Parsed message executed on chain %d [leaf=%s]\n", chainID, leaf)
		}
	}
	return nil
}

func getBalance(ctx context.Context, omniRPCClient omniClient.RPCClient, addr common.Address, chainID int) (balance *big.Int, err error) {
	client, err := omniRPCClient.GetChainClient(ctx, chainID)
	if err != nil {
		err = fmt.Errorf("could not get chain client: %w", err)
		return
	}
	balance, err = client.BalanceAt(ctx, addr, nil)
	if err != nil {
		err = fmt.Errorf("could not get balance: %w", err)
		return
	}
	return
}

const balanceThreshold = 0.1

func checkAgentBalances(ctx context.Context, omniRPCClient omniClient.RPCClient, chains map[int]chainConfig) (ok bool, err error) {
	fmt.Println("Checking agent balances...")
	ok = true
	for chainID, chainCfg := range chains {
		for agent, addr := range chainCfg.Agents {
			balance, err := getBalance(ctx, omniRPCClient, common.HexToAddress(addr), chainID)
			if err != nil {
				ok = false
				return false, err
			}
			balanceFlt := new(big.Float).SetInt(balance)
			balanceFlt.Quo(balanceFlt, new(big.Float).SetInt(big.NewInt(params.Ether)))
			balanceFlt64, _ := balanceFlt.Float64()
			if balanceFlt64 < balanceThreshold {
				ok = false
				fmt.Printf("Balance for %s on chain %d is below threshold of %f: %f [%s]\n", agent, chainID, balanceThreshold, balanceFlt64, addr)
			} else {
				fmt.Printf("Balance for %s on chain %d: %s [%s]\n", agent, chainID, balanceFlt.String(), addr)
			}
		}
	}
	return ok, nil
}

func checkAgentStatuses(ctx context.Context, omniRPCClient omniClient.RPCClient, cfg loadConfig) (ok bool, err error) {
	fmt.Println("Checking agent statuses...")

	cClient, err := ethergoChain.NewFromURL(ctx, omniRPCClient.GetEndpoint(cfg.SummitDomainID, 1))
	if err != nil {
		return ok, err
	}

	bondingManagerContract, err := evm.NewBondingManagerContract(ctx, cClient, common.HexToAddress(cfg.BondingManagerAddr))
	if err != nil {
		return ok, err
	}

	ok = true
	for chainID, chainCfg := range cfg.Chains {
		for agent, addr := range chainCfg.Agents {
			if strings.Contains(agent, "notary") {
				status, err := bondingManagerContract.GetAgentStatus(ctx, common.HexToAddress(addr))
				if err != nil {
					return false, fmt.Errorf("could not get agent status for agent: %s: %w", agent, err)
				}
				fmt.Printf("Got agent status for agent %s on chain %d: %s\n", agent, chainID, status.Flag().String())
				if status.Flag() != types.AgentFlagActive {
					ok = false
				}
			}
		}
	}
	return ok, nil
}

var pingPongParser *pingpongclient.PingPongClientFilterer
var originParser origin.Parser
var bondingManagerParser bondingmanager.Parser
var destinationParser *destination.DestinationFilterer
var numIters, numExecuted, executeInterval, numRoutes, numSent, totalExecuted int
var messages = &sync.Map{}
var sentTxes = &sync.Map{}

const eventBufferSize = 1000
const executionTimeout = 600

func main() {
	var loadConfigPath string
	var privateKey string
	flag.StringVar(&loadConfigPath, "c", "", "path to load config")
	flag.StringVar(&privateKey, "p", "", "private key")
	flag.IntVar(&numIters, "n", 1, "number of message route iterations")
	flag.IntVar(&numRoutes, "r", 0, "number of routes")
	flag.IntVar(&executeInterval, "i", 30, "execution interval (seconds)")
	flag.Parse()
	if len(loadConfigPath) == 0 {
		panic("expected load config path to be set (use c flag)")
	}
	if len(privateKey) == 0 {
		panic("expected private key to be set (use -p flag)")
	}

	// TODO: respect context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load the chain configs.
	loadCfg, err := getLoadConfig(loadConfigPath)
	if err != nil {
		panic(err)
	}

	// Connect to OmniRPC.
	omniRPCClient := omniClient.NewOmnirpcClient(loadCfg.OmniRPCUrl, metrics.NewNullHandler(), omniClient.WithCaptureReqRes())

	pingPongAddr := loadCfg.Chains[loadCfg.SummitDomainID].MessageAddr
	pingPongParser, err = pingpongclient.NewPingPongClientFilterer(common.HexToAddress(pingPongAddr), nil)
	if err != nil {
		panic(err)
	}

	originAddr := loadCfg.Chains[loadCfg.SummitDomainID].OriginAddr
	originParser, err = origin.NewParser(common.HexToAddress(originAddr))
	if err != nil {
		panic(err)
	}

	destinationAddr := loadCfg.Chains[loadCfg.SummitDomainID].DestinationAddr
	destinationParser, err = destination.NewDestinationFilterer(common.HexToAddress(destinationAddr), nil)
	if err != nil {
		panic(err)
	}

	// Load the private key.
	signer, err := getSigner(privateKey)
	if err != nil {
		panic(err)
	}

	// Get routes.
	routes, err := getMessageRoutes(loadCfg.Chains, numRoutes)
	if err != nil {
		panic(err)
	}
	fmt.Println("Routes:")
	for _, route := range routes {
		fmt.Printf("--- %d -> %d\n", route[0], route[1])
	}

	// Connect to contracts.
	messageContracts := map[int]domains.PingPongClientContract{}
	destinationContracts := map[int]domains.DestinationContract{}
	for cid, c := range loadCfg.Chains {
		// chainCfg := c
		chainID := cid
		chainClient, err := omniRPCClient.GetChainClient(ctx, chainID)
		if err != nil {
			panic(err)
		}

		messageContracts[chainID], err = evm.NewPingPongClientContract(ctx, chainClient, common.HexToAddress(c.MessageAddr))
		if err != nil {
			panic(err)
		}

		cClient, err := ethergoChain.NewFromURL(ctx, omniRPCClient.GetEndpoint(chainID, 1))
		destinationContracts[chainID], err = evm.NewDestinationContract(ctx, cClient, common.HexToAddress(c.DestinationAddr))
		if err != nil {
			panic(err)
		}
	}

	validateAgents := func() {
		// Check agent balances.
		ok, err := checkAgentBalances(ctx, omniRPCClient, loadCfg.Chains)
		if err != nil {
			panic(err)
		}
		if !ok {
			panic("agent balances are below threshold")
		}

		// Check agent statuses.
		ok, err = checkAgentStatuses(ctx, omniRPCClient, loadCfg)
		if err != nil {
			panic(err)
		}
		if !ok {
			panic("agent statuses are not active")
		}
	}

	runRoutes := func() {
		g, _ := errgroup.WithContext(ctx)
		// Send messages.

		fmt.Printf("Running %d iterations.\n\n", numIters)
		for i := 0; i < numIters; i++ {
			for _, r := range routes {
				route := r
				g.Go(func() error {
					origin := route[0]
					destPingPongAddr := common.HexToAddress(loadCfg.Chains[route[1]].MessageAddr)
					contract, ok := messageContracts[origin]
					if !ok {
						panic(fmt.Errorf("could not get contract for chain %d", origin))
					}
					var tx *ethTypes.Transaction
					err = retry.WithBackoff(ctx, func(context.Context) error {
						tx, err = contract.DoPing(ctx, signer, uint32(route[1]), destPingPongAddr, 0)
						if err != nil {
							fmt.Printf("Error doing ping: %v [chain=%d]\n", err, origin)
						}
						return err
					}, retry.WithMaxTotalTime(120*time.Second))
					if err != nil {
						fmt.Printf("Error sending message: %v\n", err)
						return err
					}
					fmt.Printf("Sent message from %d to %d: %s\n", route[0], route[1], types.GetTxLink(uint32(route[0]), tx))
					sentTxes.Store(tx.Hash(), true)

					chainClient, err := omniRPCClient.GetChainClient(ctx, int(route[0]))
					if err != nil {
						panic(err)
					}

					var receipt *ethTypes.Receipt
					var rcptErr error
					err = retry.WithBackoff(ctx, func(context.Context) error {
						receipt, rcptErr = chainClient.TransactionReceipt(ctx, tx.Hash())
						return rcptErr
					}, retry.WithMaxTotalTime(executionTimeout*time.Second))
					if err != nil {
						fmt.Printf("error getting transaction receipt: %v: %v [chain=%d, txHash=%s]\n", err, rcptErr, origin, tx.Hash())
						return err
					}
					if receipt.Status != ethTypes.ReceiptStatusSuccessful {
						fmt.Printf("status not successful: %v\n", receipt.Status)
						return fmt.Errorf("receipt status is not successful: %v", receipt.Status)
					}
					for _, log := range receipt.Logs {
						fmt.Printf("Passing log from %d to handleLog with txHash %s.\n", origin, tx.Hash())
						handleLog(log, uint32(origin))
					}
					return nil
				})
				time.Sleep(250 * time.Millisecond)
			}
		}

		unexecutedMsgs := map[common.Hash]types.Message{}
		g.Go(func() error {
			startTime := time.Now()
			numRoutesActual := len(routes)
			expectedNumExecuted := 1 * numRoutesActual * numIters
			numExecuted := 0
			executedMap := map[common.Hash]bool{}
			for {
				messages.Range(func(key, value interface{}) bool {
					leaf := key.(common.Hash)
					message := value.(types.Message)
					_, ok := executedMap[leaf]
					if !ok {
						contract, ok := destinationContracts[int(message.DestinationDomain())]
						if !ok {
							panic(fmt.Errorf("no destination contract found for chain: %d", message.DestinationDomain()))
						}
						status, err := contract.MessageStatus(ctx, message)
						if err != nil {
							fmt.Printf("error getting message status [leaf=%s]: %v\n", leaf, err)
							return true
						}
						if status == 2 {
							executedMap[leaf] = true
							numExecuted++
							totalExecuted++
							fmt.Printf("Verified message %s was executed. [total=%d]\n", leaf, numExecuted)
						}
					}
					if numExecuted >= expectedNumExecuted {
						fmt.Printf("Processed %d iterations and %d routes.\n", numIters, numRoutesActual)
						return false
					}
					time.Sleep(2 * time.Second)
					return true
				})
				if numExecuted >= expectedNumExecuted {
					return nil
				}
				elapsed := time.Since(startTime).Seconds()
				if elapsed > executionTimeout {
					messages.Range(func(key, value interface{}) bool {
						_, ok := executedMap[key.(common.Hash)]
						if !ok {
							unexecutedMsgs[key.(common.Hash)] = value.(types.Message)
						}
						return true
					})
					return fmt.Errorf("Timed out waiting for messages to be executed: %f", elapsed)
				}
			}
		})

		err = g.Wait()
		if err != nil {
			fmt.Printf("%v\n", err)
		}

		// Clear caches
		messages = &sync.Map{}
		sentTxes = &sync.Map{}
		unexecutedByChainID := map[uint32][]string{}
		for leaf, msg := range unexecutedMsgs {
			messages.Store(leaf, msg)
			_, ok := unexecutedByChainID[msg.DestinationDomain()]
			if !ok {
				unexecutedByChainID[msg.DestinationDomain()] = []string{}
			}
			unexecutedByChainID[msg.DestinationDomain()] = append(unexecutedByChainID[msg.DestinationDomain()], leaf.String())
		}
		for chainID, leafs := range unexecutedByChainID {
			fmt.Printf("Unexecuted messages on destination chain %d: %v\n", chainID, leafs)
		}
	}

	interval := time.Duration(executeInterval) * time.Second
	for {
		validateAgents()

		start := time.Now()
		runRoutes()
		elapsed := time.Since(start)

		// Calculate remaining time to wait
		waitTime := interval - elapsed
		if waitTime < 0 {
			waitTime = 0
			fmt.Printf("Completed routes [total_executed=%d].\n", totalExecuted)
		} else {
			fmt.Printf("Completed routes [total_executed=%d]. Waiting for %f seconds...\n", totalExecuted, waitTime.Seconds())
		}

		// Non-blocking wait
		select {
		case <-time.After(waitTime):
			// Continue to the next iteration
		}
	}
}
