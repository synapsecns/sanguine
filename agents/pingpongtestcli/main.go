// Package main has the main file for the pingpongtestcli utility
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

//nolint:gosec,cyclop
func main() {
	fmt.Println("Enter Sending Chain URL (eg https://polygon-rpc.com, https://api.avax.network/ext/bc/C/rpc, https://optimism-mainnet.public.blastapi.io): ")
	var sendingChainURL string
	fmt.Scanln(&sendingChainURL)

	ctx := context.Background()
	underlyingClient, err := chain.NewFromURL(context.Background(), sendingChainURL)
	if err != nil {
		fmt.Printf("could not get evm: %s", err)
		return
	}

	fmt.Println("Enter Sending Test Client Contract Address (eg 0x07303feddAd86BF1ac260F1d9886E420D9c7144C): ")
	var sendingTestClientContract string
	fmt.Scanln(&sendingTestClientContract)

	boundPingPongClient, err := evm.NewPingPongClientContract(ctx, underlyingClient, common.HexToAddress(sendingTestClientContract))
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

	fmt.Println("Enter Receiving Test Client Contract Address (eg 0x07303feddAd86BF1ac260F1d9886E420D9c7144C): ")
	var receivingPingPongClientAddress string
	fmt.Scanln(&receivingPingPongClientAddress)

	recipient := common.HexToAddress(receivingPingPongClientAddress)

	fmt.Println("Enter Destination Domain ID (eg. polygon: 137, avalanche 43114, optimism: 10): ")
	var domainIDStr string
	fmt.Scanln(&domainIDStr)

	destID64, err := strconv.ParseUint(domainIDStr, 10, 32)
	if err != nil {
		fmt.Printf("could not parse destination id: %s", err)
		return
	}

	destinationID := uint32(destID64)

	err = boundPingPongClient.DoPing(ctx, localSigner, destinationID, recipient, uint16(1))
	if err != nil {
		fmt.Printf("could not send ping: %s", err)
		return
	}
}
