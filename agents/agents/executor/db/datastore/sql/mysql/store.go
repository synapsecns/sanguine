package mysql

import (
	"context"
	"fmt"
	hostfile "github.com/flowerwrong/go-hostsfile"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/datastore/sql/base"
	common_base "github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"gorm.io/gorm/schema"
	"net"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Store is the sqlite store. It extends the base store for sqlite specific queries.
type Store struct {
	*base.Store
}

// MaxIdleConns is exported here for testing. Tests execute too slowly with a reconnect each time.
var MaxIdleConns = 10

// NamingStrategy is for table prefixes.
var NamingStrategy = schema.NamingStrategy{}

// NewMysqlStore creates a new mysql store for a given data store.
func NewMysqlStore(ctx context.Context, dbURL string, handler metrics.Handler) (*Store, error) {
	logger.Debug("creating mysql store")

	//TODO: REMOVE ME OR @TRAJAN0x will get pissed
	// this is a fun little workaround for a bug in the way dns resolution works when using kubefwd
	// for some unknown reason (I suspect it's netgo vs cgo) the dns resolution fails when using kubefwd
	// in mysql. If this is a persistent issue and kubefwd is used a lot, it might be worth looking into how the mysql driver does dns resolution
	// and merging there
	if false {
		const amysql = "agents-mysql"
		res, err := net.LookupIP("scribe-mysql-metrics")
		if err != nil {
			panic(res)
		}
		ip, err := hostfile.Lookup(amysql)
		if err != nil {
			return nil, fmt.Errorf("could not lookup ip for mysql: %w", err)
		}

		dbURL = strings.Replace(dbURL, amysql, ip, 1)
	}

	gdb, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{
		Logger:               common_base.GetGormLogger(logger),
		FullSaveAssociations: true,
		NamingStrategy:       NamingStrategy,
		NowFunc:              time.Now,
	})

	if err != nil {
		return nil, fmt.Errorf("could not create mysql connection: %w", err)
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		return nil, fmt.Errorf("could not get sql db: %w", err)
	}

	// fixes a timeout issue https://stackoverflow.com/a/42146536
	sqlDB.SetMaxIdleConns(MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	handler.AddGormCallbacks(gdb)

	err = gdb.WithContext(ctx).AutoMigrate(base.GetAllModels()...)

	if err != nil {
		return nil, fmt.Errorf("could not migrate on mysql: %w", err)
	}
	return &Store{base.NewStore(gdb)}, nil
}

// var _ db.Service = &Store{}
