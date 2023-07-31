package core_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	common "github.com/synapsecns/sanguine/core"
	"os"
	"testing"
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

func TestGetEnvBool(t *testing.T) {
	type args struct {
		name       string
		defaultVal bool
	}
	tests := []struct {
		name     string
		args     args
		want     bool
		envVal   string
		setupEnv bool
	}{
		{
			name: "Environment variable not set",
			args: args{
				name:       "NOT_SET",
				defaultVal: true,
			},
			want:     true,
			setupEnv: false,
		},
		{
			name: "Environment variable set to true",
			args: args{
				name:       "SET_TRUE",
				defaultVal: false,
			},
			want:     true,
			envVal:   "true",
			setupEnv: true,
		},
		{
			name: "Environment variable set to false",
			args: args{
				name:       "SET_FALSE",
				defaultVal: true,
			},
			want:     false,
			envVal:   "false",
			setupEnv: true,
		},
		{
			name: "Environment variable set to non-boolean",
			args: args{
				name:       "SET_NON_BOOLEAN",
				defaultVal: true,
			},
			want:     true,
			envVal:   "non-boolean",
			setupEnv: true,
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupEnv {
				t.Setenv(tt.args.name, tt.envVal)
			} else {
				_ = os.Unsetenv(tt.args.name)
			}

			if got := common.GetEnvBool(tt.args.name, tt.args.defaultVal); got != tt.want {
				t.Errorf("GetEnvBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
