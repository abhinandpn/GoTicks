package repository

import (
	"context"
	"database/sql"
	"errors"
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

func GetUserByID(ctx context.Context, db *sql.DB, userID uuid.UUID) (*models.User, error) {

	if userID == uuid.Nil {
		return nil, errors.New("invalid user ID")
	}

	query := `
		SELECT id, name, number, email, password_hash, created_at
		FROM users
		WHERE id = $1
	`

	var user models.User

	err := db.QueryRowContext(ctx, query, userID).Scan(
		&user.ID,
		&user.Name,
		&user.Number,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(ctx context.Context, db *sql.DB, email string) (*models.User, error) {

	if email == "" {
		return nil, errors.New("email is required")
	}

	query := `
		SELECT id, name, number, email, password_hash, created_at
		FROM users
		WHERE email = $1
		LIMIT 1
	`

	var user models.User

	err := db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Number,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func GetUserByPhone(ctx context.Context, db *sql.DB, phone string) (*models.User, error) {

	if phone == "" {
		return nil, errors.New("phone is required")
	}

	query := `
		SELECT id, name, number, email, password_hash, created_at
		FROM users
		WHERE number = $1
		LIMIT 1
	`

	var user models.User

	err := db.QueryRowContext(ctx, query, phone).Scan(
		&user.ID,
		&user.Name,
		&user.Number,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}
