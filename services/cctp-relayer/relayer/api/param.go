package api

import (
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type Param interface {
	Name() string
	Parse(ctx *gin.Context) (interface{}, error)
}

func getRawParam(name string, ctx *gin.Context) (string, error) {
	value := ctx.Param(name)
	if value == "" {
		return "", fmt.Errorf("required parameter '%s' is missing", name)
	}
	return value, nil
}

type HashParam struct{}

func (h HashParam) Name() string { return "hash" }

func (h HashParam) Parse(ctx *gin.Context) (interface{}, error) {
	value, err := getRawParam(h.Name(), ctx)
	if err != nil {
		return "", err
	}
	ok := common.IsHexAddress(value)
	if !ok {
		return "", fmt.Errorf("invalid hash: %s", value)
	}
	return value, nil
}

type OriginParam struct{}

func (o OriginParam) Name() string { return "origin" }

func (o OriginParam) Parse(ctx *gin.Context) (interface{}, error) {
	rawValue, err := getRawParam(o.Name(), ctx)
	if err != nil {
		return 0, err
	}
	value, err := strconv.Atoi(rawValue)
	if err != nil {
		return 0, fmt.Errorf("could not parse origin: %s", rawValue)
	}
	return uint32(value), nil
}
