package quote_test

import (
	"math/big"

	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/service/quote"
)

// helper function to create a node with a specific transaction ID and value.
func createNode(transactionID string, value int64) *quote.DLLNode {
	return quote.NewNode(transactionID, big.NewInt(value))
}

func (t *QuoteSuite) TestAddNode() {
	dll := quote.DLL{}
	node1 := createNode("tx123", 100)
	node2 := createNode("tx456", 200)

	// Add first node
	dll.AddNode(node1)
	Equal(t.T(), node1, dll.Head())
	Equal(t.T(), node1, dll.Tail())

	// Add second node
	dll.AddNode(node2)
	Equal(t.T(), node2, dll.Head())
	Equal(t.T(), node1, dll.Tail())
}

func (t *QuoteSuite) TestRemoveNode() {
	dll := quote.DLL{}
	node1 := createNode("tx123", 100)
	node2 := createNode("tx456", 200)
	dll.AddNode(node1)
	dll.AddNode(node2)

	// Remove head
	dll.RemoveNode(node2)
	Equal(t.T(), node1, dll.Head())
	Equal(t.T(), node1, dll.Tail())

	// Remove last node
	dll.RemoveNode(node1)
	Equal(t.T(), quote.DLL{}, dll)
}

func (t *QuoteSuite) TestHeadInsertedMoreThanAnHourAgo() {
	dll := quote.DLL{}
	dll.ClearStale()
	Equal(t.T(), quote.DLL{}, dll)

	// Add a node and manipulate its timestamp to simulate an hour passing
	node := createNode("tx123", 100)
	dll.NewTimestamp(0, node)
	dll.AddNode(node)
	dll.ClearStale()
	Equal(t.T(), quote.DLL{}, dll)
}

func (t *QuoteSuite) TestNewBridgeReqs() {
	handler := quote.NewBridgeReqs(2)
	Equal(t.T(), 2, handler.Capacity())
	Equal(t.T(), 0, handler.GetHashmapLen())
}

func (t *QuoteSuite) TestBridgeReqHandlerPutAndGet() {
	handler := quote.NewBridgeReqs(2)

	// Test Put
	handler.Put("tx1", big.NewInt(100))
	Equal(t.T(), 1, handler.GetHashmapLen())
	Equal(t.T(), big.NewInt(100), handler.Sum())
	Equal(t.T(), big.NewInt(100), handler.GetHead().Value)

	// Test Get - existing item
	value, exists := handler.GetAndDelete("tx1")
	True(t.T(), exists)
	Equal(t.T(), big.NewInt(100), value)

	// Test capacity overflow
	handler.Put("tx1", big.NewInt(100))
	handler.Put("tx2", big.NewInt(200))
	handler.Put("tx3", big.NewInt(300))
	Equal(t.T(), 2, handler.GetHashmapLen())

	// Test Get - non-existing item
	_, exists = handler.GetAndDelete("tx1")
	False(t.T(), exists)
}

func (t *QuoteSuite) TestBridgeReqHandlerUpdateValue() {
	handler := quote.NewBridgeReqs(2)
	handler.Put("tx1", big.NewInt(100))
	handler.Put("tx1", big.NewInt(150))

	value, exists := handler.GetAndDelete("tx1")
	True(t.T(), exists)
	Equal(t.T(), big.NewInt(150), value)
	Equal(t.T(), 0, handler.GetHashmapLen())
}

func (t *QuoteSuite) TestBridgeReqHandlerCapacityHandling() {
	handler := quote.NewBridgeReqs(2)
	handler.Put("tx1", big.NewInt(100))
	handler.Put("tx2", big.NewInt(200))
	handler.Put("tx3", big.NewInt(300))

	Equal(t.T(), 2, handler.GetHashmapLen())
	Equal(t.T(), big.NewInt(500), handler.Sum())

	// Test gets
	_, exists := handler.GetAndDelete("tx1")
	False(t.T(), exists) // This guy was removed because of capacity overflow

	_, exists = handler.GetAndDelete("tx2")
	True(t.T(), exists)

	_, exists = handler.GetAndDelete("tx3")
	True(t.T(), exists)
}
