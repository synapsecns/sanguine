package observer_test

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/observer"
	"golang.org/x/sync/semaphore"
	"sync"
	"time"
)

// TestObserverRemoveNonExistentListener makes sure observer doesn't break if we try to remove a non-existent listener.
func (s *ObserverSuite) TestObserverRemoveNonExistentListener() {
	testKeyObserver := observer.NewKeyObserver()

	unusedChan := make(chan interface{})
	testKeyObserver.RemoveListener(gofakeit.Word(), unusedChan)
}

// TestObserverEmitNoChannels makes sure the observer doesn't break if we try to emit to no listeners.
func (s *ObserverSuite) TestObserverEmitNoChannels() {
	testKeyObserver := observer.NewKeyObserver()

	testKeyObserver.Emit(gofakeit.Word(), gofakeit.Person())
}

// TestManyListeners tests emission over many listeners.
func (s *ObserverSuite) TestManyListeners() {
	// listenerCount is how many listeners to create
	const listenerCount = 100
	// removeAfterMessages removes the listeners after this many messages
	const removeAfterMessages = 100

	// timeout the test if it doesn't pass after 5 seconds
	ctx, cancel := context.WithTimeout(s.GetTestContext(), time.Second*5)
	defer cancel()

	// create a new waitgroup used for managing the listeners
	var wg sync.WaitGroup

	// create a new testKeyObserver
	testKeyObserver := observer.NewKeyObserver()
	// create a waitgroup of size listeners
	wg.Add(listenerCount)
	// create a semaphore. This will be acquired on listener creation and released
	// once the listener has been added to prevent race conditions in the test
	initializationLock := semaphore.NewWeighted(listenerCount)

	listenerName := gofakeit.Word()
	// create the listeners
	for i := 0; i < listenerCount; i++ {
		assert.Nil(s.T(), initializationLock.Acquire(ctx, 1))
		go func() {
			listener := make(chan interface{})
			defer close(listener)
			defer wg.Done()

			testKeyObserver.AddListener(listenerName, listener)
			initializationLock.Release(1)
			msgCount := 0
			for range listener {
				msgCount++
				if msgCount == removeAfterMessages {
					testKeyObserver.RemoveListener(listenerName, listener)
					return
				}
			}
		}()
	}

	assert.Nil(s.T(), initializationLock.Acquire(ctx, listenerCount))
	for i := 0; i < removeAfterMessages; i++ {
		testKeyObserver.Emit(listenerName, gofakeit.Word())
	}

	wg.Wait()
	// make sure the test wasn't canceled
	assert.Nil(s.T(), ctx.Err())
}
