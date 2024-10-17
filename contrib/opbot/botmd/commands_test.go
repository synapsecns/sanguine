package botmd_test

import (
	"testing"

	"github.com/synapsecns/sanguine/contrib/opbot/botmd"
)

func TestStripLinks(t *testing.T) {
	testLink := "<https://example.com|example>"
	expected := "example"

	if got := botmd.StripLinks(testLink); got != expected {
		t.Errorf("StripLinks(%s) = %s; want %s", testLink, got, expected)
	}
}
