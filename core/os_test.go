package core_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	common "github.com/synapsecns/sanguine/core"
	"os"
)

// TestGetEnv makes sure that default variables are set/fetched.
func (s *CoreSuite) TestGetEnv() {
	testWord := gofakeit.Word()
	testValue := gofakeit.Word()

	Equal(s.T(), common.GetEnv(testWord, testValue), testValue)

	err := os.Setenv(testWord, gofakeit.Word())
	Nil(s.T(), err)
	NotEqual(s.T(), testValue, common.GetEnv(testWord, testValue))

	err = os.Setenv(testWord, testValue)
	Nil(s.T(), err)
	Equal(s.T(), testValue, common.GetEnv(testWord, testValue))
}

func (s *CoreSuite) TestGetEnvInt() {
	Equal(s.T(), common.GetEnvInt(gofakeit.Word(), 1), 1)
	Nil(s.T(), os.Setenv("invalid", "invalid"))
	Equal(s.T(), common.GetEnvInt("invalid", 1), 1)
	Nil(s.T(), os.Setenv("valid", "3"))
	Equal(s.T(), common.GetEnvInt("valid", 1), 3)
}
