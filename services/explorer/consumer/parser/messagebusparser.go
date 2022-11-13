package parser

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/contracts/messagebus"
	"github.com/synapsecns/sanguine/services/explorer/db"
	model "github.com/synapsecns/sanguine/services/explorer/db/sql"
	messageBusTypes "github.com/synapsecns/sanguine/services/explorer/types/messagebus"
	"time"
)

type MessageParser struct {
	// consumerDB is the database to store parsed data in
	consumerDB db.ConsumerDB
	// Filterer is the message Filterer we use to parse events
	Filterer *messagebus.MessageBusUpgradeableFilterer
	// messageAddress is the address of the message
	messageAddress common.Address
	// consumerFetcher is the Fetcher for sender and timestamp
	consumerFetcher *fetcher.ScribeFetcher
}

// NewMessageParser creates a new parser for a given message.
func NewMessageParser(consumerDB db.ConsumerDB, messageAddress common.Address, consumerFetcher *fetcher.ScribeFetcher) (*MessageParser, error) {
	filterer, err := messagebus.NewMessageBusUpgradeableFilterer(messageAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", messagebus.MessageBusReceiverUpgradeableFilterer{}, err)
	}
	return &MessageParser{consumerDB, filterer, messageAddress, consumerFetcher}, nil
}

// EventType returns the event type of a message log.
func (m *MessageParser) EventType(log ethTypes.Log) (_ messageBusTypes.EventType, ok bool) {
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
func eventToMessageEvent(event messageBusTypes.EventLog, chainID uint32) model.MessageEvent {
	return model.MessageEvent{
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
	}
}

// ParseAndStore parses the message logs and stores them in the database.
//
// nolint:gocognit,cyclop,dupl
func (m *MessageParser) ParseAndStore(ctx context.Context, log ethTypes.Log, chainID uint32) error {
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
			return nil, fmt.Errorf("unknown topic: %s", logTopic.Hex())
		}
	}(log)

	if err != nil {
		// Switch failed.
		return err
	}

	// populate message event type so following operations can mature the event data.
	messageEvent := eventToMessageEvent(iFace, chainID)

	// Get timestamp from consumer
	timeStamp, err := m.consumerFetcher.FetchClient.GetBlockTime(ctx, int(chainID), int(iFace.GetBlockNumber()))

	// If we have a timestamp, populate the following attributes of messageEvent.
	if err == nil {
		timeStampBig := uint64(*timeStamp.Response)
		messageEvent.TimeStamp = &timeStampBig
	}

	err = m.consumerDB.StoreEvent(ctx, &messageEvent)
	if err != nil {
		return fmt.Errorf("could not store event: %w", err)
	}
	
	return nil
}
