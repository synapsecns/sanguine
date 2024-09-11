package metrics

// HeadersToMap converts a string of headers to a map.
func HeadersToMap(val string) map[string]string {
	return headersToMap(val)
}
