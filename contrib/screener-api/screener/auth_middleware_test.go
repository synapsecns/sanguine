package screener

import (
	"testing"
	"time"
)

func TestSignatureTimestampFresh(t *testing.T) {
	t.Parallel()

	now := time.Unix(1_700_000_000, 0)
	maxSkew := 5 * time.Minute

	tests := []struct {
		name      string
		timestamp string
		want      bool
	}{
		{name: "exact now", timestamp: "1700000000", want: true},
		{name: "within skew past", timestamp: "1699999900", want: true},
		{name: "within skew future", timestamp: "1700000100", want: true},
		{name: "stale past hour", timestamp: "1699996400", want: false},
		{name: "far future", timestamp: "1700003600", want: false},
		{name: "non numeric", timestamp: "not-a-unix-ts", want: false},
		{name: "empty", timestamp: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := signatureTimestampFresh(tt.timestamp, now, maxSkew); got != tt.want {
				t.Fatalf("signatureTimestampFresh(%q) = %v, want %v", tt.timestamp, got, tt.want)
			}
		})
	}
}
