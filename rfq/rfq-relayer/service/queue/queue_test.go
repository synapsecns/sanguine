// nolint: goconst
package queue_test

import (
	"time"

	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db/model"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/service/queue"
)

func (t *QueueSuite) TestEnqueue() {
	t.RunOnAllDBs(func(testDB db.TestDB) {
		ctx := t.GetTestContext()
		q, err := queue.NewQueue(ctx, 2, 100, testDB)
		Nil(t.T(), err)
		transactionID := "request1"
		transactionID2 := "request2"
		// Test enqueueing to a non-full queue
		err = q.Enqueue(ctx, transactionID)
		Nil(t.T(), err)
		Equal(t.T(), 1, q.Len())

		events := []model.DeadlineQueue{}
		err = testDB.UNSAFE_DB().WithContext(ctx).Model(&model.DeadlineQueue{}).Where(&model.DeadlineQueue{}).Scan(&events).Error
		Nil(t.T(), err)
		Equal(t.T(), 1, len(events)) // Make sure there's something in the database

		// Relay queue
		err = q.Enqueue(ctx, transactionID2)
		Nil(t.T(), err)
		Equal(t.T(), 2, q.Len())

		twoEvents := []model.DeadlineQueue{}
		err = testDB.UNSAFE_DB().WithContext(ctx).Model(&model.DeadlineQueue{}).Where(&model.DeadlineQueue{}).Scan(&twoEvents).Error
		Nil(t.T(), err)
		Equal(t.T(), 2, len(twoEvents)) // Make sure there's 2 records in the database

		// Test enqueueing to a full queue
		err = q.Enqueue(ctx, transactionID2)
		// Expect error as length of queue is 2
		NotNil(t.T(), err)
		Equal(t.T(), 2, q.Len())

		twoEventsSecondCheck := []model.DeadlineQueue{}
		err = testDB.UNSAFE_DB().WithContext(ctx).Model(&model.DeadlineQueue{}).Where(&model.DeadlineQueue{}).Scan(&twoEventsSecondCheck).Error
		Nil(t.T(), err)
		Equal(t.T(), 2, len(twoEventsSecondCheck)) // Make sure there's still 2 records in the database
	})
}

func (t *QueueSuite) TestDequeue() {
	t.RunOnAllDBs(func(testDB db.TestDB) {
		ctx := t.GetTestContext()
		q, err := queue.NewQueue(ctx, 2, 100, testDB)
		Nil(t.T(), err)
		transactionID := "request1"

		err = q.Enqueue(ctx, transactionID)
		Nil(t.T(), err)

		// Test dequeueing from a non-empty queue
		dequeuedNode, err := q.Dequeue(ctx)
		Nil(t.T(), err)
		Equal(t.T(), 0, q.Len())
		Equal(t.T(), transactionID, dequeuedNode)

		// Test dequeueing from an empty queue
		_, err = q.Dequeue(ctx)
		NotNil(t.T(), err)
	})
}

func (t *QueueSuite) TestHasLiveElements() {
	t.RunOnAllDBs(func(testDB db.TestDB) {
		ctx := t.GetTestContext()
		q, err := queue.NewQueue(ctx, 2, 5, testDB)
		Nil(t.T(), err)

		transactionID := "request1"
		// Test peeking at an empty queue
		_, err = q.HasLiveElements(200)
		NotNil(t.T(), err)

		// Enqueue a node and test peeking
		err = q.Enqueue(ctx, transactionID)
		Nil(t.T(), err)
		found, err := q.HasLiveElements(time.Now().Unix()) // Current time is before the deadline
		Nil(t.T(), err)
		False(t.T(), found)

		// Wait for deadline to pass
		<-time.After(5 * time.Second)
		found, err = q.HasLiveElements(time.Now().Unix()) // Current time is after the deadline
		Nil(t.T(), err)
		True(t.T(), found)
	})
}

// This tests the queue's ability to load from the database (the relayer should be restart safe).
// This is important for any case in which the relayer goes down or restarts. The queue should be able to load from the
// database and continue where it left off without any stuck claims.
func (t *QueueSuite) TestLoadFromDB() {
	t.RunOnAllDBs(func(testDB db.TestDB) {
		ctx := t.GetTestContext()
		entry1 := &model.DeadlineQueue{
			Timestamp:     time.Now().Unix(),
			TransactionID: "request1",
		}
		entry2 := &model.DeadlineQueue{
			Timestamp:     time.Now().Unix(),
			TransactionID: "request2",
		}

		// Test cold start (nothing in database)
		q1, err := queue.NewQueue(ctx, 2, 5, testDB)
		Nil(t.T(), err)
		Equal(t.T(), 0, q1.Len())
		err = q1.Enqueue(ctx, entry1.TransactionID)
		Nil(t.T(), err)
		err = q1.Enqueue(ctx, entry2.TransactionID)
		Nil(t.T(), err)
		q1.Kill() // Simulated shutdown

		// Test warm start/post restart (things in the database)
		// New queue should load from database
		q2, err := queue.NewQueue(ctx, 2, 5, testDB)
		Nil(t.T(), err)
		Equal(t.T(), 2, q2.Len())

		// Test dequeueing from a non-empty queue
		dequeuedNode, err := q2.Dequeue(ctx)
		Nil(t.T(), err)
		Equal(t.T(), entry1.TransactionID, dequeuedNode)
		Equal(t.T(), 1, q2.Len())

		// Check that the database is updated
		remainingEvents := []model.DeadlineQueue{}
		err = testDB.UNSAFE_DB().WithContext(ctx).Model(&model.DeadlineQueue{}).Where(&model.DeadlineQueue{}).Scan(&remainingEvents).Error
		Nil(t.T(), err)
		Equal(t.T(), 1, len(remainingEvents)) // Make sure there's 1 record in the database
	})
}
