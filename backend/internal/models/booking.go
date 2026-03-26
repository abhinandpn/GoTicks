package models

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID           uuid.UUID `json:"id" db:"id"`
	UserID       uuid.UUID `json:"user_id" db:"user_id"`
	ShowID       uuid.UUID `json:"show_id" db:"show_id"`
	TotalAmmount float64   `json:"total_amount" db:"total_amount"`
	Status       string    `json:"status" db:"status"`
	CreatedAt    time.Time    `json:"created_at" db:"created_at"`
}

type BookingSeat struct {
	Id        uuid.UUID `json:"id" db:"id"`
	BookingId uuid.UUID `json:"booking_id" db:"booking_id"`
	SeatId    uuid.UUID `json:"seat_id" db:"seat_id"`
	Price     float64   `json:"price" db:"price"`
}
