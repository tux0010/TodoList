package database

import (
	"database/sql"
)

type DatabaseDriver interface {
	ExecuteQuery(query string, values ...interface{}) error
	ExecuteQueryWithResponse(query string, values ...interface{}) (interface{}, error)
}

type Database struct {
	db *sql.DB
}

func NewDatabase(conn *sql.DB) *Database {
	return &Database{db: conn}
}

// ExecuteQuery will run a Database query and return an error if it fails
func (d *Database) ExecuteQuery(query string, values ...interface{}) error {
	_, err := d.db.Exec(query, values...)

	return err
}

// ExecuteQueryWithResponse will run a Database query and return the resultant rows
func (d *Database) ExecuteQueryWithResponse(query string, values ...interface{}) (interface{}, error) {
	return d.db.Query(query, values...)
}
