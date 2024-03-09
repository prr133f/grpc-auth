package database

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PgxIface interface {
	Begin(context.Context) (pgx.Tx, error)
	Close()
	Ping(context.Context) error
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(context context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

type Postgres struct {
	DB  PgxIface
	Log *zap.Logger
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

func NewPG(ctx context.Context, DSN string, logger *zap.Logger) (*Postgres, error) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, DSN)
		if err != nil {
			logger.Fatal("error while creating pool", zap.Error(err))
		}

		pgInstance = &Postgres{DB: db, Log: logger}
	})

	return pgInstance, nil
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.DB.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.DB.Close()
}
