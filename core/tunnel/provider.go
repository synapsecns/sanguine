package tunnel

// Provider is a tunnel provider.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=Provider -linecomment
type Provider uint8

const (
	// Moe is the moe provider: https://github.com/fasmide/remotemoe
	Moe Provider = iota + 1
	// Ngrok is the ngrok provider: https://ngrok.com/
	Ngrok
)
