package payment

import "errors"

var (
	ErrNotFound  = errors.New("record does not exist")
	ErrDuplicate = errors.New("record already exists")
)
