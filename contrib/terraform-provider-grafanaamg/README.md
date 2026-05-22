# Terraform Provider Grafana AMG

[![Go Reference](https://pkg.go.dev/badge/github.com/synapsecns/sanguine/contrib/terraform-provider-grafanaamg.svg)](https://pkg.go.dev/github.com/synapsecns/sanguine/contrib/terraform-provider-grafanaamg)
[![Go Report Card](https://goreportcard.com/badge/github.com/synapsecns/sanguine/contrib/terraform-provider-grafanaamg)](https://goreportcard.com/report/github.com/synapsecns/sanguine/contrib/terraform-provider-grafanaamg)

`terraform-provider-grafanaamg` wraps the upstream Grafana Terraform provider and configures it for Amazon Managed Grafana by creating a short-lived AMG service account token during provider configuration.

This keeps the Grafana API token out of Terraform state and avoids a separate Terraform workspace just to mint provider credentials. All Grafana resources and data sources are delegated to the upstream provider.

## Usage

Configure Terraform to use this provider binary for the `grafana` provider name:

```hcl
terraform {
  required_providers {
    grafana = {
      source  = "synapsecns/grafanaamg"
      version = ">= 0.1.0"
    }
  }
}

provider "grafana" {
  url = "https://g-0000000000.grafana-workspace.us-east-1.amazonaws.com"
}
```

Set the AMG workspace and service account identifiers through environment variables before running Terraform:

```bash
export GRAFANA_AMG_WORKSPACE_ID="g-0000000000"
export GRAFANA_AMG_SERVICE_ACCOUNT_ID="sa-000000000000000000000"
export GRAFANA_AMG_REGION="us-east-1"
```

The provider uses the default AWS SDK credential chain. The AWS principal must be allowed to call `grafana:CreateWorkspaceServiceAccountToken` and `grafana:DeleteWorkspaceServiceAccountToken` for the configured workspace and service account.

If `GRAFANA_AUTH` is already set, the wrapper does not create an AMG token and delegates directly to the upstream Grafana provider.

## Configuration

| Environment variable | Required | Default | Description |
| --- | --- | --- | --- |
| `GRAFANA_AMG_WORKSPACE_ID` | Yes | | AMG workspace ID. |
| `GRAFANA_AMG_SERVICE_ACCOUNT_ID` | Yes | | AMG service account ID used to mint Terraform tokens. |
| `GRAFANA_AMG_REGION` | No | AWS SDK default | AWS region for AMG API calls. |
| `GRAFANA_AMG_TOKEN_TTL_SECONDS` | No | `3600` | Lifetime for each generated service account token. |
| `GRAFANA_AMG_TOKEN_NAME_PREFIX` | No | `terraform-provider-grafanaamg` | Prefix for generated token names. |
| `GRAFANA_AMG_DELETE_TOKEN_ON_STOP` | No | `true` | Delete generated tokens when Terraform stops the provider. Tokens still expire by TTL if cleanup is interrupted. |

## Release

This module uses the Sanguine Terraform provider release shape:

- GoReleaser builds Terraform-style binaries named `terraform-provider-grafanaamg_v{{ .Version }}`.
- Release archives are zipped per OS/architecture.
- `terraform-registry-manifest.json` is included in release and checksum artifacts.
- The monorepo tag prefix is `contrib/terraform-provider-grafanaamg/`.

## Limitations

The AMG workspace and service account must already exist before this provider is configured. Terraform provider initialization happens before resources in the same graph can create those values.
