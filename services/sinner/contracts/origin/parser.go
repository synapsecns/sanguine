// Package origin is the origin contract parser.
package origin

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/sinner/logger"

	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/services/sinner/db"
	"github.com/synapsecns/sanguine/services/sinner/db/model"
	sinnerTypes "github.com/synapsecns/sanguine/services/sinner/types"
)

// parserImpl is the parser for the origin contract.
type parserImpl struct {
	filterer *origin.OriginFilterer
	// parser is the parser interface.
	parser origin.Parser
	// db is the database
	db db.EventDB
	// chainID is the chainID of the underlying chain
	chainID uint32
}

// NewParser creates a new parser for the origin contract.
func NewParser(originAddress common.Address, db db.EventDB, chainID uint32) (sinnerTypes.EventParser, error) {
	// Get agents parser to utilize event type parsing.
	agentsParser, err := origin.NewParser(originAddress)
	if err != nil {
		return nil, fmt.Errorf("could not create agents parser %w", err)
	}

	// Get filterer to get ABI IFace from abi.
	filter, err := origin.NewOriginFilterer(originAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", origin.OriginFilterer{}, err)
	}

	parser := &parserImpl{
		filterer: filter,
		parser:   agentsParser,
		db:       db,
		chainID:  chainID,
	}
	return parser, nil
}

// ParseAndStore parses and stores the log.
func (p *parserImpl) ParseAndStore(ctx context.Context, log ethTypes.Log, tx sinnerTypes.TxSupplementalInfo) error {
	eventType, ok := p.parser.EventType(log)

	if !ok {
		logger.ReportSinnerError(fmt.Errorf("unknown origin log topic"), 0, logger.UnknownTopic)
		return nil
	}
	if eventType == origin.SentEvent {
		parsedEvent, err := p.parseSent(log, tx)
		if err != nil {
			return fmt.Errorf("error while parsing origin sent event. Err: %w", err)
		}

		g, storeCtx := errgroup.WithContext(ctx)

		g.Go(func() error {
			err := p.db.StoreOrUpdateMessageStatus(storeCtx, parsedEvent.TxHash, parsedEvent.MessageHash, sinnerTypes.Origin)
			if err != nil {
				return fmt.Errorf("error while storing origin sent event. Err: %w", err)
			}
			return nil
		})

		g.Go(func() error {
			err := p.db.StoreOriginSent(storeCtx, parsedEvent)
			if err != nil {
				return fmt.Errorf("error while storing origin sent event. Err: %w", err)
			}
			return nil
		})

		err = g.Wait()
		if err != nil {
			return fmt.Errorf("error while storing origin sent event. Err: %w", err)
		}
	}
	return nil
}

// parseSent parses the sent event.
func (p *parserImpl) parseSent(log ethTypes.Log, tx sinnerTypes.TxSupplementalInfo) (*model.OriginSent, error) {
	iFace, err := p.filterer.ParseSent(log)
	if err != nil {
		return nil, fmt.Errorf("could not parse sent log. err: %w", err)
	}
	parsedEvent := model.OriginSent{
		ContractAddress:    iFace.Raw.Address.String(),
		BlockNumber:        iFace.Raw.BlockNumber,
		TxHash:             iFace.Raw.TxHash.String(),
		TxIndex:            iFace.Raw.TxIndex,
		DestinationChainID: iFace.Destination,
		Message:            common.Bytes2Hex(iFace.Message),
		Nonce:              iFace.Nonce,
		MessageHash:        common.Bytes2Hex(iFace.MessageHash[:]),
		Timestamp:          uint64(tx.Timestamp),
		Sender:             tx.Sender,
	}

	parsedEvent.ChainID = p.chainID

	// This case will be hit unless there was a failure in producing the message
	if len(iFace.Message) > 0 {
		parsedMessage, err := types.DecodeMessage(iFace.Message)
		if err != nil {
			return nil, fmt.Errorf("could not decode message. err: %w", err)
		}

		messageLeaf, err := parsedMessage.ToLeaf()
		if err != nil {
			return nil, fmt.Errorf("could not get leaf from message. err: %w", err)
		}
		parsedEvent.MessageLeaf = common.Bytes2Hex(messageLeaf[:])
		parsedEvent.OptimisticSeconds = parsedMessage.OptimisticSeconds()

		messageHeader := parsedMessage.Header()
		parsedEvent.MessageFlag = uint8(messageHeader.Flag())

		messageBody := parsedMessage.BaseMessage()

		sender := messageBody.Sender()
		parsedEvent.Sender = common.Bytes2Hex(sender[:])

		recipient := messageBody.Recipient()
		parsedEvent.Recipient = common.Bytes2Hex(recipient[:])

		messageRequest := messageBody.Request()
		parsedEvent.Version = messageRequest.Version()
		parsedEvent.GasLimit = messageRequest.GasLimit()
		parsedEvent.GasDrop = messageRequest.GasDrop().String()

		messageTips := messageBody.Tips()
		parsedEvent.SummitTip = messageTips.SummitTip().String()
		parsedEvent.AttestationTip = messageTips.AttestationTip().String()
		parsedEvent.ExecutionTip = messageTips.ExecutionTip().String()
		parsedEvent.DeliveryTip = messageTips.DeliveryTip().String()
	}
	return &parsedEvent, nil
}

var _ sinnerTypes.EventParser = &parserImpl{}
