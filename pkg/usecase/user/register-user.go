package user

import (
	"context"
	"errors"
	"giants/pkg/entity/user"
)

type RegisterUser struct {
	userStore user.IUserStore
}

func NewRegisterUser(us user.IUserStore) *RegisterUser {
	return &RegisterUser{userStore: us}
}

type RegisterUserInputDto struct {
	Email string
}

func (uc RegisterUser) Run(dto RegisterUserInputDto) error {
	uObj, ok := user.NewUser(dto.Email)
	if !ok {
		return errors.New("invalid email")
	}

	err := uc.userStore.Add(context.Background(), uObj)
	if err != nil {
		if errors.Is(err, user.ErrDuplicate) {
			return errors.New("email already taken")
		}

		return err
	}

	return nil
}
