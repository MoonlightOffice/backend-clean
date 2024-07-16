package payment

import (
	"time"

	"giants/pkg/util"
)

type Payment struct {
	UserId    string
	PaymentId string
	CreatedAt time.Time
	Amount    int64
	Currency  Currency
	Metadata  map[string]interface{}
}

type Currency string

const (
	CurrencyJPY Currency = "jpy"
	CurrencyUSD Currency = "usd"
)

func NewPayment(userId string, amount int64, currency Currency) *Payment {
	return &Payment{
		UserId:    userId,
		PaymentId: util.NewID(),
		CreatedAt: time.Now(),
		Amount:    amount,
		Currency:  currency,
	}
}
