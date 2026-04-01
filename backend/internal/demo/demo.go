package demo

import (
	"context"
	"fmt"
	"log"

	"github.com/abhinandpn/GoTicks/backend/internal/migration"
	"github.com/abhinandpn/GoTicks/backend/internal/models"
	"github.com/abhinandpn/GoTicks/backend/internal/repository"
)

func DemoCreateUser() {

	sqlDB, err := migration.DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	NewUser := models.User{
		Name:         "John Doe",
		Number:       "1234567890",
		Email:        "test@gmail.com",
		PasswordHash: "hashedpassword",
	}

	err = repository.CreateUser(ctx, sqlDB, NewUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User created successfully data:", NewUser)
}
