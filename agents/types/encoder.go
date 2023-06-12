package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/libs4go/crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/math"
)

const (
	uint16Len = 2
	uint32Len = 4
	uint40Len = 5
)

// EncodeGasData encodes a gasdata.
func EncodeGasData(gasData GasData) ([]byte, error) {
	b := make([]byte, 0)
	markupBytes := make([]byte, uint16Len)
	etherPriceBytes := make([]byte, uint16Len)
	amortAttCostBytes := make([]byte, uint16Len)
	execBufferBytes := make([]byte, uint16Len)
	dataPriceBytes := make([]byte, uint16Len)
	gasPriceBytes := make([]byte, uint16Len)

	binary.BigEndian.PutUint16(markupBytes, gasData.Markup())
	binary.BigEndian.PutUint16(etherPriceBytes, gasData.EtherPrice())
	binary.BigEndian.PutUint16(amortAttCostBytes, gasData.AmortAttCost())
	binary.BigEndian.PutUint16(execBufferBytes, gasData.ExecBuffer())
	binary.BigEndian.PutUint16(dataPriceBytes, gasData.DataPrice())
	binary.BigEndian.PutUint16(gasPriceBytes, gasData.GasPrice())

	b = append(b, gasPriceBytes...)
	b = append(b, dataPriceBytes...)
	b = append(b, execBufferBytes...)
	b = append(b, amortAttCostBytes...)
	b = append(b, etherPriceBytes...)
	b = append(b, markupBytes...)

	return b, nil
}

// DecodeGasData decodes a gasData.
func DecodeGasData(toDecode []byte) (GasData, error) {
	if len(toDecode) != gasDataSize {
		return nil, fmt.Errorf("invalid gasData length, expected %d, got %d", gasDataSize, len(toDecode))
	}

	gasPrice := binary.BigEndian.Uint16(toDecode[gasDataOffsetGasPrice:gasDataOffsetDataPrice])
	dataPrice := binary.BigEndian.Uint16(toDecode[gasDataOffsetDataPrice:gasDataOffsetExecBuffer])
	execBuffer := binary.BigEndian.Uint16(toDecode[gasDataOffsetExecBuffer:gasDataOffsetAmortAttCost])
	amortAttCost := binary.BigEndian.Uint16(toDecode[gasDataOffsetAmortAttCost:gasDataOffsetEtherPrice])
	etherPrice := binary.BigEndian.Uint16(toDecode[gasDataOffsetEtherPrice:gasDataOffsetMarkup])
	markup := binary.BigEndian.Uint16(toDecode[gasDataOffsetMarkup:gasDataSize])

	return gasData{
		markup:       markup,
		etherPrice:   etherPrice,
		amortAttCost: amortAttCost,
		execBuffer:   execBuffer,
		dataPrice:    dataPrice,
		gasPrice:     gasPrice,
	}, nil
}

// EncodeChainGas encodes a chaingas.
func EncodeChainGas(chainGas ChainGas) ([]byte, error) {
	b := make([]byte, 0)
	domainBytes := make([]byte, uint32Len)
	domain := chainGas.Domain()
	binary.BigEndian.PutUint32(domainBytes, domain)
	b = append(b, domainBytes...)

	gasDataEncoded, err := EncodeGasData(chainGas.GasData())
	if err != nil {
		return nil, fmt.Errorf("failed to encode gas data for chain gas %w", err)
	}

	b = append(b, gasDataEncoded...)

	return b, nil
}

// DecodeChainGas decodes a chainGas.
func DecodeChainGas(toDecode []byte) (ChainGas, error) {
	if len(toDecode) != chainGasSize {
		return nil, fmt.Errorf("invalid chainGas length, expected %d, got %d", chainGasSize, len(toDecode))
	}

	domain := binary.BigEndian.Uint32(toDecode[chainGasOffsetDomain:chainGasOffsetGasData])
	gasData, err := DecodeGasData(toDecode[chainGasOffsetGasData:chainGasSize])
	if err != nil {
		return nil, fmt.Errorf("failed to decode gas data for chain gas %w", err)
	}

	return chainGas{
		gasData: gasData,
		domain:  domain,
	}, nil
}

// EncodeState encodes a state.
func EncodeState(state State) ([]byte, error) {
	b := make([]byte, 0)
	originBytes := make([]byte, uint32Len)
	nonceBytes := make([]byte, uint32Len)

	binary.BigEndian.PutUint32(originBytes, state.Origin())
	binary.BigEndian.PutUint32(nonceBytes, state.Nonce())
	root := state.Root()

	b = append(b, root[:]...)
	b = append(b, originBytes...)
	b = append(b, nonceBytes...)
	b = append(b, math.PaddedBigBytes(state.BlockNumber(), uint40Len)...)
	b = append(b, math.PaddedBigBytes(state.Timestamp(), uint40Len)...)

	gasDataEncoded, err := EncodeGasData(state.GasData())
	if err != nil {
		return nil, fmt.Errorf("failed to encode gas data for state %w", err)
	}
	b = append(b, gasDataEncoded...)

	return b, nil
}

// DecodeState decodes a state.
func DecodeState(toDecode []byte) (State, error) {
	if len(toDecode) != stateSize {
		return nil, fmt.Errorf("invalid state length, expected %d, got %d", stateSize, len(toDecode))
	}

	root := toDecode[stateOffsetRoot:stateOffsetOrigin]
	origin := binary.BigEndian.Uint32(toDecode[stateOffsetOrigin:stateOffsetNonce])
	nonce := binary.BigEndian.Uint32(toDecode[stateOffsetNonce:stateOffsetBlockNumber])
	blockNumber := new(big.Int).SetBytes(toDecode[stateOffsetBlockNumber:stateOffsetTimestamp])
	timestamp := new(big.Int).SetBytes(toDecode[stateOffsetTimestamp:stateOffsetGasData])

	gasDataToDecode := toDecode[stateOffsetGasData:stateSize]
	gasData, err := DecodeGasData(gasDataToDecode)
	if err != nil {
		return nil, fmt.Errorf("failed to decode gas data for state %w", err)
	}

	var rootB32 [32]byte
	copy(rootB32[:], root)

	return state{
		root:        rootB32,
		origin:      origin,
		nonce:       nonce,
		blockNumber: blockNumber,
		timestamp:   timestamp,
		gasData:     gasData,
	}, nil
}

// EncodeSnapshot encodes a snapshot.
func EncodeSnapshot(snapshot Snapshot) ([]byte, error) {
	states := snapshot.States()

	if len(states) == 0 {
		return nil, fmt.Errorf("no states to encode")
	}

	encodedStates := make([]byte, 0)

	for _, state := range states {
		encodedState, err := EncodeState(state)
		if err != nil {
			return nil, fmt.Errorf("could not encode state: %w", err)
		}
		encodedStates = append(encodedStates, encodedState...)
	}

	return encodedStates, nil
}

// DecodeSnapshot decodes a snapshot.
func DecodeSnapshot(toDecode []byte) (Snapshot, error) {
	var states []State

	if len(toDecode)%stateSize != 0 {
		return nil, fmt.Errorf("invalid snapshot length, expected multiple of %d, got %d", stateSize, len(toDecode))
	}

	for i := 0; i < len(toDecode); i += stateSize {
		state, err := DecodeState(toDecode[i : i+stateSize])
		if err != nil {
			return nil, fmt.Errorf("could not decode state: %w", err)
		}
		states = append(states, state)
	}

	return snapshot{
		states: states,
	}, nil
}

// EncodeAttestation encodes an attestation.
func EncodeAttestation(attestation Attestation) ([]byte, error) {
	b := make([]byte, 0)
	nonceBytes := make([]byte, uint32Len)

	binary.BigEndian.PutUint32(nonceBytes, attestation.Nonce())
	snapshotRoot := attestation.SnapshotRoot()
	dataHash := attestation.DataHash()

	b = append(b, snapshotRoot[:]...)
	b = append(b, dataHash[:]...)
	b = append(b, nonceBytes...)
	b = append(b, math.PaddedBigBytes(attestation.BlockNumber(), uint40Len)...)
	b = append(b, math.PaddedBigBytes(attestation.Timestamp(), uint40Len)...)

	return b, nil
}

// DecodeAttestation decodes an attestation.
func DecodeAttestation(toDecode []byte) (Attestation, error) {
	if len(toDecode) != attestationSize {
		return nil, fmt.Errorf("invalid attestation length, expected %d, got %d", attestationSize, len(toDecode))
	}

	snapshotRoot := toDecode[attestationOffsetRoot:attestationOffsetDataHash]
	dataHash := toDecode[attestationOffsetDataHash:attestationOffsetNonce]
	nonce := binary.BigEndian.Uint32(toDecode[attestationOffsetNonce:attestationOffsetBlockNumber])
	blockNumber := new(big.Int).SetBytes(toDecode[attestationOffsetBlockNumber:attestationOffsetTimestamp])
	timestamp := new(big.Int).SetBytes(toDecode[attestationOffsetTimestamp:attestationSize])

	var snapshotRootB32, dataHashB32 [32]byte
	copy(snapshotRootB32[:], snapshotRoot)
	copy(dataHashB32[:], dataHash)

	return attestation{
		snapshotRoot: snapshotRootB32,
		dataHash:     dataHashB32,
		nonce:        nonce,
		blockNumber:  blockNumber,
		timestamp:    timestamp,
	}, nil
}

// EncodeSignature encodes a signature.
func EncodeSignature(sig Signature) ([]byte, error) {
	return ecdsa.Sig2Bytes(secp256k1.S256(), sig.R(), sig.S(), sig.V()), nil
}

// DecodeSignature decodes a signature.
func DecodeSignature(toDecode []byte) (sig Signature, err error) {
	r, s, v, err := ecdsa.Bytes2Sig(secp256k1.S256(), toDecode)
	if err != nil {
		return nil, fmt.Errorf("could not decode signature: %w", err)
	}

	return NewSignature(v, r, s), nil
}

// HashRawBytes takes the raw bytes and produces a hash.
func HashRawBytes(rawBytes []byte) (common.Hash, error) {
	hashedDigest := crypto.Keccak256Hash(rawBytes)

	signedHash := crypto.Keccak256Hash([]byte("\x19Ethereum Signed Message:\n32"), hashedDigest.Bytes())
	return signedHash, nil
}

const (
	uint64Len = 8
	uint96Len = 12
)

// EncodeTips encodes a list of tips.
//
//nolint:makezero
func EncodeTips(tips Tips) ([]byte, error) {
	b := make([]byte, 0)

	b = append(b, math.PaddedBigBytes(tips.SummitTip(), uint64Len)...)
	b = append(b, math.PaddedBigBytes(tips.AttestationTip(), uint64Len)...)
	b = append(b, math.PaddedBigBytes(tips.ExecutionTip(), uint64Len)...)
	b = append(b, math.PaddedBigBytes(tips.DeliveryTip(), uint64Len)...)

	return b, nil
}

// DecodeTips decodes a tips typed mem view.
func DecodeTips(toDecode []byte) (Tips, error) {
	summitTip := new(big.Int).SetBytes(toDecode[0:8])
	attestationTip := new(big.Int).SetBytes(toDecode[8:16])
	executionTip := new(big.Int).SetBytes(toDecode[16:24])
	deliveryTip := new(big.Int).SetBytes(toDecode[24:])

	return NewTips(summitTip, attestationTip, executionTip, deliveryTip), nil
}

type headerEncoder struct {
	Flag              MessageFlag
	OriginDomain      uint32
	Nonce             uint32
	DestinationDomain uint32
	OptimisticSeconds uint32
}

// EncodeHeader encodes a message header.
func EncodeHeader(header Header) ([]byte, error) {
	newHeader := headerEncoder{
		Flag:              header.Flag(),
		OriginDomain:      header.OriginDomain(),
		Nonce:             header.Nonce(),
		DestinationDomain: header.DestinationDomain(),
		OptimisticSeconds: header.OptimisticSeconds(),
	}

	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, newHeader)
	if err != nil {
		return nil, fmt.Errorf("could not write binary: %w", err)
	}

	return buf.Bytes(), nil
}

// DecodeHeader decodes a header from a byte slice.
func DecodeHeader(header []byte) (Header, error) {
	reader := bytes.NewReader(header)

	var encoded headerEncoder

	err := binary.Read(reader, binary.BigEndian, &encoded)
	if err != nil {
		return nil, fmt.Errorf("failed to decode header: %w", err)
	}

	decoded := headerImpl{
		flag:              encoded.Flag,
		originDomain:      encoded.OriginDomain,
		nonce:             encoded.Nonce,
		destinationDomain: encoded.DestinationDomain,
		optimisticSeconds: encoded.OptimisticSeconds,
	}

	return decoded, nil
}

// EncodeRequest encodes a request.
func EncodeRequest(m Request) ([]byte, error) {
	b := make([]byte, 0)

	versionBytes := make([]byte, uint32Len)
	binary.BigEndian.PutUint32(versionBytes, m.Version())

	gasLimitBytes := make([]byte, uint64Len)
	binary.BigEndian.PutUint64(gasLimitBytes, m.GasLimit())

	b = append(b, versionBytes...)
	b = append(b, gasLimitBytes...)
	b = append(b, math.PaddedBigBytes(m.GasDrop(), uint96Len)...)

	return b, nil
}

// DecodeRequest decodes a request typed mem view.
func DecodeRequest(toDecode []byte) Request {
	version := binary.BigEndian.Uint32(toDecode[VersionOffset:GasLimitOffset])
	gasLimit := binary.BigEndian.Uint64(toDecode[GasLimitOffset:GasDropOffset])
	gasDrop := new(big.Int).SetBytes(toDecode[GasDropOffset:RequestSize])

	return NewRequest(version, gasLimit, gasDrop)
}

// EncodeBaseMessage encodes a base message.
func EncodeBaseMessage(m BaseMessage) ([]byte, error) {
	b := make([]byte, 0)

	encodedTips, err := EncodeTips(m.Tips())
	if err != nil {
		return []byte{}, fmt.Errorf("could not encode tips part of message: %w", err)
	}
	b = append(b, encodedTips...)

	senderRef := m.Sender()
	recipientRef := m.Recipient()

	b = append(b, senderRef[:]...)
	b = append(b, recipientRef[:]...)

	encodedRequest, err := EncodeRequest(m.Request())
	if err != nil {
		return []byte{}, fmt.Errorf("could not encode request part of message: %w", err)
	}
	b = append(b, encodedRequest...)
	b = append(b, m.Content()...)

	return b, nil
}

// DecodeBaseMessage decodes a base message typed mem view.
func DecodeBaseMessage(toDecode []byte) (BaseMessage, error) {
	if len(toDecode) < BaseMessageContentOffset {
		return nil, fmt.Errorf("invalid attestation length, expected at least %d, got %d", BaseMessageContentOffset, len(toDecode))
	}

	tipsBytes := toDecode[BaseMessageTipsOffset:BaseMessageSenderOffset]
	senderBytes := toDecode[BaseMessageSenderOffset:BaseMessageRecipientOffset]
	recipientBytes := toDecode[BaseMessageRecipientOffset:BaseMessageRequestOffset]
	requestBytes := toDecode[BaseMessageRequestOffset:BaseMessageContentOffset]
	content := toDecode[BaseMessageContentOffset:]

	var sender [32]byte
	var recipient [32]byte
	copy(sender[:], senderBytes)
	copy(recipient[:], recipientBytes)

	decodedTips, err := DecodeTips(tipsBytes)
	if err != nil {
		return nil, fmt.Errorf("could not decode tips part of message: %w", err)
	}

	request := DecodeRequest(requestBytes)

	return NewBaseMessage(sender, recipient, decodedTips, request, content), nil
}

// EncodeMessage encodes a message.
func EncodeMessage(m Message) ([]byte, error) {
	encodedHeader, err := EncodeHeader(m.Header())
	if err != nil {
		return []byte{}, fmt.Errorf("could not encode header: %w", err)
	}

	buf := new(bytes.Buffer)

	buf.Write(encodedHeader)

	if m.Header().Flag() == MessageFlagBase {
		encodedBaseMessage, err := EncodeBaseMessage(m.BaseMessage())
		if err != nil {
			return []byte{}, fmt.Errorf("could not encode header: %w", err)
		}
		buf.Write(encodedBaseMessage)
	} else {
		buf.Write(m.Body())
	}

	return buf.Bytes(), nil
}

// DecodeMessage decodes a message from a byte slice.
func DecodeMessage(message []byte) (Message, error) {
	rawHeader := message[:MessageBodyOffset]

	header, err := DecodeHeader(rawHeader)
	if err != nil {
		return nil, fmt.Errorf("could not decode header: %w", err)
	}

	rawBody := message[MessageBodyOffset:]

	var decoded Message

	var content []byte
	if header.Flag() == MessageFlagBase {
		baseMessage, err := DecodeBaseMessage(rawBody)
		if err != nil {
			return nil, fmt.Errorf("could not decode base message: %w", err)
		}

		decoded = messageImpl{
			header:      header,
			baseMessage: baseMessage,
			body:        rawBody,
		}
	} else {
		content = rawBody
		decoded = messageImpl{
			header: header,
			body:   content,
		}
	}

	return decoded, nil
}

// EncodeAgentStatus encodes a agent status.
func EncodeAgentStatus(agentStatus AgentStatus) ([]byte, error) {
	b := make([]byte, 0)
	domainBytes := make([]byte, uint32Len)
	indexBytes := make([]byte, uint32Len)

	binary.BigEndian.PutUint32(domainBytes, agentStatus.Domain())
	binary.BigEndian.PutUint32(indexBytes, agentStatus.Index())

	b = append(b, agentStatus.Flag())
	b = append(b, domainBytes...)
	b = append(b, indexBytes...)

	return b, nil
}

// DecodeAgentStatus decodes an agent status.
func DecodeAgentStatus(toDecode []byte) (AgentStatus, error) {
	if len(toDecode) != agentStatusSize {
		return nil, fmt.Errorf("invalid agent status length, expected %d, got %d", agentStatusSize, len(toDecode))
	}

	flagBytes := toDecode[agentStatusOffsetFlag:agentStatusOffsetDomain]
	domain := binary.BigEndian.Uint32(toDecode[agentStatusOffsetDomain:agentStatusOffsetIndex])
	index := binary.BigEndian.Uint32(toDecode[agentStatusOffsetDomain:agentStatusSize])

	return agentStatus{
		flag:   flagBytes[0],
		domain: domain,
		index:  index,
	}, nil
}
