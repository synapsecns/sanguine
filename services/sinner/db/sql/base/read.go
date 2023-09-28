package base

import (
	"context"
	"fmt"
)

// RetrieveMessageStatus retrieve message status.
func (s Store) RetrieveMessageStatus(_ context.Context, txhash string) (string, error) {
	// TODO implement
	fmt.Println("gm", txhash)
	return "gm", nil
}
