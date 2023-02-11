package provider_test

import (
	"github.com/synapsecns/sanguine/contrib/terraform-provider-iap/provider"
	"testing"
)

// Make sure the provider loads.
func TestProviderLoad(t *testing.T) {
	prov := provider.Provider()
	if prov == nil {
		t.Fatal("Provider should not be nil")
	}
}
