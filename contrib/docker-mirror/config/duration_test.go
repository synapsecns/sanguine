package config_test

import (
	"github.com/synapsecns/sanguine/contrib/docker-mirror/config"
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	tests := []struct {
		input string
		want  time.Duration
	}{
		{"1y", time.Duration(1000*60*60*24*365) * time.Millisecond},
		{"1w", time.Duration(1000*60*60*24*7) * time.Millisecond},
		{"1d", time.Duration(1000*60*60*24) * time.Millisecond},
		{"1h", time.Duration(1000*60*60) * time.Millisecond},
		{"1m", time.Duration(1000*60) * time.Millisecond},
		{"1s", time.Duration(1000) * time.Millisecond},
		{"1ms", time.Duration(1) * time.Millisecond},
		{"10ms", time.Duration(10) * time.Millisecond},
		{"10s", time.Duration(10000) * time.Millisecond},
		{"10m", time.Duration(600000) * time.Millisecond},
		{"10h", time.Duration(36000000) * time.Millisecond},
		{"10d", time.Duration(864000000) * time.Millisecond},
		{"10w", time.Duration(6048000000) * time.Millisecond},
		{"10y", time.Duration(315360000000) * time.Millisecond},
	}
	for i, test := range tests {
		got, err := config.ParseDuration(test.input)
		if err != nil {
			t.Errorf("Test %d: unexpected error: %v", i, err)
			continue
		}
		if got != config.Duration(test.want) {
			t.Errorf("Test %d: got %v, want %v", i, got, test.want)
		}
	}
}

func TestDurationString(t *testing.T) {
	tests := []struct {
		input time.Duration
		want  string
	}{
		{time.Duration(1000*60*60*24*365) * time.Millisecond, "1y"},
		{time.Duration(1000*60*60*24*7) * time.Millisecond, "1w"},
		{time.Duration(1000*60*60*24) * time.Millisecond, "1d"},
		{time.Duration(1000*60*60) * time.Millisecond, "1h"},
		{time.Duration(1000*60) * time.Millisecond, "1m"},
		{time.Duration(1000) * time.Millisecond, "1s"},
		{time.Duration(1) * time.Millisecond, "1ms"},
		{time.Duration(10) * time.Millisecond, "10ms"},
		{time.Duration(10000) * time.Millisecond, "10s"},
		{time.Duration(600000) * time.Millisecond, "10m"},
		{time.Duration(36000000) * time.Millisecond, "10h"},
		{time.Duration(864000000) * time.Millisecond, "10d"},
		{time.Duration(6048000000) * time.Millisecond, "10w"},
		{time.Duration(315360000000) * time.Millisecond, "10y"},
	}
	for i, test := range tests {
		got := config.Duration(test.input).String()
		if got != test.want {
			t.Errorf("Test %d: got %v, want %v", i, got, test.want)
		}
	}
}
