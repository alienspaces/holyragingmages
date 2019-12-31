package service

import (
	"github.com/jmoiron/sqlx"
)

// Database provides methods to get database connection
type Database interface {
	GetConnection() *sqlx.DB
	CloseConnection()
	BeginTrx() *sqlx.Tx
}
