package parser_test

import (
	"github.com/synapsecns/sanguine/services/explorer/consumer/parsers"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestBoolToUint8(t *testing.T) {
	inputTrue := true
	inputFalse := false
	checkTrue := parser.BoolToUint8(&inputTrue)
	checkFalse := parser.BoolToUint8(&inputFalse)
	Equal(t, uint8(1), *checkTrue)
	Equal(t, uint8(0), *checkFalse)
}

func TestToNullString(t *testing.T) {
	inputValid := "TEST_STRING"
	checkValid := parser.ToNullString(&inputValid)
	checkInvalid := parser.ToNullString(nil)
	Equal(t, inputValid, checkValid.String)
	True(t, checkValid.Valid)
	Equal(t, "", checkInvalid.String)
	False(t, checkInvalid.Valid)
}
