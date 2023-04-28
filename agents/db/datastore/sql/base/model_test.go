package base_test

import (
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
		CMNonce:             gofakeit.Uint32(),
		CMDestination:       gofakeit.Uint32(),
		CMBody:              []byte(gofakeit.Paragraph(4, 1, 4, " ")),
		CMOptimisticSeconds: gofakeit.Uint32(),
	}

	Equal(t, cm.OriginDomain(), cm.CMOrigin)
	Equal(t, cm.Nonce(), cm.CMNonce)
	Equal(t, cm.DestinationDomain(), cm.CMDestination)
	Equal(t, cm.Body(), cm.CMBody)

	toLeaf, err := cm.ToLeaf()
	Nil(t, err)
	Equal(t, toLeaf[:], cm.CMLeaf)

	Equal(t, cm.OptimisticSeconds(), cm.CMOptimisticSeconds)
	Equal(t, cm.Message(), cm.CMMessage)
	Equal(t, core.BytesToSlice(cm.Leaf()), cm.CMLeaf)
}
