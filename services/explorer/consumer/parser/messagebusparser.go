package parser

import (
	"context"
	"fmt"
	"time"

	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher/tokenprice"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/contracts/messagebus"
	"github.com/synapsecns/sanguine/services/explorer/db"
	model "github.com/synapsecns/sanguine/services/explorer/db/sql"
	messageBusTypes "github.com/synapsecns/sanguine/services/explorer/types/messagebus"
)

// MessageBusParser parses messagebus logs.
type MessageBusParser struct {
	// consumerDB is the database to store parsed data in
	consumerDB db.ConsumerDB
	// Filterer is the message Filterer we use to parse events
	Filterer *messagebus.MessageBusUpgradeableFilterer
	// messageAddress is the address of the message
	messageBusAddress common.Address
	// consumerFetcher is the Fetcher for sender and timestamp
	consumerFetcher fetcher.ScribeFetcher
	// tokenPriceService contains the token price service/cache
	tokenPriceService tokenprice.Service
}

// NewMessageBusParser creates a new parser for a given message.
func NewMessageBusParser(consumerDB db.ConsumerDB, messageBusAddress common.Address, consumerFetcher fetcher.ScribeFetcher, tokenPriceService tokenprice.Service) (*MessageBusParser, error) {
	filterer, err := messagebus.NewMessageBusUpgradeableFilterer(messageBusAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", messagebus.MessageBusReceiverUpgradeableFilterer{}, err)
	}
	return &MessageBusParser{consumerDB, filterer, messageBusAddress, consumerFetcher, tokenPriceService}, nil
}

// EventType returns the event type of a message log.
func (m *MessageBusParser) EventType(log ethTypes.Log) (_ messageBusTypes.EventType, ok bool) {
	for _, logTopic := range log.Topics {
		eventType := messagebus.EventTypeFromTopic(logTopic)
		if eventType == nil {
			continue
		}
		return *eventType, true
	}
	// return an unknown event to avoid cases where user failed to check the event type
	return messageBusTypes.EventType(len(messageBusTypes.AllEventTypes()) + 2), false
}

// eventToMessageEvent stores a message event.
func eventToMessageEvent(event messageBusTypes.EventLog, chainID uint32) model.MessageBusEvent {
	return model.MessageBusEvent{
		InsertTime:      uint64(time.Now().UnixNano()),
		ContractAddress: event.GetContractAddress().String(),
		ChainID:         chainID,
		EventType:       event.GetEventType().Int(),
		BlockNumber:     event.GetBlockNumber(),
		TxHash:          event.GetTxHash().String(),
		EventIndex:      event.GetEventIndex(),
		Sender:          "",
		MessageID:       ToNullString(event.GetMessageID()),
		SourceChainID:   event.GetSourceChainID(),

		Status:             ToNullString(event.GetStatus()),
		SourceAddress:      ToNullString(event.GetSourceAddress()),
		DestinationAddress: ToNullString(event.GetDestinationAddress()),
		DestinationChainID: event.GetDestinationChainID(),
		Nonce:              event.GetNonce(),
		Message:            ToNullString(event.GetMessage()),
		Receiver:           ToNullString(event.GetReceiver()),
		Options:            ToNullString(event.GetOptions()),
		Fee:                event.GetFee(),
		TimeStamp:          nil,
		RevertedReason:     ToNullString(event.GetRevertReason()),
		FeeUSD:             nil,
	}
}

// ParserType returns the type of parser.
func (m *MessageBusParser) ParserType() string {
	return "messagebus"
}

// Parse parses the message logs.
//
// nolint:gocognit,cyclop,dupl
func (m *MessageBusParser) Parse(ctx context.Context, log ethTypes.Log, chainID uint32) (interface{}, error) {
	logTopic := log.Topics[0]
	iFace, err := func(log ethTypes.Log) (messageBusTypes.EventLog, error) {
		switch logTopic {
		case messagebus.Topic(messageBusTypes.ExecutedEvent):
			iFace, err := m.Filterer.ParseExecuted(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse token : %w", err)
			}
			return iFace, nil
		case messagebus.Topic(messageBusTypes.MessageSentEvent):
			iFace, err := m.Filterer.ParseMessageSent(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse sent message: %w", err)
			}
			return iFace, nil
		case messagebus.Topic(messageBusTypes.CallRevertedEvent):
			iFace, err := m.Filterer.ParseCallReverted(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse sent message: %w", err)
			}
			return iFace, nil

		default:
			logger.Warnf("ErrUnknownTopic in messagebus: %s %s chain: %d address: %s", log.TxHash, logTopic.String(), chainID, log.Address.Hex())

			return nil, fmt.Errorf(ErrUnknownTopic)
		}
	}(log)

	if err != nil {
		// Switch failed.

		return nil, err
	}
	if iFace == nil {
		// Unknown topic.
		return nil, fmt.Errorf("unknwn topic")
	}

	// populate message event type so following operations can mature the event data.
	messageEvent := eventToMessageEvent(iFace, chainID)

	// Get timestamp from consumer
	timeStamp, err := m.consumerFetcher.FetchBlockTime(ctx, int(chainID), int(iFace.GetBlockNumber()))
	if err != nil {
		return nil, fmt.Errorf("could not get block time: %w", err)
	}

	// If we have a timestamp, populate the following attributes of messageEvent.
	timeStampBig := uint64(*timeStamp)
	messageEvent.TimeStamp = &timeStampBig

	if messageEvent.Fee != nil {
		switch messageEvent.ChainID {
		case 8217: // Klaytn
			coinGeckoID := "klay-token"
			feeValue, err := m.getFeeValue(ctx, messageEvent, coinGeckoID)
			if err != nil {
				return nil, err
			}
			messageEvent.FeeUSD = feeValue
		case 53935: // DFK
			coinGeckoID := "defi-kingdoms"
			feeValue, err := m.getFeeValue(ctx, messageEvent, coinGeckoID)
			if err != nil {
				return nil, err
			}
			messageEvent.FeeUSD = feeValue
		case 1666600000: // Harmony
			coinGeckoID := "harmony"
			feeValue, err := m.getFeeValue(ctx, messageEvent, coinGeckoID)
			if err != nil {
				return nil, err
			}
			messageEvent.FeeUSD = feeValue
		default:
			// pass
		}
	}
	return messageEvent, nil
}

func (m *MessageBusParser) getFeeValue(ctx context.Context, messageEvent model.MessageBusEvent, coinGeckoID string) (*float64, error) {
	tokenPrice := m.tokenPriceService.GetPriceData(ctx, int(*messageEvent.TimeStamp), coinGeckoID)
	if tokenPrice == nil {
		return nil, fmt.Errorf("MESSAGEBUS could not get token price for coingeckotoken:  %s chain: %d txhash %s %d", coinGeckoID, messageEvent.ChainID, messageEvent.TxHash, messageEvent.TimeStamp)
	}
	price := GetAmountUSD(messageEvent.Fee, 18, tokenPrice)
	if price != nil {
		return price, nil
	}
	return nil, fmt.Errorf("MESSAGEBUS could not convert token price:  %s chain: %d txhash %s %d", coinGeckoID, messageEvent.ChainID, messageEvent.TxHash, messageEvent.TimeStamp)
}
