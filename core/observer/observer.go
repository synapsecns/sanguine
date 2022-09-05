package observer

import (
	"github.com/cheekybits/genny/generic"
	"sync"
)

// Key is the type used for ginny to generate generics.
type Key generic.Type

// KeyObserver implements an event manager that allows us to add/remove listeners at will
// listeners can be added on an event "e" and pass a channel in. Every time an event "e"
// is emitted, all listeners listening on channel "e" will receive that event.
// see https://flaviocopes.com/golang-event-listeners/ for details.
//
// Deprecated: this should be replaced by ethereums's event subscription: https://pkg.go.dev/github.com/ethereum/go-ethereum@v1.10.5/event#Subscription
type KeyObserver struct {
	listeners map[Key][]chan interface{}
	mux       sync.RWMutex
}

// NewKeyObserver creates a new observer.
func NewKeyObserver() *KeyObserver {
	return &KeyObserver{
		listeners: nil,
		mux:       sync.RWMutex{},
	}
}

// AddListener adds an event listener to the KeyObserver struct instance. This listener
// will receive all events on event "e" on channel "ch".
func (b *KeyObserver) AddListener(e Key, ch chan interface{}) {
	b.mux.Lock()
	defer b.mux.Unlock()
	if b.listeners == nil {
		b.listeners = make(map[Key][]chan interface{})
	}

	b.listeners[e] = append(b.listeners[e], ch)
}

// RemoveListener removes an event listener from the KeyObserver struct instance.
// Note: listeners must be removed before the channel is closed to prevent a race condition.
func (b *KeyObserver) RemoveListener(e Key, ch chan interface{}) {
	b.mux.Lock()
	defer b.mux.Unlock()
	if _, ok := b.listeners[e]; ok {
		for i := range b.listeners[e] {
			if b.listeners[e][i] == ch {
				b.listeners[e] = append(b.listeners[e][:i], b.listeners[e][i+1:]...)
				break
			}
		}
	}
}

// Emit emits an event on the KeyObserver struct instance. The event is received by all listeners
// listening on event.
func (b *KeyObserver) Emit(e Key, response interface{}) {
	b.mux.RLock()
	defer b.mux.RUnlock()
	if _, ok := b.listeners[e]; ok {
		for _, handler := range b.listeners[e] {
			go func(handler chan interface{}) {
				handler <- response
			}(handler)
		}
	}
}
