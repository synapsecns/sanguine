// Package provider wraps the upstream Grafana Terraform provider with AMG auth.
package provider

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/grafana"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

const (
	envGrafanaAuth = "GRAFANA_AUTH"

	envWorkspaceID      = "GRAFANA_AMG_WORKSPACE_ID"
	envServiceAccountID = "GRAFANA_AMG_SERVICE_ACCOUNT_ID"
	envRegion           = "GRAFANA_AMG_REGION"
	envTokenTTLSeconds  = "GRAFANA_AMG_TOKEN_TTL_SECONDS"    // #nosec G101 -- environment variable name, not a credential.
	envTokenNamePrefix  = "GRAFANA_AMG_TOKEN_NAME_PREFIX"    // #nosec G101 -- environment variable name, not a credential.
	envDeleteToken      = "GRAFANA_AMG_DELETE_TOKEN_ON_STOP" // #nosec G101 -- environment variable name, not a credential.

	defaultTokenTTLSeconds = int32(3600)
	defaultTokenNamePrefix = "terraform-provider-grafanaamg" // #nosec G101 -- token name prefix, not a credential.
)

// Server delegates Terraform provider RPCs to the upstream Grafana provider.
type Server struct {
	tfprotov5.ProviderServer

	mu     sync.Mutex
	tokens []createdToken
}

type createdToken struct {
	workspaceID      string
	serviceAccountID string
	tokenID          string
	region           string
}

type amgConfig struct {
	workspaceID      string
	serviceAccountID string
	region           string
	tokenTTLSeconds  int32
	tokenNamePrefix  string
	deleteToken      bool
}

type grafanaAPI interface {
	CreateWorkspaceServiceAccountToken(ctx context.Context, input *grafana.CreateWorkspaceServiceAccountTokenInput, optFns ...func(*grafana.Options)) (*grafana.CreateWorkspaceServiceAccountTokenOutput, error)
	DeleteWorkspaceServiceAccountToken(ctx context.Context, input *grafana.DeleteWorkspaceServiceAccountTokenInput, optFns ...func(*grafana.Options)) (*grafana.DeleteWorkspaceServiceAccountTokenOutput, error)
}

var newGrafanaClient = grafanaClient

// New wraps an upstream Terraform provider server with AMG token configuration.
func New(upstream tfprotov5.ProviderServer) *Server {
	return &Server{ProviderServer: upstream}
}

// ConfigureProvider mints an AMG token before configuring the upstream provider.
func (s *Server) ConfigureProvider(ctx context.Context, req *tfprotov5.ConfigureProviderRequest) (*tfprotov5.ConfigureProviderResponse, error) {
	cfg, enabled, err := loadAMGConfig()
	if err != nil {
		return diagnosticResponse("Invalid AMG provider configuration.", err.Error()), nil
	}

	if !enabled || strings.TrimSpace(os.Getenv(envGrafanaAuth)) != "" {
		resp, configureErr := s.ProviderServer.ConfigureProvider(ctx, req)
		if configureErr != nil {
			return resp, fmt.Errorf("configure upstream Grafana provider: %w", configureErr)
		}

		return resp, nil
	}

	token, err := s.createToken(ctx, cfg)
	if err != nil {
		return diagnosticResponse("Failed to create Amazon Managed Grafana token.", err.Error()), nil
	}

	previousAuth, hadPreviousAuth := os.LookupEnv(envGrafanaAuth)
	setEnvErr := os.Setenv(envGrafanaAuth, token.key)
	if setEnvErr != nil {
		return diagnosticResponse("Failed to configure Grafana provider auth.", setEnvErr.Error()), nil
	}
	defer restoreEnv(envGrafanaAuth, previousAuth, hadPreviousAuth)

	resp, err := s.ProviderServer.ConfigureProvider(ctx, req)
	if err != nil {
		s.deleteTokenIfEnabled(ctx, cfg, token.id)
		return resp, fmt.Errorf("configure upstream Grafana provider: %w", err)
	}

	if hasErrors(resp) {
		s.deleteTokenIfEnabled(ctx, cfg, token.id)
		return resp, nil
	}

	if cfg.deleteToken {
		s.mu.Lock()
		s.tokens = append(s.tokens, createdToken{
			workspaceID:      cfg.workspaceID,
			serviceAccountID: cfg.serviceAccountID,
			tokenID:          token.id,
			region:           cfg.region,
		})
		s.mu.Unlock()
	}

	return resp, nil
}

// StopProvider stops the upstream provider and cleans up minted AMG tokens.
func (s *Server) StopProvider(ctx context.Context, req *tfprotov5.StopProviderRequest) (*tfprotov5.StopProviderResponse, error) {
	resp, err := s.ProviderServer.StopProvider(ctx, req)

	s.mu.Lock()
	tokens := append([]createdToken(nil), s.tokens...)
	s.tokens = nil
	s.mu.Unlock()

	for _, token := range tokens {
		deleteErr := deleteToken(ctx, token)
		if deleteErr != nil {
			log.Printf("failed to delete AMG Grafana service account token %s: %v", token.tokenID, deleteErr)
		}
	}

	if err != nil {
		return resp, fmt.Errorf("stop upstream Grafana provider: %w", err)
	}

	return resp, nil
}

func (s *Server) deleteTokenIfEnabled(ctx context.Context, cfg amgConfig, tokenID string) {
	if !cfg.deleteToken {
		return
	}

	token := createdToken{
		workspaceID:      cfg.workspaceID,
		serviceAccountID: cfg.serviceAccountID,
		tokenID:          tokenID,
		region:           cfg.region,
	}

	err := deleteToken(ctx, token)
	if err != nil {
		log.Printf("failed to delete AMG Grafana service account token %s after failed provider configuration: %v", tokenID, err)
	}
}

type mintedToken struct {
	id  string
	key string
}

func (s *Server) createToken(ctx context.Context, cfg amgConfig) (mintedToken, error) {
	client, err := newGrafanaClient(ctx, cfg.region)
	if err != nil {
		return mintedToken{}, err
	}

	name := fmt.Sprintf("%s-%d", cfg.tokenNamePrefix, time.Now().UnixNano())

	out, err := client.CreateWorkspaceServiceAccountToken(ctx, &grafana.CreateWorkspaceServiceAccountTokenInput{
		Name:             aws.String(name),
		SecondsToLive:    aws.Int32(cfg.tokenTTLSeconds),
		ServiceAccountId: aws.String(cfg.serviceAccountID),
		WorkspaceId:      aws.String(cfg.workspaceID),
	})
	if err != nil {
		return mintedToken{}, fmt.Errorf("create AMG workspace service account token: %w", err)
	}

	if out.ServiceAccountToken == nil || aws.ToString(out.ServiceAccountToken.Key) == "" || aws.ToString(out.ServiceAccountToken.Id) == "" {
		return mintedToken{}, fmt.Errorf("AWS returned an empty AMG service account token")
	}

	return mintedToken{
		id:  aws.ToString(out.ServiceAccountToken.Id),
		key: aws.ToString(out.ServiceAccountToken.Key),
	}, nil
}

func deleteToken(ctx context.Context, token createdToken) error {
	client, err := newGrafanaClient(ctx, token.region)
	if err != nil {
		return err
	}

	_, err = client.DeleteWorkspaceServiceAccountToken(ctx, &grafana.DeleteWorkspaceServiceAccountTokenInput{
		ServiceAccountId: aws.String(token.serviceAccountID),
		TokenId:          aws.String(token.tokenID),
		WorkspaceId:      aws.String(token.workspaceID),
	})
	if err != nil {
		return fmt.Errorf("delete AMG workspace service account token: %w", err)
	}

	return nil
}

func grafanaClient(ctx context.Context, region string) (grafanaAPI, error) {
	opts := []func(*config.LoadOptions) error{}
	if strings.TrimSpace(region) != "" {
		opts = append(opts, config.WithRegion(region))
	}

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("load AWS config: %w", err)
	}

	return grafana.NewFromConfig(cfg), nil
}

func loadAMGConfig() (amgConfig, bool, error) {
	workspaceID := strings.TrimSpace(os.Getenv(envWorkspaceID))
	serviceAccountID := strings.TrimSpace(os.Getenv(envServiceAccountID))

	if workspaceID == "" && serviceAccountID == "" {
		return amgConfig{}, false, nil
	}

	if workspaceID == "" || serviceAccountID == "" {
		return amgConfig{}, false, fmt.Errorf("%s and %s must both be set", envWorkspaceID, envServiceAccountID)
	}

	ttl, err := parseTokenTTL()
	if err != nil {
		return amgConfig{}, false, err
	}

	deleteOnStop, err := parseDeleteOnStop()
	if err != nil {
		return amgConfig{}, false, err
	}

	prefix := strings.TrimSpace(os.Getenv(envTokenNamePrefix))
	if prefix == "" {
		prefix = defaultTokenNamePrefix
	}

	return amgConfig{
		workspaceID:      workspaceID,
		serviceAccountID: serviceAccountID,
		region:           strings.TrimSpace(os.Getenv(envRegion)),
		tokenTTLSeconds:  ttl,
		tokenNamePrefix:  prefix,
		deleteToken:      deleteOnStop,
	}, true, nil
}

func parseTokenTTL() (int32, error) {
	rawTTL := strings.TrimSpace(os.Getenv(envTokenTTLSeconds))
	if rawTTL == "" {
		return defaultTokenTTLSeconds, nil
	}

	parsedTTL, err := strconv.ParseInt(rawTTL, 10, 32)
	if err != nil || parsedTTL <= 0 {
		return 0, fmt.Errorf("%s must be a positive integer", envTokenTTLSeconds)
	}

	return int32(parsedTTL), nil
}

func parseDeleteOnStop() (bool, error) {
	rawDelete := strings.TrimSpace(os.Getenv(envDeleteToken))
	if rawDelete == "" {
		return true, nil
	}

	parsedDelete, err := strconv.ParseBool(rawDelete)
	if err != nil {
		return false, fmt.Errorf("%s must be a boolean", envDeleteToken)
	}

	return parsedDelete, nil
}

func diagnosticResponse(summary string, err string) *tfprotov5.ConfigureProviderResponse {
	return &tfprotov5.ConfigureProviderResponse{
		Diagnostics: []*tfprotov5.Diagnostic{{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  summary,
			Detail:   err,
		}},
	}
}

func hasErrors(resp *tfprotov5.ConfigureProviderResponse) bool {
	if resp == nil {
		return false
	}

	for _, diag := range resp.Diagnostics {
		if diag != nil && diag.Severity == tfprotov5.DiagnosticSeverityError {
			return true
		}
	}

	return false
}

func restoreEnv(key, value string, existed bool) {
	if existed {
		_ = os.Setenv(key, value)
		return
	}

	_ = os.Unsetenv(key)
}
