package config

import (
	"fmt"
	"os"
	"time"

	"backend.com/go-backend/app/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file!")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Gorm dafabase config
	config := &gorm.Config{
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		NowFunc: func() time.Time {
			return time.Now().UTC() // for consistent timezone
		},
	}

	database, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        dsn,
	}), config)
	if err != nil {
		panic("Failed to connect to database!")
	}

	sqlDB, err := database.DB()
	if err != nil {
		panic(fmt.Sprintf("Failed to get DB instance: %v", err))
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Migrations with transactions
	err = database.Transaction(func(tx *gorm.DB) error {
		models := []interface{}{
			&models.User{},
			&models.Listing{},
			&models.Realtor{},
		}

		fmt.Println("Migrating models...")
		for _, model := range models {
			fmt.Printf("Migrating model: %T\n", model)
		}
		return tx.AutoMigrate(models...)
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to migrate: %v", err))
	}

	// Verify Migration
	if !database.Migrator().HasTable(&models.Listing{}) {
		panic("Listing table does not exist!")
	}

	fmt.Println("Database connected and migrated successfully!")
	DB = database // Assign the database connection to the global variable
}

func SessionStorage() redis.Store {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file!")
	}
	store, _ := redis.NewStore(10, "tcp", os.Getenv("REDIS_URL"), "", []byte(os.Getenv("SESSION_SECRET")))

	store.Options(sessions.Options{
		Path:     "/api/v1",
		MaxAge:   3 * 24 * 60 * 60,
		HttpOnly: true,
		Secure:   true,
		SameSite: 3,
	})
	return store
}
