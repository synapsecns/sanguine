// Package db contains the database implementation of the quoting api.
package db

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core"
	common_base "github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/ethergo/submitter/db/txdb"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/db/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/schema"
	"os"
	"time"

	"github.com/synapsecns/sanguine/core/metrics"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
	"gorm.io/gorm"
)

const (
	maxOpenConns = 1048
)

// Database is the database object.
type Database struct {
	*gorm.DB
	metrics        metrics.Handler
	submitterStore submitterDB.Service
}

// NewDatabase creates a new instance of the database.
// TODO: consider returning an interface instead of a concrete type.
// TODO: make dbType a string
func NewDatabase(ctx context.Context, handler metrics.Handler, skipMigrations bool, dbType, dsn string) (db *Database, err error) {
	// @Bobby lets adjust this to how we want the config to work
	var dialector gorm.Dialector
	// TODO: clean this up
	if dbType == "sqlite" {
		dialector = sqlite.Open(dsn)
	} else {
		connString := dsn
		// fallback to env vars
		if dsn == "" {
			dbname := os.Getenv("MYSQL_DATABASE")
			connString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", core.GetEnv("MYSQL_USER", "root"), os.Getenv("MYSQL_PASSWORD"), core.GetEnv("MYSQL_HOST", "127.0.0.1"), core.GetEnvInt("MYSQL_PORT", 3306), dbname)
		}
		dialector = mysql.Open(connString)
	}

	logger.Debug("creating store")
	ctx, span := handler.Tracer().Start(ctx, "start-db")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// Modify this as we need
	gdb, err := gorm.Open(dialector, &gorm.Config{
		Logger:                                   common_base.GetGormLogger(logger),
		FullSaveAssociations:                     true,
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		NamingStrategy:                           schema.NamingStrategy{},
	})

	if err != nil {
		return nil, fmt.Errorf("could not create mysql connection: %w", err)
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		return nil, fmt.Errorf("could not get sql db: %w", err)
	}

	// fixes a timeout issue https://stackoverflow.com/a/42146536
	sqlDB.SetMaxIdleConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxOpenConns(maxOpenConns)

	handler.AddGormCallbacks(gdb)

	if !skipMigrations {
		err = gdb.Transaction(func(tx *gorm.DB) error {
			//nolint: wrapcheck
			return gdb.WithContext(ctx).AutoMigrate(models.GetAllModels()...)
		})
	}
	if err != nil {
		return nil, fmt.Errorf("could not migrate on mysql: %w", err)
	}

	txDB := txdb.NewTXStore(gdb, handler)

	return &Database{gdb, handler, txDB}, nil
}

// InsertQuote inserts a quote into the database.
func (db *Database) InsertQuote(c context.Context, q *models.Quote) (id uint, err error) {
	if err = db.WithContext(c).Create(q).Error; err != nil {
		return
	}
	id = q.ID
	return
}

// UpdateQuote updates a quote in the database.
func (db *Database) UpdateQuote(c context.Context, q *models.Quote) error {
	return db.WithContext(c).Save(q).Error
}

// GetQuote gets a quote from the database.
func (db *Database) GetQuote(c context.Context, id uint) (q models.Quote, err error) {
	result := db.WithContext(c).First(&q, "id = ?", id)
	err = result.Error
	return
}

// GetQuotes gets quotes from the database.
func (db *Database) GetQuotes(c context.Context, req *models.Request) (qs []models.Quote, err error) {
	result := db.WithContext(c).Where(&models.Quote{
		OriginChainID: req.OriginChainID,
		DestChainID:   req.DestChainID,
		OriginToken:   req.OriginToken,
		DestToken:     req.DestToken,
	}).Where(
		"origin_amount >= ?",
		req.OriginAmount,
	).Where(
		"updated_at >= ?",
		req.UpdatedAtLast,
	).Order("price desc").Find(&qs)
	err = result.Error
	return
}

// DeleteQuote deletes a quote from the database.
func (db *Database) DeleteQuote(c context.Context, id uint) error {
	return db.WithContext(c).Delete(&models.Quote{}, id).Error
}
