package tracely

import (
	"encoding/hex"
	"fmt"
	"github.com/DenrianWeiss/tracely/model"
	"math/big"
	"strconv"
	"strings"
)

func GenSpace(c int) string {
	return strings.Repeat(" ", c)
}

func BytesToInt(b []byte) int64 {
	hexE := hex.EncodeToString(b)
	bI, _ := big.NewInt(0).SetString(hexE, 16)
	return int64(bI.Uint64())
}

func ParseCall(step model.TraceStep) map[string]string {
	result := make(map[string]string)
	stack := GenerateStack(step.Stack)
	memory := GenerateMemory(step.Memory)
	gas := AccessStack(stack, 0)
	addr := AccessStack(stack, 1)
	value := AccessStack(stack, 2)
	argsOffset := AccessStack(stack, 3)
	argsLength := AccessStack(stack, 4)

	// Get Args.
	arg := AccessMemory(memory, int(BytesToInt(argsOffset)), int(BytesToInt(argsLength)))
	result["gas"] = strconv.FormatUint(uint64(BytesToInt(gas)), 10)
	result["addr"] = "0x" + hex.EncodeToString(addr)[24:]
	result["value"] = strconv.FormatUint(uint64(BytesToInt(value)), 10)
	result["arg"] = hex.EncodeToString(arg)
	return result
}

func ParseDelegateCall(step model.TraceStep) map[string]string {
	result := make(map[string]string)
	stack := GenerateStack(step.Stack)
	memory := GenerateMemory(step.Memory)
	gas := AccessStack(stack, 0)
	addr := AccessStack(stack, 1)
	argsOffset := AccessStack(stack, 2)
	argsLength := AccessStack(stack, 3)

	// Get Args.
	arg := AccessMemory(memory, int(BytesToInt(argsOffset)), int(BytesToInt(argsLength)))
	result["gas"] = strconv.FormatUint(uint64(BytesToInt(gas)), 10)
	result["addr"] = "0x" + hex.EncodeToString(addr)[24:]
	result["arg"] = hex.EncodeToString(arg)
	return result
}

func ParseRevert(step model.TraceStep) map[string]string {
	result := make(map[string]string)
	stack := GenerateStack(step.Stack)
	memory := GenerateMemory(step.Memory)
	if len(stack) <= 2 {
		return result
	}
	offset := AccessStack(stack, 0)
	length := AccessStack(stack, 1)
	if BytesToInt(length) == 0 {
		return result
	}
	reason := AccessMemory(memory, int(BytesToInt(offset)), int(BytesToInt(length)))
	result["reason"] = string(reason)
	return result
}

func ParseReturn(step model.TraceStep) map[string]string {
	r := ParseRevert(step)
	r["result"] = hex.EncodeToString([]byte(r["reason"]))
	delete(r, "reason")
	return r
}

func PrintStep(step model.TraceStep) {
	if step.Op == "CALL" {
		pCall := ParseCall(step)
		fmt.Printf("%sCALL address %s, gas %s, value %s, payload %s\n",
			GenSpace((step.Depth-1)*2),
			pCall["addr"],
			pCall["gas"],
			pCall["value"],
			pCall["arg"])
	} else if step.Op == "DELEGATECALL" || step.Op == "STATICCALL" {
		pCall := ParseDelegateCall(step)
		fmt.Printf("%s%s address %s, gas %s, payload %s\n",
			GenSpace((step.Depth-1)*2),
			step.Op,
			pCall["addr"],
			pCall["gas"],
			pCall["arg"])
	} else if step.Op == "REVERT" {
		pCall := ParseRevert(step)
		fmt.Printf("%s%s: %s\n",
			GenSpace((step.Depth-1)*2),
			step.Op,
			pCall["reason"],
		)
	} else if step.Op == "RETURN" {
		pCall := ParseReturn(step)
		fmt.Printf("%s%s: %s\n",
			GenSpace((step.Depth-1)*2),
			step.Op,
			pCall["result"],
		)
	} else {
		fmt.Printf("%s%s\n", GenSpace((step.Depth-1)*2), step.Op)
	}
}
