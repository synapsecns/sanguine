package db

import (
	"context"

	"github.com/synapsecns/sanguine/rfq/quoting-api/config"
	"github.com/synapsecns/sanguine/rfq/quoting-api/db/models"

	"github.com/synapsecns/sanguine/core/metrics"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
	metrics        metrics.Handler
	submitterStore submitterDB.Service
}

// TODO:
func NewDatabase(ctx context.Context, cfg *config.Config) *Database {
	return &Database{}
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
