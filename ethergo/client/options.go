package client

// Options is a type for client options.
type Options func(c *clientImpl)

// Capture option allows for capturing the request and response data.
func Capture(captureReqRes bool) Options {
	return func(c *clientImpl) {
		c.captureRequestRes = captureReqRes
	}
}
