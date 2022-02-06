package model

import (
	"database/sql"
	"time"
)

type (
	Order struct {
		ID           int64          `db:"id"`
		CustomerName string         `db:"customer_name"`
		Template     sql.NullString `db:"template"`
		Status       string         `db:"status"`
		CreatedAt    time.Time      `db:"created_at"`
		ExpiredAt    time.Time      `db:"expired_at"`
	}
)
