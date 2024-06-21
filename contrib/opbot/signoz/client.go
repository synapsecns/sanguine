package signoz

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Start Generation Here
type Client struct {
	apiKey string
	client *resty.Client
}

func NewClient(url, apiKey string) *Client {
	client := resty.New()
	client.SetBaseURL(fmt.Sprintf("https://%s/api/v3/query_range", apiKey))
	client.SetHeader("SIGNOZ-API-KEY", apiKey)
	return &Client{
		apiKey: apiKey,
		client: client,
	}
}

func (c *Client) SearchTraces(deploymentName, httpMethod string, hasError bool, tags map[string]string) (*resty.Response, error) {
	queryParams := map[string]string{
		"deployment_name": deploymentName,
		"httpMethod":      httpMethod,
		"hasError":        map[bool]string{true: "true", false: "false"}[hasError],
	}

	for key, value := range tags {
		queryParams[key] = value
	}

	return c.client.R().
		SetQueryParams(queryParams).
		Get("/query_range")
}
