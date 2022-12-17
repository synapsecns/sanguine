// nolint
package newrelic

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/filipemendespi/newrelic-context/nrmock"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Model struct {
	ID    int
	Value string
}

var db *gorm.DB
var testTxn *newrelic.Transaction
var segmentsHistory []*nrmock.DatastoreSegment
var sampleLicense = "0123456789012345678901234567890123456789"

func TestMain(m *testing.M) {
	var err error
	// prepare db
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Microsecond, // Slow SQL threshold
			LogLevel:      logger.Info,      // Log level
			Colorful:      false,            // Disable color
		},
	)
	db, err = gorm.Open(sqlite.Open("./foo.db"), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	if err := db.Migrator().CreateTable(&Model{}); err != nil {
		panic(err)
	}
	if err := db.Create(&Model{Value: "to-select"}).Error; err != nil {
		panic(err)
	}
	AddGormCallbacks(db, nil)

	// mock newrelic
	originalBuilder := segmentBuilder
	segmentBuilder = func(
		startTime newrelic.SegmentStartTime,
		product newrelic.DatastoreProduct,
		query string,
		operation string,
		collection string,
	) segment {
		segment := originalBuilder(startTime, product, query, operation, collection).(*newrelic.DatastoreSegment)
		mock := &nrmock.DatastoreSegment{DatastoreSegment: segment, StartTime: startTime}
		segmentsHistory = append(segmentsHistory, mock)
		return mock
	}

	app, _ := newrelic.NewApplication(
		newrelic.ConfigAppName("My app"),
		newrelic.ConfigLicense(sampleLicense),
	)
	testTxn = app.StartTransaction("txn-name")

	res := m.Run()
	// TODO: because we copied this, this is not how we normally do cleanups
	// this will result in extra db being present if we crash during test
	_ = os.Remove("./foo.db")
	os.Exit(res)
}

func TestWrappedGorm(t *testing.T) {
	segmentsHistory = []*nrmock.DatastoreSegment{}
	txn := SetTxnToGorm(testTxn, db)
	txnDB := txn.WithContext(nil)
	dbInsert(t, txnDB)
	lastSegment := segmentsHistory[0]
	if lastSegment.Product != newrelic.DatastoreSQLite {
		t.Errorf("wrong product: %v", lastSegment.Product)
	}
	if lastSegment.ParameterizedQuery != "INSERT INTO `models` (`value`) VALUES (?)" {
		t.Errorf("wrong query: %v", lastSegment.ParameterizedQuery)
	}
	if lastSegment.Operation != "INSERT" {
		t.Errorf("wrong operation: %v", lastSegment.Operation)
	}
	if lastSegment.Collection != "models" {
		t.Errorf("wrong collection: %v", lastSegment.Collection)
	}
	// gorm always tries to wrap insert into transaction
	lastSegment = segmentsHistory[1]
	if lastSegment.Operation != "COMMIT/ROLLBACK" {
		t.Errorf("wrong operation: %v", lastSegment.Operation)
	}
	if lastSegment.Collection != "models" {
		t.Errorf("wrong collection: %v", lastSegment.Collection)
	}
	// no transaction on select
	dbSelect(t, txnDB)

	// to update we have to create a row first +2 transactions
	dbUpdate(t, txnDB)
	lastSegment = segmentsHistory[5]
	if lastSegment.Operation != "UPDATE" {
		t.Errorf("wrong operation: %v", lastSegment.Operation)
	}
	// gorm always tries to wrap update into transaction
	lastSegment = segmentsHistory[6]
	if lastSegment.Operation != "COMMIT/ROLLBACK" {
		t.Errorf("wrong operation: %v", lastSegment.Operation)
	}
	if lastSegment.Collection != "models" {
		t.Errorf("wrong collection: %v", lastSegment.Collection)
	}
	// to delete we have to create a row first +2 transactions
	dbDelete(t, txnDB)

	lastSegment = segmentsHistory[9]
	if lastSegment.Operation != "DELETE" {
		t.Errorf("wrong operation: %v", lastSegment.Operation)
	}
	// gorm always tries to wrap delete into transaction
	lastSegment = segmentsHistory[10]
	if lastSegment.Operation != "COMMIT/ROLLBACK" {
		t.Errorf("wrong operation: %v", lastSegment.Operation)
	}
	if lastSegment.Collection != "models" {
		t.Errorf("wrong collection: %v", lastSegment.Collection)
	}

	dbSelectNoRecord(t, txnDB)
	lastSegment = segmentsHistory[11]
	if lastSegment.Operation != "SELECT" {
		t.Error("must report SELECT operation even no record result")
	}
	if lastSegment.ParameterizedQuery != "SELECT * FROM `models` WHERE `models`.`value` = ? ORDER BY `models`.`id` LIMIT 1" {
		t.Error("wrong query", lastSegment.ParameterizedQuery)
	}

	historyLen := len(segmentsHistory)
	dbInsert(t, db)
	dbSelect(t, db)
	dbUpdate(t, db)
	dbDelete(t, db)
	if len(segmentsHistory) > historyLen {
		t.Error("main db was affected")
	}
}

func TestDBManualTransaction(t *testing.T) {
	segmentsHistory = []*nrmock.DatastoreSegment{}
	txnDB := SetTxnToGorm(testTxn, db)

	// when transaction has been started manually gorm won't wrap insert/update/delete into another tx
	// commit time in such case must be measured manually by the user
	tx := txnDB.Begin()

	m := &Model{Value: "manual-tx-test"}
	if err := tx.Create(m).Error; err != nil {
		t.Error(err)
	}
	m.Value = "updated"
	if err := tx.Save(m).Error; err != nil {
		t.Error(err)
	}
	if err := tx.Delete(m).Error; err != nil {
		t.Error(err)
	}

	tx.Commit()

	if len(segmentsHistory) != 3 {
		t.Errorf("expected 3 segments, got: %v", len(segmentsHistory))
	}
	op := segmentsHistory[0].Operation
	if op != "INSERT" {
		t.Errorf("wrong operation: %v", op)
	}
	op = segmentsHistory[1].Operation
	if op != "UPDATE" {
		t.Errorf("wrong operation: %v", op)
	}
	op = segmentsHistory[2].Operation
	if op != "DELETE" {
		t.Errorf("wrong operation: %v", op)
	}
}

func dbInsert(t *testing.T, db *gorm.DB) {
	if err := db.Create(&Model{Value: "test"}).Error; err != nil {
		t.Error(err)
	}
}

func dbSelect(t *testing.T, db *gorm.DB) {
	if err := db.First(&Model{Value: "to-select"}).Error; err != nil {
		t.Error(err)
	}
}

func dbSelectNoRecord(t *testing.T, db *gorm.DB) {
	if err := db.Where(Model{Value: "not found"}).First(&Model{}).Error; err != gorm.ErrRecordNotFound {
		t.Error(err)
	}
}

func dbUpdate(t *testing.T, db *gorm.DB) {
	m := &Model{Value: "to-update"}
	if err := db.Create(m).Error; err != nil {
		t.Error(err)
	}
	m.Value = "updated"
	if err := db.Save(m).Error; err != nil {
		t.Error(err)
	}
}

func dbDelete(t *testing.T, db *gorm.DB) {
	m := &Model{Value: "to-delete"}
	if err := db.Create(m).Error; err != nil {
		t.Error(err)
	}
	if err := db.Delete(m).Error; err != nil {
		t.Error(err)
	}
}
