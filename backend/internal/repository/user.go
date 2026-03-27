package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/abhinandpn/GoTicks/backend/internal/auth"
	"github.com/abhinandpn/GoTicks/backend/internal/models"
	"github.com/google/uuid"
)

func CreateUser(ctx context.Context, db *sql.DB, user models.User) error {
	query := `
		INSERT INTO users (id, name, number, email, password_hash, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	Password, err := auth.HashPassword(user.PasswordHash)
	if err != nil {
		return err
	}
	fmt.Printf("Hashed password: %s\n", Password)

	// Ensure values are set (important)
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	_, err = db.ExecContext(
		ctx,
		query,
		user.ID,
		user.Name,
		user.Number,
		user.Email,
		Password,
		user.CreatedAt,
	)

	return err
}

func UpdateUser(ctx context.Context, db *sql.DB, user models.User) error {

	// update make pending for - get the user info / get the update user info (need to add)
	query := `
		UPDATE users
		SET name = $1,
		    number = $2,
		    email = $3,
		    password_hash = $4
		WHERE id = $5
	`

	result, err := db.ExecContext(
		ctx,
		query,
		user.Name,
		user.Number,
		user.Email,
		user.PasswordHash,
		user.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %s", user.ID)
	}

	return nil
}
func DeleteUser(ctx context.Context, db *sql.DB, userID uuid.UUID) error {

	query := `DELETE FROM users WHERE id = $1`

	result, err := db.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %s", userID)
	}

	return nil
}
