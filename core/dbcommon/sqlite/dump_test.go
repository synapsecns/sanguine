package sqlite3dump

import (
	"bufio"
	"bytes"
	"database/sql"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertEqualIgnoreLineSeparators(t *testing.T, expected []byte, actual []byte) {
	expectedStr := strings.ReplaceAll(string(expected), "\r\n", "\n")
	actualStr := strings.ReplaceAll(string(actual), "\r\n", "\n")
	assert.Equal(t, expectedStr, actualStr)
}

func TestCars(t *testing.T) {
	var b bytes.Buffer
	out := bufio.NewWriter(&b)
	err := Dump("testdata/cars.db", out)
	assert.Nil(t, err)
	out.Flush()
	pythonOutput, _ := ioutil.ReadFile("testdata/python.sql")
	assertEqualIgnoreLineSeparators(t, pythonOutput, b.Bytes())
	ioutil.WriteFile("out.sql", b.Bytes(), 0644)
}

func TestMigrate(t *testing.T) {
	var b bytes.Buffer
	out := bufio.NewWriter(&b)

	db, err := sql.Open("sqlite3", "testdata/cars.db")
	assert.Nil(t, err)
	defer db.Close()

	err = DumpMigration(db, out)
	assert.Nil(t, err)

	out.Flush()
	pythonOutput, _ := ioutil.ReadFile("testdata/migrate.sql")
	assertEqualIgnoreLineSeparators(t, pythonOutput, b.Bytes())
}

func TestDump(t *testing.T) {
	cases := map[string]struct {
		dbFile     string
		expectFile string
		options    []Option
	}{
		"No Options": {
			dbFile:     "cars.db",
			expectFile: "python.sql",
		},
		"WithMigration": {
			dbFile:     "cars.db",
			expectFile: "migrate.sql",
			options:    []Option{WithMigration()},
		},
		"WithDropIfExists": {
			dbFile:     "cars.db",
			expectFile: "drop_if_exists.sql",
			options:    []Option{WithDropIfExists(true)},
		},
		"WithTransaction - false": {
			dbFile:     "cars.db",
			expectFile: "without_tx.sql",
			options:    []Option{WithTransaction(false)},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			expect, err := ioutil.ReadFile(filepath.Join("testdata", c.expectFile))
			require.NoError(t, err, "failed to open expect file")

			var b bytes.Buffer
			out := bufio.NewWriter(&b)
			dbFilePath := filepath.Join("testdata", c.dbFile)
			err = Dump(dbFilePath, out, c.options...)
			require.NoError(t, err)
			out.Flush()

			got := b.Bytes()
			assertEqualIgnoreLineSeparators(t, expect, got)
		})
	}
}
