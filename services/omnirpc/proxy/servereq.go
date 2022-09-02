package proxy

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	resty "github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func parseRPCPayload(body []byte) (method string, err error) {
	rpcPayload := struct {
		ID     json.RawMessage   `json:"id"`
		Method string            `json:"method"`
		Params []json.RawMessage `json:"params"`
	}{}

	err = json.Unmarshal(body, &rpcPayload)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse json RPC payload")
	}

	return rpcPayload.Method, nil
}

func (r *RPCProxy) serveRPCReq(c *gin.Context, chainID int) {
	rpcList := r.rpcMap.ChainID(chainID)

	if len(rpcList) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("no endpoint for chain %d", chainID),
		})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}

	for _, endpoint := range rpcList {
		endpointURL, err := url.Parse(endpoint)
		if err != nil {
			continue
		}

		// websockets not yet supported
		if !slices.Contains([]string{"http", "https"}, endpointURL.Scheme) {
			continue
		}

		client := resty.New()
		resp, err := client.R().
			SetContext(c).
			SetBody(body).
			Post(endpoint)

		if err != nil {
			// continue until we exhaust endpoints
			continue
		}

		if resp.StatusCode() < 200 || resp.StatusCode() > 400 {
			// error
			continue
		}

		c.Header("x-forwarded-from", endpoint)
		c.Data(http.StatusOK, gin.MIMEJSON, resp.Body())
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": fmt.Sprintf("no rpc online for chain %d, attempted: %s", chainID, strings.Join(rpcList, ",")),
	})
}
