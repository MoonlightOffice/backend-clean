package payment

import "context"

type IPaymentStore interface {
	// errors: ErrDuplicate
	Add(context.Context, *Payment) error

	// errors: ErrNotFound
	Find(ctx context.Context, userId, paymentId string) (*Payment, error)
}
