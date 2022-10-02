package collection

import (
	"bytes"
	// used to embed the collection.
	_ "embed"
	"fmt"
	postman "github.com/rbretecher/go-postman-collection"
)

//go:embed collection.json
var ethCollection []byte

// CreateCollection creates a new postman collection.
func CreateCollection() ([]byte, error) {
	collection, err := postman.ParseCollection(bytes.NewReader(ethCollection))
	if err != nil {
		return nil, fmt.Errorf("could not parse collection: %w", err)
	}

	// there's currently a bug where we can't deepcopy items and have them marshaled so we do this twice
	confCollection, err := postman.ParseCollection(bytes.NewReader(ethCollection))
	if err != nil {
		return nil, fmt.Errorf("could not parse collection: %w", err)
	}

	collection.Variables = []*postman.Variable{
		{
			ID:          "ENVIRONMENT",
			Name:        "Base url",
			Description: "Base url of the omnirpc server",
		},
		{
			ID:          "CHAIN_ID",
			Name:        "Chain ID",
			Description: "Chain ID to send requests to",
		},
		{
			ID:          "CONFIRMATIONS",
			Name:        "Confirmations",
			Description: "Confirmation count",
		},
	}

	var omniNamespace, confNamespace *postman.Items

	// copy the eth unsupported namespaces for omni
	omniNamespace = getEthNamespace(collection)
	confNamespace = getEthNamespace(confCollection)

	for i, item := range omniNamespace.Items {
		item.Request.URL.Raw = fmt.Sprintf("{{%s}}/rpc/{{%s}}", collection.Variables[0].ID, collection.Variables[1].ID)
		item.Request.URL.Path = []string{
			"rpc",
			fmt.Sprintf("{{%s}}", collection.Variables[1].ID),
		}
		omniNamespace.Items[i] = item
	}

	for i, item := range confNamespace.Items {
		item.Request.URL.Raw = fmt.Sprintf("{{%s}}/confirmations/{{%s}}/rpc/{{%s}}", collection.Variables[0].ID, collection.Variables[2].ID, collection.Variables[1].ID)
		item.Request.URL.Path = []string{
			"confirmations",
			fmt.Sprintf("{{%s}}", collection.Variables[2].ID),
			"rpc",
			fmt.Sprintf("{{%s}}", collection.Variables[1].ID),
		}
		confNamespace.Items[i] = item
	}

	confNamespace.Name = "eth (confirmations)"

	collection.Items = []*postman.Items{omniNamespace, confNamespace}

	var b bytes.Buffer
	err = collection.Write(&b, postman.V210)
	if err != nil {
		return nil, fmt.Errorf("could not write collection: %w", err)
	}

	return b.Bytes(), nil
}

// getEthNamespace gets the eth namespace from a collection.
func getEthNamespace(collection *postman.Collection) *postman.Items {
	for _, coll := range collection.Items {
		if coll.Name == "eth" {
			return coll
		}
	}
	// TODO: error here, maybe?
	return nil
}
