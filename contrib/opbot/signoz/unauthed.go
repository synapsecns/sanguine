package signoz

import (
	"context"
	"fmt"
	"github.com/dubonzi/otelresty"
	"github.com/go-http-utils/headers"
	"github.com/go-resty/resty/v2"
	"github.com/synapsecns/sanguine/core/metrics"
)

// UnauthenticatedClient is an unauthenticated client.
type UnauthenticatedClient struct {
	client *resty.Client
}

// NewUnauthenticatedClient creates a new unauthenticated client.
func NewUnauthenticatedClient(handler metrics.Handler, url string) *UnauthenticatedClient {
	client := resty.New()
	client.SetBaseURL(url)
	client.SetHeader(headers.UserAgent, "query-service")
	otelresty.TraceClient(client, otelresty.WithTracerProvider(handler.GetTracerProvider()))

	return &UnauthenticatedClient{
		client: client,
	}
}

// Login logs in a user.
func (c *UnauthenticatedClient) Login(ctx context.Context, email, password string) (_ *LoginResponse, err error) {
	type RequestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var res LoginResponse

	resp, err := c.client.R().
		SetBody(RequestBody{Email: email, Password: password}).
		SetContext(ctx).
		SetResult(&res).
		Post("/api/v1/login")
	if err != nil {
		return nil, fmt.Errorf("error logging in: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("login failed: %w", err)
	}

	return &res, nil
}

// LoginResponse is the response from the login endpoint.
type LoginResponse struct {
	// UserID is the user ID.
	UserID string `json:"userId"`
	// AccessJWT is the access JWT.
	AccessJwt string `json:"accessJwt"`
	// AccessJWTExpiry is the access JWT expiry.
	AccessJwtExpiry int `json:"accessJwtExpiry"`
	// RefreshJWT is the refresh JWT.
	RefreshJwt string `json:"refreshJwt"`
	// RefreshJWTExpiry is the refresh JWT expiry.
	RefreshJwtExpiry int `json:"refreshJwtExpiry"`
}
