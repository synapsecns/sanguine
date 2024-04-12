package tracely

import (
	"fmt"
	"strings"
)

func SeperatePayload(in *string) {
	s := strings.TrimPrefix(*in, "0x")
	sLen := len(s)
	if sLen%64 != 0 {
		sSig := s[:8]
		fmt.Printf("function sig: %s\n", sSig)
		s = s[8:]
	}
	for len(s) >= 64 {
		sPart := s[:64]
		s = s[64:]
		fmt.Printf("%s\n", sPart)
	}
}
