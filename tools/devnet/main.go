package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	execConfig "github.com/synapsecns/sanguine/agents/config/executor"
	"github.com/synapsecns/sanguine/agents/contracts/test/pingpongclient"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	omniConfig "github.com/synapsecns/sanguine/services/omnirpc/config"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v2"
)

type chainConfig struct {
	Command     interface{} `yaml:"command"`
	Name        string
	ChainID     uint32
	Deployments map[string]deploymentConfig
}

func (c chainConfig) getChainClient(port uint16) (chainClient chain.Chain, err error) {
	chainURL := fmt.Sprintf("http://localhost:%d/rpc/%d", port, c.ChainID)
	return chain.NewFromURL(context.Background(), chainURL)
}

type dockerComposeConfig struct {
	Services map[string]chainConfig `yaml:"services"`
}

type deploymentConfig struct {
	Name            string
	ContractAddress string
	Contract        interface{}
}

func (d *deploymentConfig) loadContract(chainClient chain.Chain) (err error) {
	// TODO: respect context
	ctx := context.Background()
	switch d.Name {
	case "PingPongClient":
		d.Contract, err = evm.NewPingPongClientContract(ctx, chainClient, common.HexToAddress(d.ContractAddress))
		if err != nil {
			return err
		}
	default:
		err = fmt.Errorf("unknown contract %s", d.Name)
	}
	return err
}

func getChainConfigs(dockerPath string) (chainConfigs map[uint32]chainConfig, err error) {
	// Read the Docker Compose YAML file.
	dockerComposePath := fmt.Sprintf("%s/%s", dockerPath, dockerComposeFile)
	fmt.Printf("dockerComposePath: %v\n", dockerComposePath)
	data, err := os.ReadFile(dockerComposePath)
	if err != nil {
		return chainConfigs, err
	}

	// Parse the YAML data into a dockerComposeConfig struct.
	var dockerComposeConfig dockerComposeConfig
	err = yaml.Unmarshal(data, &dockerComposeConfig)
	if err != nil {
		return chainConfigs, err
	}

	chainConfigs = map[uint32]chainConfig{}
	for name, chainCfg := range dockerComposeConfig.Services {
		commandStr, ok := chainCfg.Command.(string)
		if !ok {
			continue
		}
		chainCfg.ChainID = extractChainID(commandStr)
		if chainCfg.ChainID > 0 {
			chainCfg.Name = name
			chainConfigs[chainCfg.ChainID] = chainCfg
			chainCfg.Deployments = map[string]deploymentConfig{}
		}
	}
	return chainConfigs, err
}

func extractChainID(command string) (chainID uint32) {
	re := regexp.MustCompile(`--chain-id=(\d+)`)
	matches := re.FindStringSubmatch(command)
	if len(matches) == 2 {
		chainID, _ := strconv.Atoi(matches[1])
		return uint32(chainID)
	}
	return 0
}

func loadOmniRPCConfig(dockerPath string) (omniRPCConfig omniConfig.Config, err error) {
	omniRPCPath := fmt.Sprintf("%s/config/%s", dockerPath, omnirpcConfig)
	data, err := os.ReadFile(omniRPCPath)
	if err != nil {
		return
	}
	return omniConfig.UnmarshallConfig(data)
}

func loadDeployments(contractName, deploymentPath string, chainConfigs map[uint32]chainConfig, omniRPCConfig omniConfig.Config) (err error) {
	for chainID, chainConfig := range chainConfigs {
		fmt.Printf("Loading deployment for chain %d: %v\n", chainID, chainConfig.Name)
		contractABIPath := fmt.Sprintf("%s/%s/%s.json", deploymentPath, chainConfig.Name, contractName)
		abiFile, err := os.Open(contractABIPath)
		if err != nil {
			return err
		}
		defer abiFile.Close()

		abiBytes, err := io.ReadAll(abiFile)
		if err != nil {
			return err
		}

		var data map[string]interface{}
		err = json.Unmarshal(abiBytes, &data)
		if err != nil {
			return err
		}

		deployment := deploymentConfig{Name: contractName}
		var ok bool
		deployment.ContractAddress, ok = data["address"].(string)
		if !ok {
			return fmt.Errorf("could not find address for %s", contractName)
		}

		chainClient, err := chainConfig.getChainClient(omniRPCConfig.Port)
		if err != nil {
			return err
		}
		err = deployment.loadContract(chainClient)
		if err != nil {
			return err
		}

		if chainConfig.Deployments == nil {
			chainConfig.Deployments = map[string]deploymentConfig{}
		}
		chainConfig.Deployments[contractName] = deployment
		chainConfigs[chainID] = chainConfig
	}
	return nil
}

func loadParser(addr common.Address) (parser *pingpongclient.PingPongClientFilterer, err error) {
	return pingpongclient.NewPingPongClientFilterer(addr, nil)
}

func getSummitChainID(dockerPath string) (summitChainID uint32, err error) {
	executorPath := fmt.Sprintf("%s/config/%s", dockerPath, executorConfig)
	executorConfig, err := execConfig.DecodeConfig(executorPath)
	if err != nil {
		return 0, err
	}
	return executorConfig.SummitChainID, nil
}

func getSigner(privateKey string) (signer *localsigner.Signer, err error) {
	localWallet, err := wallet.FromHex(privateKey)
	if err != nil {
		return
	}
	signer = localsigner.NewSigner(localWallet.PrivateKey())
	return signer, nil
}

func getMessageRoutes(chainConfigs map[uint32]chainConfig, summitChainID uint32, numRoutes int) (routes [][2]uint32, err error) {
	chainIDs := []uint32{}
	for chainID := range chainConfigs {
		if chainID == summitChainID {
			continue
		}
		chainIDs = append(chainIDs, chainID)
	}
	fmt.Printf("got chain IDs: %v\n", chainIDs)
	routes = [][2]uint32{}
	for i, chainID := range chainIDs {
		if len(routes) >= numRoutes {
			return routes, nil
		}
		origin := chainID
		destination := chainIDs[(i+1)%len(chainIDs)]
		if origin == destination {
			return nil, fmt.Errorf("cannot generate unique origin / destination pair")
		}
		routes = append(routes, [2]uint32{origin, destination})
	}
	return routes, nil
}

const scribeConnectTimeout = 30 * time.Second

func makeScribeClient(parentCtx context.Context, handler metrics.Handler, url string) (*grpc.ClientConn, pbscribe.ScribeServiceClient, error) {
	fmt.Printf("make scribe client with url: %v\n", url)
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

// streamLogs uses the grpcConnection to Scribe, with a chainID and address to get all logs from that address.
func streamLogs(ctx context.Context, chainID uint32, address string, conn pbscribe.ScribeServiceClient, omniRPCClient omniClient.RPCClient) error {
	fmt.Printf("streaming logs for chain %d on addr %s\n", chainID, address)
	// chainClient, err := omniRPCClient.GetChainClient(ctx, int(chainID))
	// if err != nil {
	// 	return err
	// }
	// fromBlock, err := chainClient.BlockNumber(ctx)
	// if err != nil {
	// 	return err
	// }
	fromBlock := 0
	fmt.Printf("Got current block number for chain %d: %d\n", chainID, fromBlock)
	toBlock := "latest"
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
	var event interface{}
	if event, err = parser.ParsePingSent(*log); err == nil {
		pingSentEvent, ok := event.(*pingpongclient.PingPongClientPingSent)
		if !ok {
			return fmt.Errorf("could not parse ping sent event")
		}
		fmt.Printf("Parsed ping sent on chain %d [ID=%d]\n", chainID, pingSentEvent.PingId.Int64())
	} else if event, err = parser.ParsePingReceived(*log); err == nil {
		pingReceivedEvent, ok := event.(*pingpongclient.PingPongClientPingReceived)
		if !ok {
			return fmt.Errorf("could not parse ping received event")
		}
		fmt.Printf("Parsed ping received on chain %d [ID=%d]\n", chainID, pingReceivedEvent.PingId.Int64())
	} else if event, err = parser.ParsePongSent(*log); err == nil {
		pongSentEvent, ok := event.(*pingpongclient.PingPongClientPongSent)
		if !ok {
			return fmt.Errorf("could not parse pong sent event")
		}
		fmt.Printf("Parsed pong sent on chain %d [ID=%d]\n", chainID, pongSentEvent.PingId.Int64())
	} else if event, err = parser.ParsePongReceived(*log); err == nil {
		pongReceivedEvent, ok := event.(*pingpongclient.PingPongClientPongReceived)
		if !ok {
			return fmt.Errorf("could not parse pong received event")
		}
		fmt.Printf("Parsed pong received on chain %d [ID=%d]\n", chainID, pongReceivedEvent.PingId.Int64())
	} else {
		return fmt.Errorf("could not parse log")
	}
	return nil
}

var dockerComposeFile = "docker-compose.yml"
var omnirpcConfig = "omnirpc.yaml"
var executorConfig = "executor-config.yml"
var parser *pingpongclient.PingPongClientFilterer

const eventBufferSize = 1000

func main() {
	var dockerPath string
	var deploymentPath string
	var privateKey string
	var numIters int
	flag.StringVar(&dockerPath, "d", "", "path to devnet docker files")
	flag.StringVar(&deploymentPath, "c", "", "path to contract deployments")
	flag.StringVar(&privateKey, "p", "", "private key")
	flag.IntVar(&numIters, "n", 1, "number of message route iterations")
	flag.Parse()
	if len(dockerPath) == 0 {
		panic("expected docker path to be set (use -d flag)")
	}
	if len(deploymentPath) == 0 {
		panic("expected deployment path to be set (use -c flag)")
	}
	if len(privateKey) == 0 {
		panic("expected private key to be set (use -p flag)")
	}

	// TODO: respect context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load the chain configs.
	chainConfigs, err := getChainConfigs(dockerPath)
	if err != nil {
		panic(err)
	}

	// Load the omnirpc config.
	omniRPCConfig, err := loadOmniRPCConfig(dockerPath)
	if err != nil {
		panic(err)
	}

	// Load the deployments.
	err = loadDeployments("PingPongClient", deploymentPath, chainConfigs, omniRPCConfig)
	if err != nil {
		panic(err)
	}

	// Load the summit chain id.
	summitChainID, err := getSummitChainID(dockerPath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("summitChainID: %v\n", summitChainID)

	pingPongAddr := chainConfigs[summitChainID].Deployments["PingPongClient"].ContractAddress
	parser, err = pingpongclient.NewPingPongClientFilterer(common.HexToAddress(pingPongAddr), nil)
	if err != nil {
		panic(err)
	}

	// Load the private key.
	signer, err := getSigner(privateKey)
	if err != nil {
		panic(err)
	}

	// Get routes.
	routes, err := getMessageRoutes(chainConfigs, summitChainID, 2)
	if err != nil {
		panic(err)
	}
	// routes = [][2]uint32{{43, 44}, {44, 43}}
	fmt.Printf("routes: %v\n", routes)

	// Connect to Scribe.
	// omniRPCURL := fmt.Sprintf("localhost:%d", omniRPCConfig.Port)
	omniRPCURL := "localhost:9002"
	_, scribeClient, err := makeScribeClient(ctx, metrics.NewNullHandler(), omniRPCURL)
	if err != nil {
		panic(err)
	}

	// Connect to OmniRPC.
	omniRPCClient := omniClient.NewOmnirpcClient("http://localhost:9001", metrics.NewNullHandler(), omniClient.WithCaptureReqRes())

	// Listen for messages.
	g, ctx := errgroup.WithContext(ctx)
	contractName := "PingPongClient"
	for _, c := range chainConfigs {
		chainCfg := c
		addr := chainCfg.Deployments[contractName].ContractAddress
		g.Go(func() error {
			return streamLogs(ctx, chainCfg.ChainID, addr, scribeClient, omniRPCClient)
		})
	}

	// Send messages.
	fmt.Printf("Running %d iterations.\n", numIters)
	for i := 0; i < numIters; i++ {
		for _, route := range routes {
			fmt.Printf("Sending message from %d to %d\n", route[0], route[1])
			contract, ok := chainConfigs[route[0]].Deployments[contractName].Contract.(domains.PingPongClientContract)
			if !ok {
				panic("could not cast contract")
			}
			destPingPongAddr := common.HexToAddress(chainConfigs[route[1]].Deployments[contractName].ContractAddress)
			tx, err := contract.DoPing(ctx, signer, route[1], destPingPongAddr, 0)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Sent ping to contract %s: %s\n", destPingPongAddr.String(), tx.Hash().String())
		}
	}

	err = g.Wait()
	if err != nil {
		panic(err)
	}
}
