package decoders

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/lmittmann/w3"
)

var (
	revertSelector       = selector("Error(string)")
	approveSelector      = selector("approve(address,uint256)")
	transferSelector     = selector("transfer(address,uint256)")
	transferFromSelector = selector("transferFrom(address,address,uint256)")
	outputSuccess        = w3.B("0x0000000000000000000000000000000000000000000000000000000000000001")
)

// selector returns the 4-byte selector of the given signature.
func selector(signature string) (selector [4]byte) {
	copy(selector[:], crypto.Keccak256([]byte(signature)))
	return
}

// reimplements standard DecodeReturn checks from w3.
func returnChecks(f *w3.Func, output []byte, returns ...any) error {
	if bytes.HasPrefix(output, revertSelector[:]) {
		//nolint
		if reason, err := abi.UnpackRevert(output); err != nil {
			// nolint
			return err
		} else {
			return fmt.Errorf("%w: %s", w3.ErrEvmRevert, reason)
		}
	}

	// Gracefully handle uncompliant ERC20 returns
	if len(returns) == 1 && len(output) == 0 &&
		(f.Selector == approveSelector ||
			f.Selector == transferSelector ||
			f.Selector == transferFromSelector) {
		// nolint
		// TODO: deal with ineffassign here. Right now this isn't used by any of the functions
		output = outputSuccess
	}

	return nil
}
