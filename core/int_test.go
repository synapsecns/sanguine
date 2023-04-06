package core_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"math/big"
)

func (c *CoreSuite) TestCopyBigInt() {
	const ogInt = 5
	ogBigInt := big.NewInt(ogInt)
	newBigInt := core.CopyBigInt(ogBigInt)

	Equal(c.T(), ogBigInt.Int64(), newBigInt.Int64())
	// make sure pointers are different
	False(c.T(), core.ArePointersEqual(ogBigInt, newBigInt))
	Equal(c.T(), ogBigInt.Int64(), int64(ogInt))
}

func (c *CoreSuite) TestCopyNullInt() {
	NotPanics(c.T(), func() {
		copiedInt := core.CopyBigInt(nil)
		Nil(c.T(), copiedInt)
	})
}
