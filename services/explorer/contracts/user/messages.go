// Package user defines all end-user defined messages we resolve to in the explorer
package user

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated/multibackend"
	"github.com/synapsecns/sanguine/services/explorer/contracts/user/dfk/dfkhero"
	"github.com/synapsecns/sanguine/services/explorer/contracts/user/dfk/dfkpet"
	"github.com/synapsecns/sanguine/services/explorer/contracts/user/dfk/dfktear"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"math/big"
	"sync"
)

// globMux is a global mutex to protect the global message parsing map
// to parse messages, we start up a simulated network and add the contracts to it
// to avoid slower start time, we do this in a go function at boot time.
// we lock the mutex here and then release it only acquiring read locks going forward since these are much cheaper.
var globMux sync.RWMutex

var decoders []Decoder

// Decoder decodes a message into a message type.
type Decoder func(ctx context.Context, message []byte) (model.MessageType, error)

func init() {
	go func() {
		err := deployParseNet()
		if err != nil {
			// TODO; is this such a good idea?
			panic(err)
		}
	}()
}

// Decode decodes message hex data.
func Decode(ctx context.Context, hexMessage string) model.MessageType {
	globMux.RLock()
	defer globMux.RUnlock()

	message := common.FromHex(hexMessage)

	if len(hexMessage) == 0 {
		return model.UnknownType{Known: false}
	}

	results := make([]model.MessageType, len(decoders))

	var wg sync.WaitGroup
	var sliceMux sync.Mutex
	for i, decoder := range decoders {
		wg.Add(1)
		i := i
		decoder := decoder // capture func literal
		go func() {
			defer wg.Done()
			messageType, err := decoder(ctx, message)
			if err != nil {
				return
			}
			if messageType != nil {
				sliceMux.Lock()
				results[i] = messageType
				sliceMux.Unlock()
			}
		}()
	}

	wg.Wait()

	// go in order of decoders from largest->smallest
	for _, result := range results {
		if result != nil {
			return result
		}
	}

	return model.UnknownType{Known: false}
}

func deployParseNet() error {
	globMux.Lock()
	defer globMux.Unlock()

	// 100 million ether
	balance := big.NewInt(0).Mul(big.NewInt(params.Ether), big.NewInt(100000000))

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return fmt.Errorf("failed to generate private key: %w", err)
	}

	chainID := params.AllEthashProtocolChanges.ChainID
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create authorized transactor: %w", err)
	}

	genesisAlloc := map[common.Address]core.GenesisAccount{
		auth.From: {
			Balance: balance,
		},
	}

	simulatedBackend := multibackend.NewSimulatedBackendWithConfig(genesisAlloc, simulated.BlockGasLimit, params.AllEthashProtocolChanges)

	_, _, heroBridge, err := dfkhero.DeployHeroBridgeUpgradeable(auth, simulatedBackend)
	if err != nil {
		return fmt.Errorf("could not deploy hero bridge: %w", err)
	}

	// Note: these should stay in order of largest to smallest
	decoders = append(decoders, func(ctx context.Context, message []byte) (model.MessageType, error) {
		messageFormat, err := heroBridge.DecodeMessage(&bind.CallOpts{Context: ctx}, message)
		if err != nil {
			return nil, fmt.Errorf("could not decode message: %w", err)
		}
		return model.HeroType{
			HeroID: messageFormat.DstHeroId.String(),
		}, nil
	})

	_, _, petBridge, err := dfkpet.DeployPetBridgeUpgradeable(auth, simulatedBackend)
	if err != nil {
		return fmt.Errorf("could not deploy hero bridge: %w", err)
	}

	decoders = append(decoders, func(ctx context.Context, message []byte) (model.MessageType, error) {
		messageFormat, err := petBridge.DecodeMessage(&bind.CallOpts{Context: ctx}, message)
		if err != nil {
			return nil, fmt.Errorf("could not decode message: %w", err)
		}
		return model.PetType{
			PetID: messageFormat.DstPet.Id.String(),
			Name:  messageFormat.DstPet.Name,
		}, nil
	})

	_, _, tearBridge, err := dfktear.DeployTearBridge(auth, simulatedBackend, common.Address{}, common.Address{})
	if err != nil {
		return fmt.Errorf("could not deploy hero bridge: %w", err)
	}

	decoders = append(decoders, func(ctx context.Context, message []byte) (model.MessageType, error) {
		messageFormat, err := tearBridge.DecodeMessage(&bind.CallOpts{Context: ctx}, message)
		if err != nil {
			return nil, fmt.Errorf("could not decode message: %w", err)
		}
		return model.TearType{
			Amount: messageFormat.DstTearAmount.String(),
		}, nil
	})

	simulatedBackend.Commit()

	return nil
}
