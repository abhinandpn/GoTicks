package models

import (
	"time"

	"github.com/google/uuid"
)

type Theater struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Location    string `json:"location" db:"location"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type Screen struct {
	Id          uuid.UUID `json:"id" db:"id"`
	TheaterId   uuid.UUID `json:"theater_id" db:"theater_id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	TotalSeats  int    `json:"total_seats" db:"total_seats"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type Show struct {
	Id        uuid.UUID `json:"id" db:"id"`
	MovieId   uuid.UUID `json:"movie_id" db:"movie_id"`
	ScreenId  uuid.UUID `json:"screen_id" db:"screen_id"`
	StartTime time.Time `json:"start_time" db:"start_time"`
	EndTime   time.Time `json:"end_time" db:"end_time"`
}

type Ticket struct {
	Id       uuid.UUID `json:"id" db:"id"`
	ShowId   uuid.UUID `json:"show_id" db:"show_id"`
	SeatType string  `json:"seat_type" db:"seat_type"`
	Price    float64 `json:"price" db:"price"`
}
