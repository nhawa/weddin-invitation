package model

import (
	"database/sql"
	"time"
)

type (
	Customer struct {
		ID          int64          `db:"id"`
		FirstName   string         `db:"first_name"`
		LastName    sql.NullString `db:"last_name"`
		PhoneNumber sql.NullString `db:"phone_number"`
		Enabled     bool           `db:"enabled"`
		CreatedAt   time.Time      `db:"created_at"`
	}
)
