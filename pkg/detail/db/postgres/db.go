package postgres

import (
	"context"
	"giants/pkg/util"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pool     *pgxpool.Pool
	maxConns = 4
	url      = os.Getenv("DATABASE_URL")
	timeout  = 10 * time.Second
)

type Postgres struct {
	*pgxpool.Conn
}

func NewPostgres() (*Postgres, error) {
	// Check if pool is initialized
	if pool == nil {
		pgxConfig, err := pgxpool.ParseConfig(url)
		if err != nil {
			return nil, util.ErrBuilder(err)
		}

		pgxConfig.MaxConns = int32(maxConns)
		pgxConfig.MaxConnIdleTime = time.Hour

		pool, err = pgxpool.NewWithConfig(context.Background(), pgxConfig)
		if err != nil {
			return nil, util.ErrBuilder(err)
		}
	}

	// Acquire connection
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()

	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, util.ErrBuilder(err)
	}

	return &Postgres{conn}, nil
}
