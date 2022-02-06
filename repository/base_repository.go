package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type (
	Repository interface {
		BeginTx() *sqlx.Tx
		GetContext() (context.Context, context.CancelFunc)
	}

	BaseRepository struct {
		contextTimeout time.Duration
		DB             *sqlx.DB
	}
)

func (b BaseRepository) BeginTx() *sqlx.Tx {
	return b.DB.MustBegin()
}

func (b BaseRepository) GetContext() (context.Context, context.CancelFunc) {
	ctx := context.Background()
	contextTimeout := time.Second * 2
	return context.WithTimeout(ctx, contextTimeout)
}
