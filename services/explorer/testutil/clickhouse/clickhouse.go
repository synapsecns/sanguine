package clickhouse

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/core/dockerutil"
)

// NewClickhouseStore creates a new clickhouse db hosted at localhost:xxxx with ory/dockertest.
func NewClickhouseStore(src string) (func(), *int, error) {
	timeout := time.Second
	port := freeport.GetPort()
	portStr := strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", net.JoinHostPort("localhost", portStr), timeout)
	if conn != nil {
		fmt.Println("Connection error to port: " + portStr)
		if conn.Close() != nil {
			return nil, nil, err
		}
		return nil, nil, fmt.Errorf("port %d is already in use: %w", port, err)
	}
	fmt.Println("Starting clickhouse docker swap on port: ", portStr)

	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, nil, fmt.Errorf("could not create docker swap: %w", err)
	}
	// pulls an image, creates a container based on it and runs it
	runOptions := &dockertest.RunOptions{
		Repository: "yandex/clickhouse-server",
		Tag:        "latest",
		Env: []string{
			"CLICKHOUSE_DB=" + "clickhouse_test",
			"CLICKHOUSE_USER=" + "clickhouse_test",
			"CLICKHOUSE_PASSWORD=" + "clickhouse_test",
			"CLICKHOUSE_DEFAULT_ACCESS_MANAGEMENT=" + "1",
		},
		// Label format: clickhouse_test_<src of test>_<port running>
		Labels: map[string]string{"clickhouse_test_" + src + "_" + portStr: "1"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"9000/tcp": {{HostIP: "localhost", HostPort: portStr + "/tcp"}},
		},
		ExposedPorts: []string{"9000"},
	}
	resource, err := pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})

	// Fetch port assigned to container
	address := fmt.Sprintf("%s:%s", "localhost", dockerutil.GetPort(resource, "9000/tcp"))

	// Docker will hard kill the container in 360 seconds (this is a test env).
	// In a continuous integration environment, this is increased to allow for the lower cpu count
	resourceLifetime := uint(360)
	pool.MaxWait = time.Minute * 2

	if os.Getenv("CI") != "" {
		resourceLifetime = 900
		pool.MaxWait = time.Minute * 5
	}

	if resource.Expire(resourceLifetime) != nil {
		return nil, nil, err
	}

	// Teardown function
	cleanup := func() {
		fmt.Println("Destroying container")
		if err := pool.Purge(resource); err != nil {
			fmt.Printf("failed to purge resource: %v \n", err)
		}
	}

	var db *sql.DB
	fmt.Println("Pinging clickhouse db...")

	if err := pool.Retry(func() error {
		db = clickHouseOpenDB(address)
		return db.Ping()
	}); err != nil {
		fmt.Printf("Could not connect to docker database: %v \n", err)
		return cleanup, nil, fmt.Errorf("could not connect to docker database: %w", err)
	}
	if err != nil {
		return cleanup, nil, fmt.Errorf("could not run resource: %w", err)
	}
	return cleanup, &port, nil
}

func clickHouseOpenDB(address string) *sql.DB {
	db := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{address},
		Auth: clickhouse.Auth{
			Database: "clickhouse_test",
			Username: "clickhouse_test",
			Password: "clickhouse_test",
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
	})
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Hour)
	return db
}
