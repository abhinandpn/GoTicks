package main

import (
	"fmt"

	"github.com/abhinandpn/GoTicks/backend/internal/config"
	"github.com/abhinandpn/GoTicks/backend/internal/db"
)

func main() {
	fmt.Println("Hello, World!")
	config.EnvLoad()
	db.SuperBaseAuth()
}
