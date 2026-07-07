package database

import (
	"fmt"
	"log"

	"github.com/Indraaai/GolangReact/backend/internal/config" // Sesuaikan jika module name berbeda

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) *gorm.DB {
	// Format DSN (Data Source Name) untuk MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	// Buka koneksi
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Gagal koneksi ke database: %v", err)
	}

	// Ambil koneksi SQL raw-nya untuk test ping
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Gagal mengambil koneksi SQL: %v", err)
	}

	// Ping database untuk memastikan benar-benar nyambung
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("❌ Database tidak merespon (ping failed): %v", err)
	}

	log.Println("✅ Berhasil terhubung ke database MySQL!")
	return db
}
