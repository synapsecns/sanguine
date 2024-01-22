package main

import (
	"encoding/hex"
	"flag"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
)

func main() {
	var rawRequest string
	flag.StringVar(&rawRequest, "r", "", "raw request")
	flag.Parse()
	if rawRequest == "" {
		panic("must provide raw request (use -r)")
	}
	if rawRequest[:2] == "0x" {
		rawRequest = rawRequest[2:]
	}
	requestBytes, err := hex.DecodeString(rawRequest)
	if err != nil {
		panic(err)
	}

	parser, err := fastbridge.NewParser(common.HexToAddress(""))
	if err != nil {
		panic(err)
	}

	log := types.Log{
		Topics: []common.Hash{
			common.HexToHash("0x120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a"),
			common.HexToHash("0xb7439e36b5527ac6298c2fd035a286d9df33c5352d96f08c48d4bf06f9df4afd"),
			common.HexToHash("0x0000000000000000000000005cf2cc2c71231c23cd5c5a008b9339da33f0fa57"),
		},
		Data: requestBytes,
	}
	_, parsedEvent, ok := parser.ParseEvent(log)
	if !ok {
		panic("could not parse event")
	}

	switch event := parsedEvent.(type) {
	case *fastbridge.FastBridgeBridgeRequested:
		handleBridgeRequested(event)
	default:
		panic("unknown event")
	}
}

func handleBridgeRequested(event *fastbridge.FastBridgeBridgeRequested) {
	fmt.Println("BridgeRequested:")
	fmt.Printf("TransactionID: %s\n", hexutil.Encode(event.TransactionId[:]))
	fmt.Printf("Sender: %s\n", event.Sender.String())
	fmt.Printf("OriginAmount: %s\n", event.OriginAmount.String())
	fmt.Printf("DestAmount: %s\n", event.DestAmount.String())
	fmt.Printf("OriginToken: %s\n", event.OriginToken.String())
	fmt.Printf("DestToken: %s\n", event.DestToken.String())
	fmt.Printf("SendChainGas: %v\n", event.SendChainGas)
}
