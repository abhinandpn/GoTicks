package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/abhinandpn/GoTicks/backend/internal/models"
	"github.com/google/uuid"
)

func CreateUser(ctx context.Context, db *sql.DB, user models.User) error {
	query := `
		INSERT INTO users (id, name, number, email, password_hash, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	// Ensure values are set (important)
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}
	_, err := db.ExecContext(
		ctx,
		query,
		user.ID,
		user.Name,
		user.Number,
		user.Email,
		user.PasswordHash,
		user.CreatedAt,
	)

	return err
}
