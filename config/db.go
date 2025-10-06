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

	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Tidak bisa load .env, lanjut pakai environment system")
	}

	databaseURL := os.Getenv("DATABASE_URL")

	var dsn string
	if databaseURL != "" {

		dsn = databaseURL
		log.Println("🌐 Menggunakan DATABASE_URL dari Railway")
	} else {

		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
		log.Println("💻 Menggunakan konfigurasi lokal dari .env")
	}

	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("❌ Gagal koneksi ke database:", err)
	}

	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Membership{},
		&models.Trainer{},
		&models.WorkoutSession{},
	}

	if err := db.AutoMigrate(modelsToMigrate...); err != nil {
		log.Panicf("❌ Gagal migrasi model: %v", err)
	}

	DB = db
	log.Println("✅ Database terkoneksi & semua tabel berhasil dimigrasi!")
}
