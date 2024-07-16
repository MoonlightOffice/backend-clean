package user

import "context"

type IUserStore interface {
	// errors: ErrDuplicate
	Add(ctx context.Context, uObj *User) error

	// errors: ErrNotFound
	FindById(ctx context.Context, userId string) (*User, error)

	FindByEmail(ctx context.Context, email string) (*User, error)
}
