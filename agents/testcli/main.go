package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
)

func main() {
	fmt.Println("Enter Sending Chain URL: ")
	var sendingChainURL string
	fmt.Scanln(&sendingChainURL)

	ctx := context.Background()
	underlyingClient, err := chain.NewFromURL(context.Background(), sendingChainURL)
	if err != nil {
		fmt.Printf("could not get evm: %s", err)
		return
	}

	fmt.Println("Enter Sending Test Client Contract Address: ")
	var sendingTestClientContract string
	fmt.Scanln(&sendingTestClientContract)

	boundTestClient, err := evm.NewTestClientContract(ctx, underlyingClient, common.HexToAddress(sendingTestClientContract))
	if err != nil {
		fmt.Printf("could not bind test client contract: %s", err)
		return
	}

	fmt.Println("Enter Signing Private Key: ")
	var signingPrivateKey string
	fmt.Scanln(&signingPrivateKey)

	localWallet, err := wallet.FromHex(signingPrivateKey)
	if err != nil {
		fmt.Printf("could not generate local wallet: %s", err)
		return
	}
	localSigner := localsigner.NewSigner(localWallet.PrivateKey())

	fmt.Println("Enter Receiving Test Client Contract Address: ")
	var receivingTestClientAddress string
	fmt.Scanln(&receivingTestClientAddress)

	recipient := common.HexToAddress(receivingTestClientAddress)

	fmt.Println("Enter Destination Domain ID: ")
	var domainIDStr string
	fmt.Scanln(&domainIDStr)

	destID64, err := strconv.ParseUint(domainIDStr, 10, 32)
	if err != nil {
		fmt.Printf("could not parse destination id: %s", err)
		return
	}

	destinationID := uint32(destID64)

	fmt.Println("Enter Optimistic Seconds: ")
	var optimisticSecondsStr string
	fmt.Scanln(&optimisticSecondsStr)

	optSec64, err := strconv.ParseUint(optimisticSecondsStr, 10, 32)
	if err != nil {
		fmt.Printf("could not parse optimistic seconds: %s", err)
		return
	}

	optimisticSeconds := uint32(optSec64)

	fmt.Println("Enter Message to Send: ")
	var message string
	fmt.Scanln(&message)
	messageBody := []byte(message)

	err = boundTestClient.SendMessage(ctx, localSigner, destinationID, recipient, optimisticSeconds, messageBody)
	if err != nil {
		fmt.Printf("could not send message: %s", err)
		return
	}
}
