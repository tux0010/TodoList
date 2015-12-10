package database

import (
	"database/sql"
)

type DatabaseDriver interface {
	SetDBConnection(conn *sql.DB)
	ExecuteQuery(query string) error
	ExceuteQueryWithResponse(query string) (interface{}, error)
}

type Database struct {
	db *sql.DB
}

// SetDBConnection will set the database connection object obtained from
// the sql.Open() method
func (d *Database) SetDBConnection(conn *sql.DB) {
	d.db = conn
}

// ExecuteQuery will run a Database query and return an error if it fails
func (d *Database) ExecuteQuery(query string) error {
	_, err := d.db.Exec(query)

	return err
}

// ExecuteQueryWithResponse will run a Database query and return the resultant rows
func (d *Database) ExecuteQueryWithResponse(query string) (interface{}, error) {
	rows, err := d.db.Query(query)

	return rows, err
}
