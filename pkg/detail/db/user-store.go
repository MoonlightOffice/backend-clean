package db

import (
	"context"
	"errors"
	"time"

	"giants/pkg/detail/db/postgres"
	"giants/pkg/entity/user"
	"giants/pkg/util"

	"github.com/jackc/pgx/v5"
)

type userStore struct{}

func NewUserStore() user.IUserStore {
	return userStore{}
}

func (us userStore) Add(ctx context.Context, uObj *user.User) error {
	pool, err := postgres.NewPostgres()
	if err != nil {
		return err
	}
	defer pool.Release()

	stmt := `INSERT INTO users (user_id, email, created_at) VALUES ($1, $2, $3)`
	_, err = pool.Exec(
		ctx,
		stmt,
		uObj.UserId,
		uObj.Email,
		uObj.CreatedAt,
	)
	if err != nil {
		if postgres.IsErrDuplicate(err) {
			return util.ErrBuilder(user.ErrDuplicate)
		}

		return util.ErrBuilder(err)
	}

	return nil
}

// errors: ErrNotFound
func (us userStore) FindById(ctx context.Context, userId string) (*user.User, error) {
	pool, err := postgres.NewPostgres()
	if err != nil {
		return nil, err
	}
	defer pool.Release()

	var (
		email     string
		createdAt time.Time
	)

	stmt := `SELECT email, created_at FROM users WHERE user_id = $1 FOR UPDATE`
	err = pool.QueryRow(ctx, stmt, userId).Scan(&email, &createdAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, util.ErrBuilder(user.ErrNotFound)
		}

		return nil, util.ErrBuilder(err)
	}

	return &user.User{
		UserId:    userId,
		Email:     email,
		CreatedAt: createdAt,
	}, nil
}

func (us userStore) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	pool, err := postgres.NewPostgres()
	if err != nil {
		return nil, err
	}
	defer pool.Release()

	var (
		userId    string
		createdAt time.Time
	)

	stmt := `SELECT user_id, created_at FROM users WHERE email = $1 FOR UPDATE`
	err = pool.QueryRow(ctx, stmt, email).Scan(&userId, &createdAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, util.ErrBuilder(user.ErrNotFound)
		}

		return nil, util.ErrBuilder(err)
	}

	return &user.User{
		UserId:    userId,
		Email:     email,
		CreatedAt: createdAt,
	}, nil
}
