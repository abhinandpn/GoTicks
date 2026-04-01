package demo

import (
	"context"
	"fmt"
	"log"

	"github.com/abhinandpn/GoTicks/backend/internal/migration"
	"github.com/abhinandpn/GoTicks/backend/internal/models"
	"github.com/abhinandpn/GoTicks/backend/internal/repository"
	"github.com/redis/go-redis/v9"
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

func CheckRedis() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	res, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("❌ Redis not working:", err)
		return
	}

	fmt.Println("✅ Redis working:", res) // should print "PONG"
}
