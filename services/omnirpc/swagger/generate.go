// package main generates a static swagger collection for use in documentation. This differs from collection which dynamically generates a collection for serving.
// this is purely meant for static linking in the repo.
package main

import (
	"github.com/synapsecns/sanguine/services/omnirpc/collection"
	"os"
)

//go:generate go run github.com/synapsecns/sanguine/services/omnirpc/swagger

func main() {
	res, err := collection.CreateCollection()
	if err != nil {
		panic(err)
	}

	file, err := os.Create("collection.json")
	if err != nil {
		panic(err)
	}

	_, err = file.Write(res)
	if err != nil {
		panic(err)
	}
}
