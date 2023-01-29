package manifest

import _ "golang.org/x/tools/benchmark/parse"

// required by go:generate
import _ "golang.org/x/mod/semver"

// required for copying the module
import _ "github.com/hashicorp/terraform-provider-kubernetes/manifest/provider"

// Note: we can't actually exclude this module from codeanalysis since we import it
//go:generate go run github.com/synapsecns/sanguine/tools/bundle -prefix ""  -pkg manifest -o manifest_gen.go github.com/hashicorp/terraform-provider-kubernetes/manifest/provider
