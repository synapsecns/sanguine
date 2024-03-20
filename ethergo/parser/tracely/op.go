package tracely

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/DenrianWeiss/tracely/model"
	"math/big"
	"strconv"
	"strings"
)

func GenSpace(c int) string {
	return strings.Repeat(" ", c)
}

func BytesToInt(b []byte) (int64, error) {
	hexE := hex.EncodeToString(b)
	bI, ok := big.NewInt(0).SetString(hexE, 16)
	if !ok {
		return 0, errors.New("failed to convert bytes to int")
	}
	return int64(bI.Uint64()), nil
}

func ParseCall(step model.TraceStep) (map[string]string, error) {
	result := make(map[string]string)
	stack := GenerateStack(step.Stack)
	memory := GenerateMemory(step.Memory)
	gas := AccessStack(stack, 0)
	addr := AccessStack(stack, 1)
	value := AccessStack(stack, 2)
	argsOffset := AccessStack(stack, 3)
	argsLength := AccessStack(stack, 4)

	// Get Args.
	offset, err := BytesToInt(argsOffset)
	if err != nil {
		return nil, fmt.Errorf("failed to convert args offset to int: %w", err)
	}

	argLength, err := BytesToInt(argsLength)
	if err != nil {
		return nil, fmt.Errorf("failed to convert args length to int: %w", err)
	}

	arg := AccessMemory(memory, int(offset), int(argLength))

	gasCast, err := BytesToInt(gas)
	if err != nil {
		return nil, fmt.Errorf("failed to convert gas to int: %w", err)
	}

	castValue, err := BytesToInt(value)
	if err != nil {
		return nil, fmt.Errorf("failed to convert value to int: %w", err)
	}

	result["gas"] = strconv.FormatUint(uint64(gasCast), 10)
	result["addr"] = "0x" + hex.EncodeToString(addr)[24:]
	result["value"] = strconv.FormatUint(uint64(castValue), 10)
	result["arg"] = hex.EncodeToString(arg)

	return result, nil
}

func ParseDelegateCall(step model.TraceStep) (map[string]string, error) {
	result := make(map[string]string)
	stack := GenerateStack(step.Stack)
	memory := GenerateMemory(step.Memory)
	gas := AccessStack(stack, 0)
	addr := AccessStack(stack, 1)
	argsOffset := AccessStack(stack, 2)
	argsLength := AccessStack(stack, 3)

	castArgOffset, err := BytesToInt(argsOffset)
	if err != nil {
		return nil, fmt.Errorf("failed to convert args offset to int: %w", err)
	}

	castArgLength, err := BytesToInt(argsLength)
	if err != nil {
		return nil, fmt.Errorf("failed to convert args length to int: %w", err)
	}

	castGas, err := BytesToInt(gas)
	if err != nil {
		return nil, fmt.Errorf("failed to convert gas to int: %w", err)
	}

	// Get Args.
	arg := AccessMemory(memory, int(castArgOffset), int(castArgLength))
	result["gas"] = strconv.FormatUint(uint64(castGas), 10)
	result["addr"] = "0x" + hex.EncodeToString(addr)[24:]
	result["arg"] = hex.EncodeToString(arg)
	return result, nil
}

func ParseRevert(step model.TraceStep) (map[string]string, error) {
	result := make(map[string]string)
	stack := GenerateStack(step.Stack)
	memory := GenerateMemory(step.Memory)
	if len(stack) <= 2 {
		return result, nil
	}
	offset := AccessStack(stack, 0)
	length := AccessStack(stack, 1)
	castLength, _ := BytesToInt(length)
	if castLength == 0 {
		return result, nil
	}

	castOffset, err := BytesToInt(offset)
	if err != nil {
		return nil, fmt.Errorf("failed to convert offset to int: %w", err)
	}

	reason := AccessMemory(memory, int(castOffset), int(castLength))
	result["reason"] = string(reason)
	return result, nil
}

func ParseReturn(step model.TraceStep) (map[string]string, error) {
	r, err := ParseRevert(step)
	if err != nil {
		return nil, err
	}
	r["result"] = hex.EncodeToString([]byte(r["reason"]))
	delete(r, "reason")
	return r, nil
}

func PrintStep(step model.TraceStep) (err error) {
	var pCall map[string]string
	if step.Op == "CALL" {
		pCall, err = ParseCall(step)
		fmt.Printf("%sCALL address %s, gas %s, value %s, payload %s\n",
			GenSpace((step.Depth-1)*2),
			pCall["addr"],
			pCall["gas"],
			pCall["value"],
			pCall["arg"])
	} else if step.Op == "DELEGATECALL" || step.Op == "STATICCALL" {
		pCall, err = ParseDelegateCall(step)
		fmt.Printf("%s%s address %s, gas %s, payload %s\n",
			GenSpace((step.Depth-1)*2),
			step.Op,
			pCall["addr"],
			pCall["gas"],
			pCall["arg"])
	} else if step.Op == "REVERT" {
		pCall, err = ParseRevert(step)
		fmt.Printf("%s%s: %s\n",
			GenSpace((step.Depth-1)*2),
			step.Op,
			pCall["reason"],
		)
	} else if step.Op == "RETURN" {
		pCall, err = ParseReturn(step)
		fmt.Printf("%s%s: %s\n",
			GenSpace((step.Depth-1)*2),
			step.Op,
			pCall["result"],
		)
	} else {
		fmt.Printf("%s%s\n", GenSpace((step.Depth-1)*2), step.Op)
	}

	if err != nil {
		fmt.Println(err)
	}
	return
}
