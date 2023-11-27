package types

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/99designs/gqlgen/graphql"
)

// JSON is a generic JSON type.
type JSON map[string]interface{}

// MarshalJSON implements the graphql.Marshaler interface.
func MarshalJSON(b JSON) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		byteData, err := json.Marshal(b)
		if err != nil {
			log.Printf("FAIL WHILE MARSHAL JSON %v\n", string(byteData))
		}
		_, err = w.Write(byteData)
		if err != nil {
			log.Printf("FAIL WHILE WRITE DATA %v\n", string(byteData))
		}
	})
}

// UnmarshalJSON converts an object to a JSON type.
func UnmarshalJSON(v interface{}) (JSON, error) {
	byteData, err := json.Marshal(v)
	if err != nil {
		return JSON{}, fmt.Errorf("FAIL WHILE MARSHAL SCHEME")
	}
	tmp := make(map[string]interface{})
	err = json.Unmarshal(byteData, &tmp)
	if err != nil {
		return JSON{}, fmt.Errorf("FAIL WHILE UNMARSHAL SCHEME")
	}
	return tmp, nil
}
