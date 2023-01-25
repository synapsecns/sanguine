package provider

import (
	"fmt"
)

// TODO: test.
func validatePort(v interface{}, k string) (ws []string, errors []error) {
	value, ok := v.(int)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %s to be int", k))
		return
	}
	if value < 1 || value > 65535 {
		errors = append(errors, fmt.Errorf("%q must be between 1 and 65535, got: %d", k, value))
	}
	return
}
