package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Record -
type Record struct {
	ID        string         `db:"id"`
	CreatedAt string         `db:"created_at"`
	UpdatedAt sql.NullString `db:"updated_at"`
	DeletedAt sql.NullString `db:"deleted_at"`
}

// NewRecordID -
func NewRecordID() string {
	uuidByte, _ := uuid.NewRandom()
	uuidString := uuidByte.String()
	return uuidString
}

// NewCreatedAt -
func NewCreatedAt() string {
	return timestamp()
}

// NewUpdatedAt -
func NewUpdatedAt() sql.NullString {
	return NewNullString(timestamp())
}

// NewDeletedAt -
func NewDeletedAt() sql.NullString {
	return NewNullString(timestamp())
}

func timestamp() string {
	// UTC - "2006-01-02T15:04:05Z07:00"
	return time.Now().UTC().Format(time.RFC3339)
}

// NewNullString - converts string type to sql.NullString type
func NewNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
