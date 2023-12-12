package quote

import (
	"math/big"
	"time"
)

const staleThreshold = 86400

// DLLNode represents a node in the doubly linked list.
type DLLNode struct {
	transactionID string
	Value         *big.Int // exported for tests
	timestamp     int64
	prev          *DLLNode
	next          *DLLNode
}

// DLL represents a doubly linked list.
type DLL struct {
	head *DLLNode
	tail *DLLNode
}

// NewNode creates a new Node.
func NewNode(transactionID string, value *big.Int) *DLLNode {
	return &DLLNode{
		transactionID: transactionID,
		Value:         value,
		timestamp:     time.Now().Unix(),
		prev:          nil,
		next:          nil,
	}
}

// Head returns the head of the doubly linked list.
func (dll *DLL) Head() *DLLNode {
	return dll.head
}

// Tail returns the tail of the doubly linked list.
func (dll *DLL) Tail() *DLLNode {
	return dll.tail
}

// NewTimestamp is just for testing purposes.
func (dll *DLL) NewTimestamp(newTime int64, node *DLLNode) {
	node.timestamp = newTime
}

// AddNode adds a node to the front of the list.
func (dll *DLL) AddNode(node *DLLNode) {
	if dll.head == nil {
		dll.head = node
		dll.tail = node
		return
	}
	dll.head.prev = node
	node.next = dll.head
	dll.head = node
}

// RemoveNode removes a given node from the list.
func (dll *DLL) RemoveNode(node *DLLNode) {
	if node == dll.head {
		dll.head = node.next
	}
	if node == dll.tail {
		dll.tail = node.prev
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
}

// ClearStale clears any stale items from the list.
func (dll *DLL) ClearStale() {
	for dll.tail != nil && dll.tail.timestamp+staleThreshold < time.Now().Unix() {
		dll.RemoveNode(dll.tail)
	}
}

// IBridgeReqHandler is the interface for the handling bridge requests.
type IBridgeReqHandler interface {
	GetAndDelete(transactionID string) (*big.Int, bool) // Pop
	Put(transactionID string, value *big.Int)
	Capacity() int
	Sum() *big.Int
	GetHashmapLen() int
	GetHead() *DLLNode
}

// bridgeReqHandlerImpl represents a doubly linked list of unconfirmed bridge requests.
type bridgeReqHandlerImpl struct {
	capacity int
	hashmap  map[string]*DLLNode
	ll       DLL
	sum      *big.Int
}

// NewBridgeReqs creates a new IBridgeReqHandler.
func NewBridgeReqs(capacity int) IBridgeReqHandler {
	return &bridgeReqHandlerImpl{
		capacity: capacity,
		hashmap:  make(map[string]*DLLNode),
		sum:      big.NewInt(0),
	}
}

// GetAndDelete gets the value of the key if the transactionID exists.
func (b *bridgeReqHandlerImpl) GetAndDelete(transactionID string) (*big.Int, bool) {
	if node, ok := b.hashmap[transactionID]; ok {
		b.ll.RemoveNode(node)
		b.sum.Sub(b.sum, node.Value)
		delete(b.hashmap, node.transactionID)
		return node.Value, true
	}
	return nil, false
}

// Put sets or inserts the value if the key is not already present.
func (b *bridgeReqHandlerImpl) Put(transactionID string, value *big.Int) {
	if node, ok := b.hashmap[transactionID]; ok {
		b.ll.RemoveNode(node)
		b.sum.Sub(b.sum, node.Value)
		b.ll.AddNode(node)
		b.sum.Add(b.sum, value)
		node.Value = value
		return
	}

	if len(b.hashmap) == b.capacity {
		tail := b.ll.tail
		b.ll.RemoveNode(tail)
		b.sum.Sub(b.sum, tail.Value)
		delete(b.hashmap, tail.transactionID)
	}

	newNode := NewNode(transactionID, value)
	b.ll.AddNode(newNode)
	b.sum.Add(b.sum, value)
	b.hashmap[transactionID] = newNode
}

// Sum returns the sum of the values in the hashmap.
func (b *bridgeReqHandlerImpl) Sum() *big.Int {
	return b.sum
}

// Capacity returns the capacity of the hashmap.
func (b *bridgeReqHandlerImpl) Capacity() int {
	return b.capacity
}

// GetHashmapLen returns the length of the hashmap.
func (b *bridgeReqHandlerImpl) GetHashmapLen() int {
	return len(b.hashmap)
}

// GetHead returns the head of the doubly linked list.
func (b *bridgeReqHandlerImpl) GetHead() *DLLNode {
	return b.ll.Head()
}

var _ IBridgeReqHandler = &bridgeReqHandlerImpl{}
