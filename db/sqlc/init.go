package db

import (
	"database/sql"
	"fmt"
	"log"
)

// Database is the extension struct of Queries
type Database struct {
	*Queries
	Conn *sql.DB
}

// NewDatabase return a Database and connect it with *sql.DB
func NewDatabase(conn *sql.DB) *Database {
	return &Database{Conn: conn}
}

// Init will initialize stuff for database
func Init(dbDriver, dbUser, dbPassword, dbHost, dbPort, dbName, dbParams string) (*Database, error) {
	var database *Database

	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbParams)

	conn, err := sql.Open(dbDriver, url)
	if err != nil {
		return database, err
	}

	database = NewDatabase(conn)
	err = database.Conn.Ping()
	if err != nil {
		return database, err
	}

	log.Println("Database connected")
	return database, nil
}
