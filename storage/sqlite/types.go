package sqlite

import (
	"time"
)

// Type should match the package name
const Type = "sqlite"

// Storage is a way to store checkup results in a SQLite database.
type Storage struct {
	// SqliteDBFile is the sqlite3 DB where check results will be stored.
	SqliteDBFile string `json:"sqlite_db_file,omitempty"`

	// Issue create statements for database schema
	Create bool `json:"create"`

	// Check files older than CheckExpiry will be
	// deleted on calls to Maintain(). If this is
	// the zero value, no old check files will be
	// deleted.
	CheckExpiry time.Duration `json:"check_expiry,omitempty"`
}

// schema is the expected table schema (can be re-applied)
const schema = `CREATE TABLE IF NOT EXISTS checks (
    name TEXT NOT NULL PRIMARY KEY,
    timestamp INT8 NOT NULL,
    results TEXT
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_checks_timestamp ON checks(timestamp);`