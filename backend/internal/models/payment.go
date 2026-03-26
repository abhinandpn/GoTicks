package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	Id            uuid.UUID `json:"id" db:"id"`
	BookingId     uuid.UUID `json:"booking_id" db:"booking_id"`
	Amount        float64   `json:"amount" db:"amount"`
	Status        string    `json:"status" db:"status"`
	PaymentMethod string    `json:"payment_method" db:"payment_method"`
	PaymentTime   time.Time `json:"time" db:"time"`
}
