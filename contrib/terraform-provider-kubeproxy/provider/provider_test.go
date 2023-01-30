package provider_test

import (
	"github.com/synapsecns/sanguine/contrib/terraform-provider-kubeproxy/provider"
	"testing"
)

func TestManifestProvider(t *testing.T) {
	fn, err := provider.ManifestProvider()
	if err != nil {
		t.Fatalf("error creating manifest prov: %v", err)
	}

	server := fn()

	prov, ok := server.(*provider.RawProviderServer)
	if !ok {
		t.Fatalf("incorrect prov type: %T", server)
	}

	if prov.CombinedSchema() == nil {
		t.Fatalf("combined schema should not be nil")
	}
	if prov.GoogleProvider() == nil {
		t.Fatalf("google prov should not be nil")
	}
	if prov.RawProviderServer == nil {
		t.Fatalf("raw prov server should not be nil")
	}
}
