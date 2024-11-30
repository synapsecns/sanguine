package sqlite3dump

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type sqlite3dumper struct {
	migration           bool
	dropIfExists        bool
	wrapWithTransaction bool
}

func newSqlite3Dumper(opts ...Option) *sqlite3dumper {
	dumper := &sqlite3dumper{
		wrapWithTransaction: true,
	}

	if len(opts) == 0 {
		return dumper
	}

	for _, option := range opts {
		option(dumper)
	}
	return dumper
}

// DumpMigration will dump the database in an SQL text format
// and not include creation tables and will include table column names
//
// Deprecated, use WithMigration() option instead.
func DumpMigration(db *sql.DB, out io.Writer) (err error) {
	s3d := newSqlite3Dumper(WithMigration())
	return s3d.dumpDB(db, out)
}

// Dump will dump the database in an SQL text format into the specified io.Writer.
// Ported from the Python equivalent: https://github.com/python/cpython/blob/3.6/Lib/sqlite3/dump.py.
// Returns an error if the database doesn't exist.
func Dump(dbName string, out io.Writer, opts ...Option) (err error) {
	s3d := newSqlite3Dumper(opts...)
	return s3d.dump(dbName, out)
}

func (s3d *sqlite3dumper) dump(dbName string, out io.Writer) (err error) {
	// return if doesn't exist
	if _, err = os.Stat(dbName); os.IsNotExist(err) {
		return
	}

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return
	}
	defer db.Close()

	return s3d.dumpDB(db, out)
}

// DumpDB dumps a raw sql.DB
func DumpDB(db *sql.DB, out io.Writer, opts ...Option) (err error) {
	s3d := newSqlite3Dumper(opts...)
	return s3d.dumpDB(db, out)
}

func (s3d *sqlite3dumper) dumpDB(db *sql.DB, out io.Writer) (err error) {
	if s3d.wrapWithTransaction {
		out.Write([]byte("BEGIN TRANSACTION;\n"))
	}

	// sqlite_master table contains the SQL CREATE statements for the database.
	tableSchemas, err := s3d.getSchemas(db, `
        SELECT "name", "type", "sql"
        FROM "sqlite_master"
            WHERE "sql" NOT NULL AND
            "type" == 'table'
            ORDER BY "name"
		`)
	if err != nil {
		return err
	}

	// Now when the type is 'index', 'trigger', or 'view'
	otherSchemas, err := s3d.getSchemas(db, `
		SELECT "name", "type", "sql"
        FROM "sqlite_master"
            WHERE "sql" NOT NULL AND
            "type" IN ('index', 'trigger', 'view')
		`)
	if err != nil {
		return err
	}

	if s3d.dropIfExists {
		allSchemas := append(otherSchemas, tableSchemas...)
		if err := s3d.writeDropStatements(out, allSchemas); err != nil {
			return err
		}
	}

	for _, schema := range tableSchemas {
		if schema.Name == "sqlite_sequence" {
			out.Write([]byte(`DELETE FROM "sqlite_sequence";` + "\n"))
		} else if schema.Name == "sqlite3_stat1" {
			out.Write([]byte(`ANALYZE "sqlite_master";` + "\n"))
		} else if strings.HasPrefix(schema.Name, "sqlite_") {
			continue
			// # NOTE: Virtual table support not implemented
			// #elif sql.startswith('CREATE VIRTUAL TABLE'):
			// #    qtable = table_name.replace("'", "''")
			// #    yield("INSERT INTO sqlite_master(type,name,tbl_name,rootpage,sql)"\
			// #        "VALUES('table','{0}','{0}',0,'{1}');".format(
			// #        qtable,
			// #        sql.replace("''")))
		} else if strings.HasSuffix(schema.Name, "_segments") || strings.HasSuffix(schema.Name, "_segdir") || strings.HasSuffix(schema.Name, "_stat") || strings.HasSuffix(schema.Name, "_idx") || strings.HasSuffix(schema.Name, "_docsize") || strings.HasSuffix(schema.Name, "_config") || strings.HasSuffix(schema.Name, "_data") || strings.HasSuffix(schema.Name, "_content") {
			// these suffixes for tables are from using FTS5, and they should be ignored
			// because they are automatically created
			continue
		} else {
			if !s3d.migration {
				out.Write([]byte(fmt.Sprintf("%s;\n", schema.SQL)))
			}
		}

		// Build the insert statement for each row of the current table
		schema.Name = strings.Replace(schema.Name, `"`, `""`, -1)
		err = s3d.writeInsStmtsForTableRows(out, db, schema.Name)
		if err != nil {
			return err
		}
	}

	for _, schema := range otherSchemas {
		out.Write([]byte(fmt.Sprintf("%s;\n", schema.SQL)))
	}

	if s3d.wrapWithTransaction {
		out.Write([]byte("COMMIT;\n"))
	}

	return
}

func (s3d *sqlite3dumper) writeDropStatements(w io.Writer, schemas []schema) (err error) {
	for _, schema := range schemas {
		var statement string

		switch schema.Type {
		case "index":
			statement = fmt.Sprintf("DROP INDEX IF EXISTS %s;\n", schema.Name)
		case "table":
			if strings.HasPrefix(schema.Name, "sqlite_") {
				// skip system tables
				continue
			}

			statement = fmt.Sprintf("DROP TABLE IF EXISTS %s;\n", schema.Name)
		default:
			continue
		}

		_, err = w.Write([]byte(statement))
		if err != nil {
			return fmt.Errorf("failed to write '%q': %s", statement, err)
		}
	}

	return nil
}

func (s3d *sqlite3dumper) writeInsStmtsForTableRows(w io.Writer, db *sql.DB, tableName string) (err error) {
	// first get the column names
	columnNames, err := s3d.pragmaTableInfo(db, tableName)
	if err != nil {
		return
	}

	// sqlite_master table contains the SQL CREATE statements for the database.
	columnSelects := make([]string, len(columnNames))
	for i, c := range columnNames {
		columnSelects[i] = fmt.Sprintf(`'||quote("%s")||'`, strings.Replace(c, `"`, `""`, -1))
	}

	q := fmt.Sprintf(`
		SELECT 'INSERT INTO "%s" VALUES(%s)' FROM "%s";
	`,
		tableName,
		strings.Join(columnSelects, ","),
		tableName,
	)
	if s3d.migration {
		q = fmt.Sprintf(`
		SELECT 'INSERT INTO "%s"(%s) VALUES(%s)' FROM "%s";
	`,
			tableName,
			strings.Join(columnNames, ","),
			strings.Join(columnSelects, ","),
			tableName,
		)
	}

	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var insert string
		err = rows.Scan(&insert)
		if err != nil {
			return
		}
		_, err = w.Write([]byte(fmt.Sprintf("%s;\n", insert)))
		if err != nil {
			return
		}
	}
	return rows.Err()
}

func (s3d *sqlite3dumper) pragmaTableInfo(db *sql.DB, tableName string) (columnNames []string, err error) {
	// sqlite_master table contains the SQL CREATE statements for the database.
	q := `
        PRAGMA table_info("` + tableName + `")
		`
	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	columnNames = []string{}
	for rows.Next() {
		var arr []interface{}
		for i := 0; i < 6; i++ {
			arr = append(arr, new(interface{}))
		}
		err = rows.Scan(arr...)
		if err != nil {
			return
		}
		columnNames = append(columnNames,
			func() (result string) {
				// check the type
				switch (*arr[1].(*interface{})).(type) {
				case string:
					result = string((*arr[1].(*interface{})).(string))
				case []uint8:
					result = string((*arr[1].(*interface{})).([]uint8))
				}
				return
			}(),
		)
	}
	err = rows.Err()
	return
}

type schema struct {
	Name string
	Type string
	SQL  string
}

func (s3d *sqlite3dumper) getSchemas(db *sql.DB, q string) (schemas []schema, err error) {
	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	schemas = []schema{}
	for rows.Next() {
		s := schema{}
		err = rows.Scan(&s.Name, &s.Type, &s.SQL)
		if err != nil {
			return
		}
		schemas = append(schemas, s)
	}
	err = rows.Err()
	return
}
