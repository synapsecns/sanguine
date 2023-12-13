// Package queue implements a thread safe queue.
package queue

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db/model"
)

// IQueue is the interface for a thread safe queue.
type IQueue interface {
	Enqueue(ctx context.Context, transactionID string) error
	Dequeue(ctx context.Context) (string, error)
	HasLiveElements(uint64) (bool, error)
	Len() int
	Kill()
}

// Queue is a thread safe queue implementation.
type Queue struct {
	mu       sync.Mutex
	capacity int
	q        []*model.DeadlineQueue
	deadline int64
	db       db.DB
}

// NewQueue creates an empty queue with desired capacity.
func NewQueue(ctx context.Context, capacity int, deadline int64, db db.DB) (*Queue, error) {
	// Load any existing records from DB
	records, err := db.GetDeadlineQueueEvents(ctx)
	if err != nil {
		return nil, fmt.Errorf("error loading queue from db: %w", err)
	}

	return &Queue{
		capacity: capacity,
		q:        records,
		deadline: deadline,
		db:       db,
	}, nil
}

// Enqueue inserts the item into the queue and underlying queue table.
func (q *Queue) Enqueue(ctx context.Context, transactionID string) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	newEntry := &model.DeadlineQueue{
		Timestamp:     time.Now().Unix(),
		TransactionID: transactionID,
	}

	if len(q.q) < q.capacity {
		q.q = append(q.q, newEntry)
		// Insert into DB
		err := q.db.StoreDeadlineQueueEvent(ctx, newEntry)
		if err != nil {
			logger.Errorf("Error storing event into queue db: %v", err)
			return fmt.Errorf("error storing event into queue db: %w", err)
		}

		return nil
	}

	return fmt.Errorf("queue is full")
}

// Dequeue removes the oldest element from the queue.
func (q *Queue) Dequeue(ctx context.Context) (string, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.q) > 0 {
		item := q.q[0]
		q.q = q.q[1:]

		// Remove from DB, non-fatal error as removal will not create data inconsistency anywhere.
		err := q.db.RemoveDeadlineQueueEvent(ctx, item.TransactionID)
		if err != nil {
			logger.Errorf("Error removing event from queue db: %v", err)
		}

		return item.TransactionID, nil
	}

	return "", errQueueEmpty
}

// HasLiveElements checks if the queue has any elements that are past the deadline.
func (q *Queue) HasLiveElements(currTime int64) (bool, error) {
	if len(q.q) > 0 {
		if currTime >= q.q[0].Timestamp+q.deadline {
			return true, nil
		}
		return false, nil
	}
	return false, errQueueEmpty
}

var errQueueEmpty = errors.New("queue is empty")

// Len returns the length of the queue.
func (q *Queue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.q)
}

// Kill removes all elements from the queue, this is used for testing purposes.
func (q *Queue) Kill() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.q = []*model.DeadlineQueue{}
}
