package clickhouse

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"net"
	"strconv"
	"time"
)

// NewClickhouseStore creates a new clickhouse db hosted at localhost:9000 with ory/dockertest
func NewClickhouseStore() error {

	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort("localhost", "9000"), timeout)
	if err != nil {
		fmt.Println("Port not open - attempting to create container...")
	}
	if conn != nil {
		defer conn.Close()
		fmt.Println("Port in user, exiting | ", net.JoinHostPort("localhost", "9000"))
		return fmt.Errorf("Port is already in use: %w", err)
	}

	pool, err := dockertest.NewPool("")

	if err != nil {
		return fmt.Errorf("could not create docker pool: %w", err)
	}
	// pulls an image, creates a container based on it and runs it
	runOptions := &dockertest.RunOptions{
		Repository: "clickhouse/clickhouse-server",
		Tag:        "latest",
		Env: []string{
			"CLICKHOUSE_DB=" + "clickhouse_test",
			"CLICKHOUSE_USER=" + "clickhouse_test",
			"CLICKHOUSE_PASSWORD=" + "clickhouse_test",
			"CLICKHOUSE_DEFAULT_ACCESS_MANAGEMENT=" + "1",
		},
		Labels:       map[string]string{"clickhouse_test": "1"},
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

	// Fetch port assigned to container
	address := fmt.Sprintf("%s:%s", "localhost", resource.GetPort("9000/tcp"))
	fmt.Println(address)

	// Docker will hard kill the container in 60 seconds (this is a test env remember)
	resource.Expire(60)

	var db *sql.DB
	fmt.Printf("Pinging new clickhouse db...\n")

	if err := pool.Retry(func() error {
		db = clickHouseOpenDB(address, nil)
		return db.Ping()
	}); err != nil {
		fmt.Errorf("could not connect to docker database: %w", err)
		if err := pool.Purge(resource); err != nil {
			fmt.Printf("failed to purge resource: %v", err)
		}
		return fmt.Errorf("could not connect to docker database: %w", err)
	}
	if err != nil {
		return fmt.Errorf("could not run resource: %w", err)
	}
	return nil
}

func clickHouseOpenDB(address string, tlsConfig *tls.Config) *sql.DB {
	db := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{address},
		Auth: clickhouse.Auth{
			Database: "clickhouse_test",
			Username: "clickhouse_test",
			Password: "clickhouse_test",
		},
		TLS: tlsConfig,
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		Debug: true,
	})
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Hour)
	return db
}
