package core_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	common "github.com/synapsecns/sanguine/core"
	"os"
)

// TestGetEnv makes sure that default variables are set/fetched.
func (c *CoreSuite) TestGetEnv() {
	testWord := gofakeit.Word()
	testValue := gofakeit.Word()

	Equal(c.T(), common.GetEnv(testWord, testValue), testValue)

	c.T().Setenv(testWord, gofakeit.Word())
	NotEqual(c.T(), testValue, common.GetEnv(testWord, testValue))

	c.T().Setenv(testWord, testValue)
	Equal(c.T(), testValue, common.GetEnv(testWord, testValue))
}

func (c *CoreSuite) TestHasEnv() {
	fakeEnvWord := gofakeit.Word()
	realEnvWord := gofakeit.Word()

	c.T().Setenv(realEnvWord, gofakeit.Word())
	c.Require().True(common.HasEnv(realEnvWord))
	c.Require().False(common.HasEnv(fakeEnvWord))
}

func (c *CoreSuite) TestGetEnvInt() {
	Equal(c.T(), common.GetEnvInt(gofakeit.Word(), 1), 1)
	Nil(c.T(), os.Setenv("invalid", "invalid"))
	Equal(c.T(), common.GetEnvInt("invalid", 1), 1)
	Nil(c.T(), os.Setenv("valid", "3"))
	Equal(c.T(), common.GetEnvInt("valid", 1), 3)
}

func (c *CoreSuite) TestIsTest() {
	True(c.T(), common.IsTest())
}
