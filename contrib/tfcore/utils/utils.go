package utils

import (
	"github.com/hashicorp/terraform-provider-google/google/transport"
	"golang.org/x/oauth2"
	"reflect"
	"unsafe"
)

func GetTokenSource(config *transport.Config) oauth2.TokenSource {
	cfgReflection := reflect.ValueOf(config)
	accessToken := GetUnexportedField(cfgReflection.FieldByName("tokenSource")).(oauth2.TokenSource)
	return accessToken
}

func GetUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}

// MustCombineMaps attempts to combine two maps. Panics if maps can not be combined.
func MustCombineMaps[T interface{}](m1, m2 map[string]T) map[string]T {
	for key, value := range m2 {
		_, exists := m1[key]
		if exists {
			panic("Key overlap found when combining maps")
		}
		m1[key] = value
	}
	return m1
}
