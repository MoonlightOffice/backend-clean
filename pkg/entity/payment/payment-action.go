package payment

import "context"

type IPaymentAction interface {
	Make(context.Context, *Payment) error

	UpdateStatus(ctx context.Context, userId, paymentId string) (*Payment, error)
}
