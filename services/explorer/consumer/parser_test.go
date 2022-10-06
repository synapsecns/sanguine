package consumer

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestBoolToUint8(t *testing.T) {
	inputTrue := true
	inputFalse := false
	checkTrue := BoolToUint8(&inputTrue)
	checkFalse := BoolToUint8(&inputFalse)
	Equal(t, uint8(1), *checkTrue)
	Equal(t, uint8(0), *checkFalse)
}

func TestToNullString(t *testing.T) {
	inputValid := "TEST_STRING"
	checkValid := ToNullString(&inputValid)
	checkInvalid := ToNullString(nil)
	Equal(t, inputValid, checkValid.String)
	True(t, checkValid.Valid)
	Equal(t, "", checkInvalid.String)
	False(t, checkInvalid.Valid)
}

// TODO after mock bridge emits events.
