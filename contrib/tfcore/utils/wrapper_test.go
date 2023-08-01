package utils_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/synapsecns/sanguine/contrib/tfcore/utils"
	"testing"
)

// nolint:staticcheck
func TestWrapSchemaResource(t *testing.T) {
	resource := &schema.Resource{
		Create: func(data *schema.ResourceData, meta interface{}) error {
			return nil
		},
		Read: func(data *schema.ResourceData, meta interface{}) error {
			return nil
		},
		Update: func(data *schema.ResourceData, meta interface{}) error {
			return nil
		},
	}

	// TODO: look into testing all fields with reflection
	wrappedResource := utils.WrapSchemaResource(resource)
	underlyingProvider := &MockWrappedProvider{}

	// Test Create hook
	err := wrappedResource.Create(nil, underlyingProvider)
	if err != nil {
		t.Fatalf("Expected Create to succeed, got %s", err)
	}

	// Test Read hook
	err = wrappedResource.Read(nil, underlyingProvider)
	if err != nil {
		t.Fatalf("Expected Read to succeed, got %s", err)
	}

	// Test Update hook
	err = wrappedResource.Update(nil, underlyingProvider)
	if err != nil {
		t.Fatalf("Expected Update to succeed, got %s", err)
	}
}

type MockWrappedProvider struct{}

func (m *MockWrappedProvider) UnderlyingProvider() interface{} {
	return nil
}

func (m *MockWrappedProvider) GoogleProvider() interface{} {
	return nil
}
