package botmd_test

import (
	"github.com/synapsecns/sanguine/contrib/opbot/botmd"
	"testing"
)

func TestStripLinks(t *testing.T) {
	testLink := "<https://example.com|example>"
	expected := "example"

	if got := botmd.StripLinks(testLink); got != expected {
		t.Errorf("StripLinks(%s) = %s; want %s", testLink, got, expected)
	}
}
