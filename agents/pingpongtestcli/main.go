// Package main has the main file for the pingpongtestcli utility
package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/contracts/test/pingpongclient"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
)

//nolint:gosec
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

	fmt.Println("Enter Receiving Test Client Contract Address: ")
	var receivingPingPongClientAddress string
	fmt.Scanln(&receivingPingPongClientAddress)

	recipient := common.HexToAddress(receivingPingPongClientAddress)

	fmt.Println("Enter Destination Domain ID: ")
	var domainIDStr string
	fmt.Scanln(&domainIDStr)

	destID64, err := strconv.ParseUint(domainIDStr, 10, 32)
	if err != nil {
		fmt.Printf("could not parse destination id: %s", err)
		return
	}

	destinationID := uint32(destID64)

	pingSentSink := make(chan *pingpongclient.PingPongClientPingSent)
	pingSentSub, err := boundPingPongClient.WatchPingSent(ctx, pingSentSink)
	if err != nil {
		fmt.Printf("could not create channel to watch for ping sent: %s", err)
		return
	}

	pongReceivedSink := make(chan *pingpongclient.PingPongClientPongReceived)
	pongReceivedSub, err := boundPingPongClient.WatchPongReceived(ctx, pongReceivedSink)
	if err != nil {
		fmt.Printf("could not create channel to watch for pong received: %s", err)
		return
	}

	err = boundPingPongClient.DoPing(ctx, localSigner, destinationID, recipient, uint16(1))
	if err != nil {
		fmt.Printf("could not send ping: %s", err)
		return
	}

	pingSentWatchCtx, pingSentCancel := context.WithTimeout(ctx, time.Second*120)
	defer pingSentCancel()

	select {
	// check for errors and fail
	case <-pingSentWatchCtx.Done():
		fmt.Printf("ping sent context completed %v", ctx.Err())
		return
	case <-pingSentSub.Err():
		fmt.Printf("ping sent context completed %v", pingSentSub.Err())
		return
	// get dispatch event
	case pingSentItem := <-pingSentSink:
		if pingSentItem == nil {
			fmt.Printf("item from pingSentSink was nil unexpectedly?\n")
		}

		innerWatchCtx, innerCancel := context.WithTimeout(ctx, time.Second*120)
		defer innerCancel()

		select {
		// check for errors and fail
		case <-innerWatchCtx.Done():
			fmt.Printf("pong received context completed %v", ctx.Err())
			return
		case <-pongReceivedSub.Err():
			fmt.Printf("pong received context completed %v", pongReceivedSub.Err())
			return
		// get dispatch event
		case pongReceivedItem := <-pongReceivedSink:
			if pongReceivedItem == nil {
				fmt.Printf("item from pongReceivedSink was nil unexpectedly?\n")
			}
			break
		}

		break
	}
}
