package main

import (
	"flag"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/inbox"
	"github.com/synapsecns/sanguine/agents/types"
)

func main() {
	var address string
	var logTopic string
	var logDataHex string
	var agentAddr string
	var chainID int

	flag.StringVar(&address, "a", "", "address")
	flag.StringVar(&logTopic, "t", "", "log topic")
	flag.StringVar(&logDataHex, "d", "", "log data (hex)")
	flag.StringVar(&agentAddr, "agent", "", "agent address")
	flag.IntVar(&chainID, "c", 0, "chain ID")
	flag.Parse()
	if len(address) == 0 {
		panic("expected address to be set (use a flag)")
	}
	if len(logTopic) == 0 {
		panic("expected log topic to be set (use t flag)")
	}
	if len(logDataHex) == 0 {
		panic("expected log data to be set (use d flag)")
	}

	log := ethTypes.Log{
		Address: common.HexToAddress(address),
		Topics:  []common.Hash{common.HexToHash(logTopic)},
		Data:    common.FromHex(logDataHex),
	}
	log.Topics = []common.Hash{
		common.HexToHash("0x5ca3d740e03650b41813a4b418830f6ba39700ae010fe8c4d1bca0e8676b9c56"),
		common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000aa36a7"),
		common.HexToHash("0x000000000000000000000000bfb161171246667dc713a9333a9ff8dae75703cb"),
	}

	// lightInboxParser, err := lightinbox.NewLightInboxFilterer(common.HexToAddress(address), nil)
	// if err != nil {
	// 	panic(err)
	// }

	// event, err := lightInboxParser.ParseInvalidStateWithSnapshot(log)
	// if err != nil {
	// 	panic(err)
	// }

	inboxParser, err := inbox.NewInboxFilterer(common.HexToAddress(address), nil)
	if err != nil {
		panic(err)
	}

	event, err := inboxParser.ParseSnapshotAccepted(log)
	if err != nil {
		panic(err)
	}

	snapshotData, err := types.NewSnapshotWithMetadata(event.SnapPayload, uint32(chainID), common.HexToAddress(agentAddr), event.SnapSignature)
	if err != nil {
		panic(err)
	}

	snapRoot, _, err := snapshotData.Snapshot.SnapshotRootAndProofs()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed snapshot with snapRoot: %v\n", common.BytesToHash(snapRoot[:]))

	states := snapshotData.Snapshot.States()
	for _, state := range states {
		stateRoot := state.Root()
		fmt.Printf("Got state on domain %d, nonce %d, blockNumber %d, root %s\n", state.Origin(), state.Nonce(), state.BlockNumber(), common.BytesToHash(stateRoot[:]))
	}
}
