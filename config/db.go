package config

import (
	"fmt"
	"log"
	"os"

	"gym-membership/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Tidak bisa load .env, lanjut pakai environment system")
	}

	// Build DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Koneksi ke database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("❌ Gagal koneksi ke database:", err)
	}

	// Auto migrate semua model
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Membership{},
		&models.Trainer{},
		&models.WorkoutSession{},
	}

	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			log.Panicf("❌ Gagal migrasi model %T: %v", model, err)
		}
	}

	// Assign ke variabel global
	DB = db
	fmt.Println("✅ Database terkoneksi & migrasi semua model berhasil!")
}
