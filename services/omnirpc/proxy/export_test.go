package proxy

// IsConfirmable exports isConfirmable for testing.
func IsConfirmable(body []byte) (bool, error) {
	return isConfirmable(body)
}

// RawResponse exports rawResponse for testing.
type RawResponse interface {
	Body() []byte
	URL() string
	Hash() string
}

func (r rawResponse) Body() []byte {
	return r.body
}

func (r rawResponse) URL() string {
	return r.url
}

func (r rawResponse) Hash() string {
	return r.hash
}

var _ RawResponse = rawResponse{}

func NewRawResponse(body []byte, url string) (RawResponse, error) {
	return newRawResponse(body, url)
}
