package main

import (
	"fmt"

	"github.com/abhinandpn/GoTicks/backend/internal/config"
	"github.com/abhinandpn/GoTicks/backend/internal/db"
	"github.com/abhinandpn/GoTicks/backend/internal/demo"
	"github.com/abhinandpn/GoTicks/backend/internal/migration"
)

func main() {
	fmt.Println("Hello, World!")
	config.EnvLoad()
	db.SuperBaseAuth()
	migration.InitDB()
	demo.DemoCreateUser()
}
