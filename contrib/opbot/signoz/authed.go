package signoz

import (
	"context"
	"fmt"
	"github.com/dubonzi/otelresty"
	"github.com/go-resty/resty/v2"
	v3 "github.com/synapsecns/sanguine/contrib/opbot/signoz/generated/v3"
	"github.com/synapsecns/sanguine/core/metrics"
	"time"
)

// Client is a signoz client.
type Client struct {
	*UnauthenticatedClient
	handler                     metrics.Handler
	client                      *resty.Client
	url                         string
	email, pass                 string
	bearerToken, refreshToken   string
	bearerExpiry, refreshExpiry int64
}

// NewClientFromUser creates a new signoz client from a user.
func NewClientFromUser(handler metrics.Handler, url, email, password string) *Client {
	if len(url) == 0 {
		panic("url is required")
	}

	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}

	res := &Client{
		UnauthenticatedClient: NewUnauthenticatedClient(handler, url),
		url:                   url,
		handler:               handler,
		client:                resty.New(),
		email:                 email,
		pass:                  password,
	}

	res.client.SetBaseURL(url)
	res.client.OnBeforeRequest(res.beforeRequest)
	otelresty.TraceClient(res.client, otelresty.WithTracerProvider(res.handler.GetTracerProvider()))

	return res
}

func (c *Client) beforeRequest(_ *resty.Client, request *resty.Request) error {
	// if bearer token & refresh token are not set, login
	// TODO: we haven't built bearer refresh so just re-login every time
	if (c.bearerToken == "" && c.refreshToken == "") || c.bearerExpiry < time.Now().Unix() {
		res, err := c.UnauthenticatedClient.Login(request.Context(), c.email, c.pass)
		if err != nil {
			return err
		}

		c.bearerToken = res.AccessJwt
		c.refreshToken = res.RefreshJwt
		c.bearerExpiry = int64(res.AccessJwtExpiry)
		c.refreshExpiry = int64(res.AccessJwtExpiry)
	}

	// nolint: revive, staticcheck
	if c.bearerExpiry < time.Now().Unix() {
		// TODO: implement refresh token logic and change above to be c.refreshExpiry
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	return nil
}

// ServiceResponse is a signoz service.
type ServiceResponse struct {
	ServiceName string  `json:"serviceName"`
	P99         float64 `json:"p99"`
	AvgDuration float64 `json:"avgDuration"`
	NumCalls    int     `json:"numCalls"`
	CallRate    float64 `json:"callRate"`
	NumErrors   int     `json:"numErrors"`
	ErrorRate   float64 `json:"errorRate"`
	Num4XX      float64 `json:"num4XX"`
	FourXXRate  float64 `json:"fourXXRate"`
}

// ServiceRequest is a request to get services.
type ServiceRequest struct {
	Start string        `json:"start"`
	End   string        `json:"end"`
	Tags  []interface{} `json:"tags"`
}

// Services gets services.
func (c *Client) Services(ctx context.Context, timePeriod TimePreferenceType) (res []ServiceResponse, err error) {
	param := GetStartAndEndTime(timePeriod)

	resp, err := c.client.R().
		SetBody(ServiceRequest{
			Start: param.Start,
			End:   param.End,
			Tags:  []interface{}{},
		}).
		SetContext(ctx).
		SetResult(&res).
		Post("/api/v1/services")
	if err != nil {
		return nil, fmt.Errorf("error getting services %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("error getting services %w", err)
	}

	return res, nil
}

func hasField(field string, fields []v3.AttributeKey) bool {
	for _, f := range fields {
		if f.Key == field {
			return true
		}
	}

	return false
}

// QueryRangeResponse is a response to a query range.
type QueryRangeResponse struct {
	Status string                 `json:"status"`
	Data   *v3.QueryRangeResponse `json:"data"`
}

// SearchTraces searches for traces.
func (c *Client) SearchTraces(ctx context.Context, timePeriod TimePreferenceType, searchTerms map[string]string) (*QueryRangeResponse, error) {
	param, err := GetStartAndEndTimeInt(timePeriod)
	if err != nil {
		return nil, err
	}

	columns := []v3.AttributeKey{
		{
			Key:      "serviceName",
			DataType: v3.AttributeKeyDataTypeString,
			Type:     v3.AttributeKeyTypeTag,
			IsColumn: true,
		},
		{
			Key:      "httpMethod",
			DataType: v3.AttributeKeyDataTypeString,
			Type:     v3.AttributeKeyTypeTag,
			IsColumn: true,
		},
		{
			Key:      "responseStatusCode",
			DataType: v3.AttributeKeyDataTypeString,
			Type:     v3.AttributeKeyTypeTag,
			IsColumn: true,
		},
		{
			Key:      "httpUrl",
			DataType: v3.AttributeKeyDataTypeString,
			Type:     v3.AttributeKeyTypeTag,
			IsColumn: true,
		},
		{
			Key:      "name",
			DataType: v3.AttributeKeyDataTypeString,
			Type:     v3.AttributeKeyTypeTag,
			IsColumn: true,
		},
	}

	filterItems := make([]v3.FilterItem, 0, len(searchTerms))
	for key, value := range searchTerms {
		filterItems = append(filterItems, v3.FilterItem{
			Key: v3.AttributeKey{
				Key:      key,
				DataType: v3.AttributeKeyDataTypeString,
				Type:     v3.AttributeKeyTypeTag,
				IsColumn: hasField(key, columns),
				IsJSON:   false,
			},
			Operator: "=",
			Value:    value,
		})
	}

	query := v3.QueryRangeParamsV3{
		Start: param.Start,
		End:   param.End,
		Step:  60,
		CompositeQuery: &v3.CompositeQuery{
			BuilderQueries: map[string]*v3.BuilderQuery{
				"A": {
					DataSource:        v3.DataSourceTraces,
					AggregateOperator: v3.AggregateOperatorNoOp,
					Filters: &v3.FilterSet{
						Operator: "AND",
						// TODO: add filters
						Items: filterItems,
					},
					// TODO
					Limit:        10,
					Offset:       0,
					Expression:   "A",
					QueryName:    "A",
					StepInterval: 10,
					ReduceTo:     v3.ReduceToOperatorSum,
					OrderBy: []v3.OrderBy{
						{
							ColumnName: "timestamp",
							Order:      "desc",
						},
					},
					SelectColumns: columns,
				},
			},
			PanelType: v3.PanelTypeList,
			QueryType: v3.QueryTypeBuilder,
		},
	}

	var res QueryRangeResponse

	resp, err := c.client.R().
		SetContext(ctx).
		SetBody(query).
		SetResult(&res).
		Post("/api/v3/query_range")

	if err != nil {
		return nil, fmt.Errorf("error getting traces %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("error getting traces %w", err)
	}

	return &res, nil
}
