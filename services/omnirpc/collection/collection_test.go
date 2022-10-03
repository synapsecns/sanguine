package collection_test

import (
	"bytes"
	"github.com/rbretecher/go-postman-collection"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/omnirpc/collection"
)

func (c *CollectionSuite) TestCollection() {
	res, err := collection.CreateCollection()
	Nil(c.T(), err)

	coll, err := postman.ParseCollection(bytes.NewReader(res))
	Nil(c.T(), err)
	// 2 collections
	Equal(c.T(), len(coll.Items), 2)
}
