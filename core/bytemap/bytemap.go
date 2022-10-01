package bytemap

// ByteSliceMap emulates `map[[]byte]V`, implemented as a Trie.
//
// It seems to perform worse than `map[string]V` even when casting `string([]byte)`.
type ByteSliceMap[V any] struct {
	value    V
	children map[byte]*ByteSliceMap[V]
}

// PutString is a convenience method to insert a value using a string key.
func (n *ByteSliceMap[V]) PutString(s string, v V) *ByteSliceMap[V] {
	return n.Put([]byte(s), v)
}

// Put inserts a value into the `ByteMap` using `[]byte` as a key.
func (n *ByteSliceMap[V]) Put(s []byte, v V) *ByteSliceMap[V] {
	for _, r := range s {
		n = n.put(r)
	}
	n.value = v
	return n
}

func (n *ByteSliceMap[V]) put(r byte) *ByteSliceMap[V] {
	if child, ok := n.children[r]; ok {
		return child
	}
	var child ByteSliceMap[V]
	if n.children == nil {
		n.children = map[byte]*ByteSliceMap[V]{r: &child}
	} else {
		n.children[r] = &child
	}
	return &child
}

// Get returns a value as mapped by the `[]byte` key and a boolean of whether the value exists in the map.
func (n *ByteSliceMap[V]) Get(s []byte) (value V, _ bool) {
	for _, r := range s {
		var ok bool
		if n, ok = n.children[r]; !ok {
			return value, false
		}
	}
	return n.value, true
}

// GetString is a convenience method to get a value using a string key.
//
// See: `Get`.
func (n *ByteSliceMap[V]) GetString(s string) (V, bool) {
	return n.Get([]byte(s))
}
