package db

import (
	"context"
	"errors"
	"testing"

	"giants/pkg/detail/db/postgres"
	"giants/pkg/entity/user"
)

func TestUser(t *testing.T) {
	defer postgres.DeleteAll()

	userStore := NewUserStore()

	// Add a new user
	uObj, _ := user.NewUser("u1")

	err := userStore.Add(context.Background(), uObj)
	if err != nil {
		t.Fatal(err)
	}

	// Check if duplicated error occurs
	err = userStore.Add(context.Background(), uObj)
	if !errors.Is(err, user.ErrDuplicate) {
		t.Fatal("Expected ErrDuplicated")
	}

	// Fetch user from db and check
	_, err = userStore.FindById(context.Background(), uObj.UserId)
	if err != nil {
		t.Fatal(err)
	}
}
