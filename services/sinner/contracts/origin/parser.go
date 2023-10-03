package origin

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/services/sinner/db"
	"github.com/synapsecns/sanguine/services/sinner/db/model"
	sinnerTypes "github.com/synapsecns/sanguine/services/sinner/types"
)

type ParserImpl struct {
	filterer *origin.OriginFilterer
	// parser is the parser interface.
	parser origin.Parser
	// db is the database
	db db.EventDB
	// txMap is a map of tx hashes to tx data
	txMap map[string]sinnerTypes.TxSupplementalInfo
}

// NewParser creates a new parser for the origin contract.
func NewParser(originAddress common.Address, db db.EventDB) (*ParserImpl, error) {
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

	parser := &ParserImpl{
		filterer: filter,
		parser:   agentsParser,
		db:       db,
	}
	return parser, nil
}

func (p ParserImpl) UpdateTxMap(txMap map[string]sinnerTypes.TxSupplementalInfo) {
	p.txMap = txMap
}

func (p ParserImpl) ParseAndStore(ctx context.Context, log ethTypes.Log) error {
	eventType, ok := p.parser.EventType(log)
	if !ok {
		return fmt.Errorf("could not parse log event type. Topics: %v", log.Topics)
	}
	switch eventType {
	case origin.SentEvent:
		parsedEvent, err := p.ParseSent(log)
		if err != nil {
			return fmt.Errorf("error while parsing origin sent event. Err: %w", err)
		}

		// TODO go func this
		err = p.db.StoreOrUpdateMessageStatus(ctx, parsedEvent.TxHash, parsedEvent.MessageHash, sinnerTypes.Origin)
		if err != nil {
			return fmt.Errorf("error while storing origin sent event. Err: %w", err)
		}

		err = p.db.StoreOriginSent(ctx, parsedEvent)
		if err != nil {
			return fmt.Errorf("error while storing origin sent event. Err: %w", err)
		}

	}
	return nil
}

func (p ParserImpl) ParseSent(log ethTypes.Log) (*model.OriginSent, error) {

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
		Message:            iFace.Message,
		Nonce:              iFace.Nonce,
		MessageHash:        string(iFace.MessageHash[:]),
	}
	parsedMessage, err := types.DecodeMessage(iFace.Message)
	if err != nil {
		return nil, fmt.Errorf("could not decode message. err: %w", err)
	}

	messageLeaf, err := parsedMessage.ToLeaf()
	if err != nil {
		return nil, fmt.Errorf("could not get leaf from message. err: %w", err)
	}
	parsedEvent.MessageLeaf = string(messageLeaf[:])
	parsedEvent.OptimisticSeconds = parsedMessage.OptimisticSeconds()

	messageHeader := parsedMessage.Header()
	parsedEvent.MessageFlag = uint8(messageHeader.Flag())

	messageBody := parsedMessage.BaseMessage()

	sender := messageBody.Sender()
	parsedEvent.Sender = string(sender[:])

	recipient := messageBody.Recipient()
	parsedEvent.Recipient = string(recipient[:])

	messageRequest := messageBody.Request()
	parsedEvent.Version = messageRequest.Version()
	parsedEvent.GasLimit = messageRequest.GasLimit()
	parsedEvent.GasDrop = messageRequest.GasDrop().String()

	messageTips := messageBody.Tips()
	parsedEvent.SummitTip = messageTips.SummitTip().String()
	parsedEvent.AttestationTip = messageTips.AttestationTip().String()
	parsedEvent.ExecutionTip = messageTips.ExecutionTip().String()
	parsedEvent.DeliveryTip = messageTips.DeliveryTip().String()

	return &parsedEvent, nil
}
