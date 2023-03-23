package base_test

import (
	"math/big"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql/base"
	"github.com/synapsecns/sanguine/core"
)

func TestCommittedMessageAccessors(t *testing.T) {
	cm := base.CommittedMessage{
		CMDomainID:          gofakeit.Uint32(),
		CMMessage:           []byte(gofakeit.Paragraph(4, 1, 4, " ")),
		CMLeaf:              common.BytesToHash([]byte(gofakeit.Paragraph(4, 1, 4, " "))).Bytes(),
		CMOrigin:            gofakeit.Uint32(),
		CMSender:            common.BytesToHash([]byte(gofakeit.Paragraph(4, 1, 4, " "))).Bytes(),
		CMNonce:             gofakeit.Uint32(),
		CMDestination:       gofakeit.Uint32(),
		CMRecipient:         common.BytesToHash([]byte(gofakeit.Paragraph(4, 1, 4, " "))).Bytes(),
		CMBody:              []byte(gofakeit.Paragraph(4, 1, 4, " ")),
		CMOptimisticSeconds: gofakeit.Uint32(),
		CMNotaryTip:         new(big.Int).SetUint64(gofakeit.Uint64()).Bytes(),
		CMBroadcasterTip:    new(big.Int).SetUint64(gofakeit.Uint64()).Bytes(),
		CMExecutorTip:       new(big.Int).SetUint64(gofakeit.Uint64()).Bytes(),
		CMProverTip:         new(big.Int).SetUint64(gofakeit.Uint64()).Bytes(),
	}

	Equal(t, cm.OriginDomain(), cm.CMOrigin)
	Equal(t, cm.Sender().Bytes(), cm.CMSender)
	Equal(t, cm.Nonce(), cm.CMNonce)
	Equal(t, cm.DestinationDomain(), cm.CMDestination)
	Equal(t, cm.Recipient().Bytes(), cm.CMRecipient)
	Equal(t, cm.Body(), cm.CMBody)

	toLeaf, err := cm.ToLeaf()
	Nil(t, err)
	Equal(t, toLeaf[:], cm.CMLeaf)

	Equal(t, cm.OptimisticSeconds(), cm.CMOptimisticSeconds)
	Equal(t, cm.Message(), cm.CMMessage)
	Equal(t, core.BytesToSlice(cm.Leaf()), cm.CMLeaf)

	Equal(t, cm.CMProverTip, cm.Tips().ProverTip().Bytes())
	Equal(t, cm.CMExecutorTip, cm.Tips().ExecutorTip().Bytes())
	Equal(t, cm.CMBroadcasterTip, cm.Tips().BroadcasterTip().Bytes())
	Equal(t, cm.CMNotaryTip, cm.Tips().NotaryTip().Bytes())
}
