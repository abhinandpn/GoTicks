package migration

import (
	"log"
	"os"

	"github.com/abhinandpn/GoTicks/backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {

	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	// check connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("DB not reachable:", err)
	}

	DB = db

	// enable UUID
	DB.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto";`)

	// ✅ run migration only when needed
	if os.Getenv("RUN_MIGRATION") == "true" {

		log.Println("Running migrations...")

		err = DB.AutoMigrate(
			&models.User{},
			&models.Theater{},
			&models.Screen{},
			&models.Movie{},
			&models.Show{},
			&models.Ticket{},
			&models.Booking{},
			&models.BookingSeat{},
			&models.Payment{},
		)

		if err != nil {
			log.Fatal("Migration failed:", err)
		}

		log.Println("Migration completed")
	}

	log.Println("Database migrated successfully")
}
