package user

import (
	"context"
	"errors"
	"giants/pkg/entity/user"
)

type FindUser struct {
	userStore user.IUserStore
}

func NewFindUser(us user.IUserStore) *FindUser {
	return &FindUser{userStore: us}
}

func (uc FindUser) ByEmail(email string) (*user.User, error) {
	uObj, err := uc.userStore.FindByEmail(context.Background(), email)
	if err != nil {
		return nil, errors.New("user with this email does not exist")
	}

	return uObj, nil
}

func (uc FindUser) ById(id string) (*user.User, error) {
	uObj, err := uc.userStore.FindById(context.Background(), id)
	if err != nil {
		return nil, errors.New("user with this id does not exist")
	}

	return uObj, nil
}
