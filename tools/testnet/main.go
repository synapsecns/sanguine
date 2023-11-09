package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
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
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v2"
)

type loadConfig struct {
	SummitDomainID int                 `yaml:"summit_domain_id"`
	OmniRPCUrl     string              `yaml:"omnirpc_url"`
	ScribeURL      string              `yaml:"scribe_url"`
	Chains         map[int]chainConfig `yaml:"chains"`
}

type chainConfig struct {
	MessageAddr     string `yaml:"message_contract_addr"`
	OriginAddr      string `yaml:"origin_addr"`
	DestinationAddr string `yaml:"destination_addr"`
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

const scribeConnectTimeout = 30 * time.Second

func makeScribeClient(parentCtx context.Context, handler metrics.Handler, url string) (*grpc.ClientConn, pbscribe.ScribeServiceClient, error) {
	fmt.Printf("Connecting to scribe with URL %v...\n", url)
	ctx, cancel := context.WithTimeout(parentCtx, scribeConnectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor(otelgrpc.WithTracerProvider(handler.GetTracerProvider()))),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor(otelgrpc.WithTracerProvider(handler.GetTracerProvider()))),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("could not dial grpc: %w", err)
	}

	scribeClient := pbscribe.NewScribeServiceClient(conn)

	// Ensure that gRPC is up and running.
	healthCheck, err := scribeClient.Check(ctx, &pbscribe.HealthCheckRequest{}, grpc.WaitForReady(true))
	if err != nil {
		return nil, nil, fmt.Errorf("could not check: %w", err)
	}
	if healthCheck.Status != pbscribe.HealthCheckResponse_SERVING {
		return nil, nil, fmt.Errorf("not serving: %s", healthCheck.Status)
	}

	return conn, scribeClient, nil
}

var startBlocks map[int]uint64 = map[int]uint64{}

// streamLogs uses the grpcConnection to Scribe, with a chainID and address to get all logs from that address.
func streamLogs(ctx context.Context, chainID uint32, address string, conn pbscribe.ScribeServiceClient, omniRPCClient omniClient.RPCClient) error {
	chainClient, err := omniRPCClient.GetChainClient(ctx, int(chainID))
	if err != nil {
		return err
	}
	startBlocks[int(chainID)], err = chainClient.BlockNumber(ctx)
	if err != nil {
		return err
	}
	fromBlock := 0
	toBlock := "latest"
	fmt.Printf("Streaming logs for chain %d on addr %s from %v to %v.\n", chainID, address, fromBlock, toBlock)
	stream, err := conn.StreamLogs(ctx, &pbscribe.StreamLogsRequest{
		Filter: &pbscribe.LogFilter{
			ContractAddress: &pbscribe.NullableString{Kind: &pbscribe.NullableString_Data{Data: address}},
			ChainId:         chainID,
		},
		FromBlock: strconv.Itoa(int(fromBlock)),
		ToBlock:   toBlock,
	})
	if err != nil {
		fmt.Println("could not stream")
		return fmt.Errorf("could not stream logs: %w", err)
	}

	for {
		response, err := stream.Recv()
		if err != nil {
			return fmt.Errorf("could not receive: %w", err)
		}

		log := response.Log.ToLog()
		if log == nil {
			return fmt.Errorf("could not convert log")
		}

		select {
		case <-ctx.Done():
			fmt.Println("context done")
			err := stream.CloseSend()
			if err != nil {
				return fmt.Errorf("could not close stream: %w", err)
			}

			return fmt.Errorf("context done: %w", ctx.Err())
		default:
			err = handleLog(log, chainID)
			if err != nil {
				return err
			}
		}
	}
}

func handleLog(log *ethTypes.Log, chainID uint32) (err error) {
	// drop logs that are before the start block for this chain
	// startBlock, ok := startBlocks[int(chainID)]
	// if !ok {
	// 	return fmt.Errorf("could not get start block for chain %d", chainID)
	// }
	// if log.BlockNumber < startBlock {
	// 	return nil
	// }

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

var pingPongParser *pingpongclient.PingPongClientFilterer
var originParser origin.Parser
var destinationParser *destination.DestinationFilterer
var numIters, numExecuted, numRoutes, numSent int
var messages = &sync.Map{}
var sentTxes = &sync.Map{}

// var messages = map[common.Hash]types.Message{}
// var sentTxes = map[common.Hash]bool{}

const eventBufferSize = 1000

func main() {
	var loadConfigPath string
	var privateKey string
	flag.StringVar(&loadConfigPath, "c", "", "path to load config")
	flag.StringVar(&privateKey, "p", "", "private key")
	flag.IntVar(&numIters, "n", 1, "number of message route iterations")
	flag.IntVar(&numIters, "r", 0, "number of routes")
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

	// // Connect to Scribe.
	// _, scribeClient, err := makeScribeClient(ctx, metrics.NewNullHandler(), loadCfg.ScribeURL)
	// if err != nil {
	// 	panic(err)
	// }

	// Connect to OmniRPC.
	omniRPCClient := omniClient.NewOmnirpcClient(loadCfg.OmniRPCUrl, metrics.NewNullHandler(), omniClient.WithCaptureReqRes())

	// Listen for messages.
	g, _ := errgroup.WithContext(ctx)
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

		// messageAddr := chainCfg.MessageAddr
		// g.Go(func() error {
		// 	return streamLogs(ctx, uint32(chainID), messageAddr, scribeClient, omniRPCClient)
		// })

		// originAddr := chainCfg.OriginAddr
		// g.Go(func() error {
		// 	return streamLogs(ctx, uint32(chainID), originAddr, scribeClient, omniRPCClient)
		// })

		// destinationAddr := chainCfg.DestinationAddr
		// g.Go(func() error {
		// 	return streamLogs(ctx, uint32(chainID), destinationAddr, scribeClient, omniRPCClient)
		// })
	}

	// Send messages.
	fmt.Printf("Running %d iterations.\n\n", numIters)
	for i := 0; i < numIters; i++ {
		for _, r := range routes {
			route := r
			g.Go(func() error {
				destPingPongAddr := common.HexToAddress(loadCfg.Chains[route[1]].MessageAddr)
				contract, ok := messageContracts[route[0]]
				if !ok {
					panic(fmt.Errorf("could not get contract for chain %d", route[0]))
				}
				var tx *ethTypes.Transaction
				err = retry.WithBackoff(ctx, func(context.Context) error {
					tx, err = contract.DoPing(ctx, signer, uint32(route[1]), destPingPongAddr, 0)
					if err != nil {
						fmt.Printf("Error doing ping: %v\n", err)
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
				origin := route[0]
				err = retry.WithBackoff(ctx, func(context.Context) error {
					receipt, err = chainClient.TransactionReceipt(ctx, tx.Hash())
					return err
				}, retry.WithMaxTotalTime(120*time.Second))
				if err != nil {
					fmt.Printf("error getting transaction receipt: %v\n", err)
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

	g.Go(func() error {
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
						fmt.Printf("Verified message %s was executed. [total=%d]\n", leaf, numExecuted)
					}
				}
				if numExecuted >= expectedNumExecuted {
					fmt.Printf("Processed %d iterations and %d routes.\n", numIters, numRoutesActual)
					cancel()
					return false
				}
				time.Sleep(2 * time.Second)
				return true
			})
			if numExecuted >= expectedNumExecuted {
				return nil
			}
		}
	})

	err = g.Wait()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
