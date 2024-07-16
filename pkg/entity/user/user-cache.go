package user

import "context"

type IUserCache interface {
	// errors: ErrDuplicate
	Save(context.Context, *User) error
}
