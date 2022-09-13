package bridge_test

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/synapse-node/contracts/bridge"
	"github.com/synapsecns/synapse-node/pkg/types"
	"github.com/synapsecns/synapse-node/testutils"
	"github.com/synapsecns/synapse-node/testutils/utils/events"
)

type EvmCrossChainEventLog interface {
	types.CrossChainEventLog
	// GetRaw gets the raw event logs
	GetRaw() ethTypes.Log
}

func (b BridgeSuite) TestEventType() {
	mockBridge := events.NewMockBridge(b.GetTestContext(), b.T())
	// mock some events
	mockBridge.MockEvents(b.GetTestContext(), events.AllEventTypes()...)

	// all events should be able to be parsed
	for _, triggeredEvents := range mockBridge.GetAllEvents() {
		for _, event := range triggeredEvents {
			castEvent, ok := event.(EvmCrossChainEventLog)
			True(b.T(), ok)

			// get the parser for this chain
			testParser := b.getParser(mockBridge, castEvent)

			// get the parsed event, make sure it matches the event type
			parsedEvent, ok := testParser.EventType(castEvent.GetRaw())
			True(b.T(), ok)

			Equal(b.T(), parsedEvent, castEvent.GetEventType())

			if parsedEvent.BridgeInitiated() {
				bridgeLog, err := testParser.GetCrossChainBridgeEvent(castEvent.GetRaw())
				Nil(b.T(), err)
				NotNil(b.T(), bridgeLog)
				Equal(b.T(), bridgeLog.GetEventType(), parsedEvent)
			} else {
				userLog, err := testParser.GetCrossChainUserEvent(castEvent.GetRaw())
				Nil(b.T(), err)
				NotNil(b.T(), userLog)
				Equal(b.T(), userLog.GetEventType(), parsedEvent)
			}
		}
	}
}

// getParser gets a parser from a mock bridge by event by chain
// it does this by checking the tx hash on both chains and making a parser for the chain in question.
func (b BridgeSuite) getParser(mockBridge *events.MockBridge, event types.CrossChainEventLog) (parser bridge.Parser) {
	b.T().Helper()

	var backend backends.SimulatedTestBackend
	wg := testutils.NewWaitGroup()
	wg.Go(func() {
		tx, _, _ := mockBridge.UnderlyingBackend.TransactionByHash(b.GetTestContext(), common.HexToHash(event.GetIdentifier()))
		if tx != nil {
			backend = mockBridge.UnderlyingBackend
		}
	})
	wg.Go(func() {
		tx, _, _ := mockBridge.SynBackend.TransactionByHash(b.GetTestContext(), common.HexToHash(event.GetIdentifier()))
		if tx != nil {
			backend = mockBridge.SynBackend
		}
	})
	wg.Wait()

	// make sure the parser isn't stil nil
	NotNil(b.T(), backend)

	// create the parser
	parser, err := bridge.NewParser(backend.Config().GetEthBridgeAddress())
	Nil(b.T(), err)

	return parser
}
