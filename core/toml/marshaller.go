package toml

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/BurntSushi/toml"
)

// Encodable is an interface for a toml with an encode method.
type Encodable interface {
	Encode() (string, error)
}

// Indent is the indent we use.
const Indent = "  "

// Encode is a helper method to allow you to encode a toml to text.
// config should be passed here by pointer
// TODO: use a toml parser that can poparse comments.
func Encode(config interface{}) (string, error) {
	var buf bytes.Buffer
	encoder := toml.NewEncoder(&buf)
	encoder.Indent = Indent
	err := encoder.Encode(config)
	if err != nil {
		return "", fmt.Errorf("could not encode file: %w", err)
	}

	// currently, there's a bug in the parser that requires maps to be on the same level as the parent.
	// TODO: fix
	splitFile := strings.Split(buf.String(), "\n")
	var newLines []string
	for _, line := range splitFile {
		// get rid of double spacing on maps
		indentLen := len(Indent) * 2
		newLines = append(newLines, strings.ReplaceAll(line, getStringOfLength(indentLen), getStringOfLength(indentLen-2)))
	}

	return strings.Join(newLines, "\n"), nil
}

// getStringOfLength generates a blank string of length.
func getStringOfLength(length int) (res string) {
	for i := 0; i < length; i++ {
		res += " "
	}
	return res
}

// MarshalTextPtr should be a pointer here
// Deprecated: this can be done through the library.
func MarshalTextPtr(config interface{}) (text []byte, err error) {
	var fieldValDeref reflect.Value

	// make sure we have a pointer, if we don't return an error
	rv := eindirect(reflect.ValueOf(config))
	if rv.Kind() == reflect.Ptr {
		// if the pointer is nil, return nothing
		if rv.IsNil() {
			return text, nil
		}
		// otherwise dereference it for decoding
		rv.Set(reflect.New(rv.Type().Elem()))
		fieldValDeref = rv.Elem()
	} else {
		return text, fmt.Errorf("this method can only be run on pointers (have %s). If your method is not a pointer, you don't need a custom text marshaller", rv.Kind().String())
	}

	var buf bytes.Buffer
	encoder := toml.NewEncoder(&buf)
	err = encoder.Encode(fieldValDeref.Interface())
	if err != nil {
		return text, fmt.Errorf("could not encode file: %w", err)
	}
	return buf.Bytes(), nil
}

// eindirect wraps a pointer.
func eindirect(v reflect.Value) reflect.Value {
	//nolint:exhaustive
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return eindirect(v.Elem())
	default:
		return v
	}
}
