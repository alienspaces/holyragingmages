package store

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // blank import intended
)

// newPostgresDB -
func newPostgresDB(c Configurer, l Logger) (*sqlx.DB, error) {

	dbHost := c.Get("APP_DB_HOST")
	dbPort := c.Get("APP_DB_PORT")
	dbName := c.Get("APP_DB_NAME")
	dbUser := c.Get("APP_DB_USER")
	dbPass := c.Get("APP_DB_PASSWORD")

	cs := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPass, dbName, dbHost, dbPort)

	l.Printf("Connect string %s", cs)

	d, err := sqlx.Connect("postgres", cs)
	if err != nil {
		l.Printf("Failed to db connect >%v<", err)
		return nil, err
	}

	return d, nil
}
