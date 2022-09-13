package clickhouse

import (
	"context"
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/services/explorer/db/datastore/sql/test"
	gormClickhouse "gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"

	"strconv"
)

// Store is the clickhouse store. It extends the base store for sqlite specific queries.
type Store struct {
	*test.Store
}

// NamingStrategy is exported here for testing.
var NamingStrategy = schema.NamingStrategy{}

// NewClickhouseStore creates a new mysql store for a given data store.
func NewClickhouseStore(ctx context.Context, dbURL string) (*Store, error) {

	pool, err := dockertest.NewPool("")

	if err != nil {
		fmt.Printf("fail")

	}
	// pulls an image, creates a container based on it and runs it
	fmt.Printf("2")
	runOptions := &dockertest.RunOptions{
		Repository: "clickhouse/clickhouse-server",
		Tag:        "latest",
		Env: []string{
			"CLICKHOUSE_DB=" + "clickhouse",
			"CLICKHOUSE_USER=" + "clickhouseUser",
			"CLICKHOUSE_PASSWORD=" + "clickhouse",
			"CLICKHOUSE_DEFAULT_ACCESS_MANAGEMENT=" + "1",
		},
		Labels:       map[string]string{"goose_test": "1"},
		PortBindings: make(map[docker.Port][]docker.PortBinding),
	}
	runOptions.PortBindings[docker.Port("9000/tcp")] = []docker.PortBinding{
		{HostPort: strconv.Itoa(9000)},
	}
	resource, err := pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})

	if err != nil {
		fmt.Printf("fail")
	}

	// Fetch port assigned to container
	//address := fmt.Sprintf("%s:%s", "localhost", resource.GetPort("9000/tcp"))
	address := dbURL
	cleanup := func() {
		if err := pool.Purge(resource); err != nil {
			logger.Debug("failed to purge resource: %v", err)
		}
	}
	var db *Store
	if err := pool.Retry(func() error {
		db, err = openGormClickhouse(ctx, address)
		return nil
	}); err != nil {
		//destroy container
		cleanup()
		return nil, err
	}
	logger.Debug("created clickhouse store")
	return db, nil
}

func openGormClickhouse(ctx context.Context, address string) (*Store, error) {
	gdb, err := gorm.Open(gormClickhouse.Open(address), &gorm.Config{
		Logger:               dbcommon.GetGormLogger(logger),
		FullSaveAssociations: true,
		NamingStrategy:       NamingStrategy,
		NowFunc:              time.Now,
	})

	if err != nil {
		return nil, fmt.Errorf("could not create clickhouse connection: %w", err)
	}

	clickhouseDB, err := gdb.DB()
	if err != nil {
		return nil, fmt.Errorf("could not get sql db: %w", err)
	}

	clickhouseDB.SetMaxIdleConns(0)

	err = gdb.WithContext(ctx).AutoMigrate(test.GetAllModels()...)

	if err != nil {
		return nil, fmt.Errorf("could not migrate on clickhouse: %w", err)
	}
	return &Store{test.NewStore(gdb)}, nil
}

// var _ db.Service = &Store{}
