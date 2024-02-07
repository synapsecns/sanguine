package api

import (
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func getRawParam(name string, ctx *gin.Context) (string, error) {
	value := ctx.Query(name)
	if value == "" {
		return "", fmt.Errorf("required parameter '%s' is missing", name)
	}
	return value, nil
}

const originParamName = "origin"
const hashParamName = "hash"

func getOriginParam(ctx *gin.Context) (uint32, error) {
	rawValue, err := getRawParam(originParamName, ctx)
	if err != nil {
		return 0, err
	}
	value, err := strconv.Atoi(rawValue)
	if err != nil {
		return 0, fmt.Errorf("could not parse origin: %s", rawValue)
	}
	return uint32(value), nil
}

func getHashParam(ctx *gin.Context) (string, error) {
	value, err := getRawParam(hashParamName, ctx)
	if err != nil {
		return "", err
	}
	value = common.HexToHash(value).String()
	return value, nil
}
