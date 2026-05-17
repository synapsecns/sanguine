package provider

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/grafana"
	"github.com/aws/aws-sdk-go-v2/service/grafana/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

type fakeProviderServer struct {
	tfprotov5.ProviderServer

	configureCalls int
	stopCalls      int
	configureResp  *tfprotov5.ConfigureProviderResponse
	configureErr   error
	stopResp       *tfprotov5.StopProviderResponse
	stopErr        error
}

func (f *fakeProviderServer) ConfigureProvider(context.Context, *tfprotov5.ConfigureProviderRequest) (*tfprotov5.ConfigureProviderResponse, error) {
	f.configureCalls++

	return f.configureResp, f.configureErr
}

func (f *fakeProviderServer) StopProvider(context.Context, *tfprotov5.StopProviderRequest) (*tfprotov5.StopProviderResponse, error) {
	f.stopCalls++

	return f.stopResp, f.stopErr
}

type fakeGrafanaClient struct {
	createInput *grafana.CreateWorkspaceServiceAccountTokenInput
	createOut   *grafana.CreateWorkspaceServiceAccountTokenOutput
	createErr   error
	deleteInput *grafana.DeleteWorkspaceServiceAccountTokenInput
	deleteErr   error
}

func (f *fakeGrafanaClient) CreateWorkspaceServiceAccountToken(_ context.Context, input *grafana.CreateWorkspaceServiceAccountTokenInput, _ ...func(*grafana.Options)) (*grafana.CreateWorkspaceServiceAccountTokenOutput, error) {
	f.createInput = input

	return f.createOut, f.createErr
}

func (f *fakeGrafanaClient) DeleteWorkspaceServiceAccountToken(_ context.Context, input *grafana.DeleteWorkspaceServiceAccountTokenInput, _ ...func(*grafana.Options)) (*grafana.DeleteWorkspaceServiceAccountTokenOutput, error) {
	f.deleteInput = input

	return &grafana.DeleteWorkspaceServiceAccountTokenOutput{}, f.deleteErr
}

func withGrafanaClient(t *testing.T, client grafanaAPI) {
	t.Helper()

	previousClient := newGrafanaClient
	newGrafanaClient = func(context.Context, string) (grafanaAPI, error) {
		return client, nil
	}

	t.Cleanup(func() {
		newGrafanaClient = previousClient
	})
}

func TestConfigureProviderDelegatesWhenAMGDisabled(t *testing.T) {
	t.Setenv(envWorkspaceID, "")
	t.Setenv(envServiceAccountID, "")
	t.Setenv(envGrafanaAuth, "")

	upstreamResp := &tfprotov5.ConfigureProviderResponse{}
	upstream := &fakeProviderServer{configureResp: upstreamResp}
	server := New(upstream)

	resp, err := server.ConfigureProvider(context.Background(), &tfprotov5.ConfigureProviderRequest{})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp != upstreamResp {
		t.Fatal("expected upstream response")
	}
	if upstream.configureCalls != 1 {
		t.Fatalf("expected one configure call, got %d", upstream.configureCalls)
	}
}

func TestConfigureProviderCreatesTokenAndDelegates(t *testing.T) {
	t.Setenv(envWorkspaceID, "g-123")
	t.Setenv(envServiceAccountID, "sa-123")
	t.Setenv(envGrafanaAuth, "")

	client := &fakeGrafanaClient{
		createOut: &grafana.CreateWorkspaceServiceAccountTokenOutput{
			ServiceAccountToken: &types.ServiceAccountTokenSummaryWithKey{
				Id:  aws.String("token-123"),
				Key: aws.String("secret"),
			},
		},
	}
	withGrafanaClient(t, client)

	upstream := &fakeProviderServer{configureResp: &tfprotov5.ConfigureProviderResponse{}}
	server := New(upstream)

	resp, err := server.ConfigureProvider(context.Background(), &tfprotov5.ConfigureProviderRequest{})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if hasErrors(resp) {
		t.Fatal("expected no diagnostics")
	}
	if upstream.configureCalls != 1 {
		t.Fatalf("expected one configure call, got %d", upstream.configureCalls)
	}
	if aws.ToString(client.createInput.WorkspaceId) != "g-123" {
		t.Fatalf("unexpected workspace ID: %s", aws.ToString(client.createInput.WorkspaceId))
	}
	if os.Getenv(envGrafanaAuth) != "" {
		t.Fatal("expected GRAFANA_AUTH to be restored after configure")
	}
}

func TestConfigureProviderDelegatesWhenGrafanaAuthSet(t *testing.T) {
	t.Setenv(envWorkspaceID, "g-123")
	t.Setenv(envServiceAccountID, "sa-123")
	t.Setenv(envGrafanaAuth, "already-configured")

	upstreamResp := &tfprotov5.ConfigureProviderResponse{}
	upstream := &fakeProviderServer{configureResp: upstreamResp}
	server := New(upstream)

	resp, err := server.ConfigureProvider(context.Background(), &tfprotov5.ConfigureProviderRequest{})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp != upstreamResp {
		t.Fatal("expected upstream response")
	}
	if upstream.configureCalls != 1 {
		t.Fatalf("expected one configure call, got %d", upstream.configureCalls)
	}
}

func TestConfigureProviderReportsTokenCreationError(t *testing.T) {
	t.Setenv(envWorkspaceID, "g-123")
	t.Setenv(envServiceAccountID, "sa-123")
	t.Setenv(envGrafanaAuth, "")

	withGrafanaClient(t, &fakeGrafanaClient{createErr: errors.New("boom")})

	resp, err := New(&fakeProviderServer{}).ConfigureProvider(context.Background(), &tfprotov5.ConfigureProviderRequest{})
	if err != nil {
		t.Fatalf("expected diagnostic response, got error %v", err)
	}
	if !hasErrors(resp) {
		t.Fatal("expected diagnostic error")
	}
}

func TestStopProviderDelegates(t *testing.T) {
	upstreamResp := &tfprotov5.StopProviderResponse{}
	upstream := &fakeProviderServer{stopResp: upstreamResp}
	server := New(upstream)

	resp, err := server.StopProvider(context.Background(), &tfprotov5.StopProviderRequest{})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp != upstreamResp {
		t.Fatal("expected upstream response")
	}
	if upstream.stopCalls != 1 {
		t.Fatalf("expected one stop call, got %d", upstream.stopCalls)
	}
}

func TestDeleteToken(t *testing.T) {
	client := &fakeGrafanaClient{}
	withGrafanaClient(t, client)

	err := deleteToken(context.Background(), createdToken{
		workspaceID:      "g-123",
		serviceAccountID: "sa-123",
		tokenID:          "token-123",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if aws.ToString(client.deleteInput.TokenId) != "token-123" {
		t.Fatalf("unexpected token ID: %s", aws.ToString(client.deleteInput.TokenId))
	}
}

func TestDeleteTokenWrapsClientError(t *testing.T) {
	withGrafanaClient(t, &fakeGrafanaClient{deleteErr: errors.New("boom")})

	err := deleteToken(context.Background(), createdToken{
		workspaceID:      "g-123",
		serviceAccountID: "sa-123",
		tokenID:          "token-123",
	})
	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestLoadAMGConfigDisabled(t *testing.T) {
	t.Setenv(envWorkspaceID, "")
	t.Setenv(envServiceAccountID, "")

	_, enabled, err := loadAMGConfig()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if enabled {
		t.Fatal("expected AMG config to be disabled")
	}
}

func TestLoadAMGConfigRequiresWorkspaceAndServiceAccount(t *testing.T) {
	t.Setenv(envWorkspaceID, "g-123")
	t.Setenv(envServiceAccountID, "")

	_, enabled, err := loadAMGConfig()
	if err == nil {
		t.Fatal("expected an error")
	}
	if enabled {
		t.Fatal("expected AMG config to be disabled on invalid input")
	}
}

func TestLoadAMGConfigDefaults(t *testing.T) {
	t.Setenv(envWorkspaceID, "g-123")
	t.Setenv(envServiceAccountID, "sa-123")
	t.Setenv(envRegion, "")
	t.Setenv(envTokenTTLSeconds, "")
	t.Setenv(envTokenNamePrefix, "")
	t.Setenv(envDeleteToken, "")

	cfg, enabled, err := loadAMGConfig()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !enabled {
		t.Fatal("expected AMG config to be enabled")
	}
	if cfg.workspaceID != "g-123" {
		t.Fatalf("unexpected workspace ID: %s", cfg.workspaceID)
	}
	if cfg.serviceAccountID != "sa-123" {
		t.Fatalf("unexpected service account ID: %s", cfg.serviceAccountID)
	}
	if cfg.tokenTTLSeconds != defaultTokenTTLSeconds {
		t.Fatalf("unexpected token TTL: %d", cfg.tokenTTLSeconds)
	}
	if cfg.tokenNamePrefix != defaultTokenNamePrefix {
		t.Fatalf("unexpected token name prefix: %s", cfg.tokenNamePrefix)
	}
	if !cfg.deleteToken {
		t.Fatal("expected token cleanup to default to true")
	}
}

func TestLoadAMGConfigOverrides(t *testing.T) {
	t.Setenv(envWorkspaceID, "g-123")
	t.Setenv(envServiceAccountID, "sa-123")
	t.Setenv(envRegion, "ap-northeast-1")
	t.Setenv(envTokenTTLSeconds, "120")
	t.Setenv(envTokenNamePrefix, "ci")
	t.Setenv(envDeleteToken, "false")

	cfg, enabled, err := loadAMGConfig()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !enabled {
		t.Fatal("expected AMG config to be enabled")
	}
	if cfg.region != "ap-northeast-1" {
		t.Fatalf("unexpected region: %s", cfg.region)
	}
	if cfg.tokenTTLSeconds != 120 {
		t.Fatalf("unexpected token TTL: %d", cfg.tokenTTLSeconds)
	}
	if cfg.tokenNamePrefix != "ci" {
		t.Fatalf("unexpected token name prefix: %s", cfg.tokenNamePrefix)
	}
	if cfg.deleteToken {
		t.Fatal("expected token cleanup override to be false")
	}
}

func TestLoadAMGConfigRejectsInvalidTTL(t *testing.T) {
	t.Setenv(envWorkspaceID, "g-123")
	t.Setenv(envServiceAccountID, "sa-123")
	t.Setenv(envTokenTTLSeconds, "0")

	_, enabled, err := loadAMGConfig()
	if err == nil {
		t.Fatal("expected an error")
	}
	if enabled {
		t.Fatal("expected AMG config to be disabled on invalid input")
	}
}

func TestLoadAMGConfigRejectsInvalidDeleteToken(t *testing.T) {
	t.Setenv(envWorkspaceID, "g-123")
	t.Setenv(envServiceAccountID, "sa-123")
	t.Setenv(envDeleteToken, "not-bool")

	_, enabled, err := loadAMGConfig()
	if err == nil {
		t.Fatal("expected an error")
	}
	if enabled {
		t.Fatal("expected AMG config to be disabled on invalid input")
	}
}

func TestDiagnosticResponse(t *testing.T) {
	resp := diagnosticResponse("summary", "detail")
	if len(resp.Diagnostics) != 1 {
		t.Fatalf("expected one diagnostic, got %d", len(resp.Diagnostics))
	}
	if resp.Diagnostics[0].Severity != tfprotov5.DiagnosticSeverityError {
		t.Fatalf("unexpected diagnostic severity: %v", resp.Diagnostics[0].Severity)
	}
	if resp.Diagnostics[0].Summary != "summary" {
		t.Fatalf("unexpected diagnostic summary: %s", resp.Diagnostics[0].Summary)
	}
}

func TestHasErrors(t *testing.T) {
	if hasErrors(nil) {
		t.Fatal("nil response should not have errors")
	}
	if hasErrors(&tfprotov5.ConfigureProviderResponse{}) {
		t.Fatal("empty response should not have errors")
	}
	if !hasErrors(&tfprotov5.ConfigureProviderResponse{
		Diagnostics: []*tfprotov5.Diagnostic{{
			Severity: tfprotov5.DiagnosticSeverityError,
		}},
	}) {
		t.Fatal("expected error diagnostic to be detected")
	}
}

func TestRestoreEnv(t *testing.T) {
	const key = "GRAFANAAMG_TEST_RESTORE"

	t.Setenv(key, "new")

	restoreEnv(key, "old", true)
	if got := os.Getenv(key); got != "old" {
		t.Fatalf("expected env to be restored, got %q", got)
	}

	t.Setenv(key, "new")

	restoreEnv(key, "", false)
	if _, ok := os.LookupEnv(key); ok {
		t.Fatal("expected env to be unset")
	}
}
