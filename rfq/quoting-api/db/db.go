package db

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core"
	common_base "github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/ethergo/submitter/db/txdb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/schema"
	"os"
	"time"

	"github.com/synapsecns/sanguine/rfq/quoting-api/db/models"

	"github.com/synapsecns/sanguine/core/metrics"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
	"gorm.io/gorm"
)

const (
	maxOpenConns = 1048
	maxIdleConns = 1048
)

type Database struct {
	*gorm.DB
	metrics        metrics.Handler
	submitterStore submitterDB.Service
}

func NewDatabase(ctx context.Context, handler metrics.Handler, skipMigrations bool) (db *Database, err error) {
	// @Bobby lets adjust this to how we want the config to work
	dbname := os.Getenv("MYSQL_DATABASE")
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", core.GetEnv("MYSQL_USER", "root"), os.Getenv("MYSQL_PASSWORD"), core.GetEnv("MYSQL_HOST", "127.0.0.1"), core.GetEnvInt("MYSQL_PORT", 3306), dbname)

	logger.Debug("creating mysql store")
	ctx, span := handler.Tracer().Start(ctx, "start-mysql")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// Modify this as we need
	gdb, err := gorm.Open(mysql.Open(connString), &gorm.Config{
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

func (db *Database) InsertQuote(c context.Context, q *models.Quote) (id uint, err error) {
	if err = db.WithContext(c).Create(q).Error; err != nil {
		return
	}
	id = q.ID
	return
}

func (db *Database) UpdateQuote(c context.Context, q *models.Quote) error {
	return db.WithContext(c).Save(q).Error
}

func (db *Database) GetQuote(c context.Context, id uint) (q models.Quote, err error) {
	result := db.WithContext(c).First(&q, "id = ?", id)
	err = result.Error
	return
}

func (db *Database) GetQuotes(c context.Context, req *models.Request) (qs []models.Quote, err error) {
	result := db.WithContext(c).Where(&models.Quote{
		OriginChainId: req.OriginChainId,
		DestChainId:   req.DestChainId,
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

func (db *Database) DeleteQuote(c context.Context, id uint) error {
	return db.WithContext(c).Delete(&models.Quote{}, id).Error
}
