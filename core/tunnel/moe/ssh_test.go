package moe_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/synapsecns/sanguine/core/tunnel/moe"
	"strings"
	"testing"
)

func TestConsumeBufferUntilURL(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "url found",
			input:   "some text http://example.com/ some more text " + moe.MoeServer,
			want:    "http://example.com/",
			wantErr: false,
		},
		{
			name:    "url not found",
			input:   "some text" + moe.MoeServer + "some more text",
			want:    "",
			wantErr: true,
		},
		{
			name:    "context canceled",
			input:   "some text",
			want:    "",
			wantErr: true,
		},
	}

	for i := range testCases {
		tc := testCases[i] // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			got, err := moe.ConsumeBufferUntilURL(ctx, reader)

			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tc.want, got)
		})
	}
}
