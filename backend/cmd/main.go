package main

import (
	"log"
	"net/http"

	"github.com/Indraaai/GolangReact/backend/internal/config"
	"github.com/Indraaai/GolangReact/backend/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Load Konfigurasi
	cfg := config.LoadConfig()

	// 2. Koneksi Database
	db := database.Connect(cfg)

	// 3. Setup Router Gin
	r := gin.Default()

	// 4. Route untuk Health Check (uji koneksi DB)
	r.GET("/health", func(c *gin.Context) {
		// Coba ambil koneksi SQL untuk test ping
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Gagal mengambil koneksi DB",
			})
			return
		}

		// Ping DB
		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Koneksi DB terputus",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Server sehat & Database terhubung!",
		})
	})

	// 5. Jalankan Server
	log.Printf("🚀 Server running on port %s", cfg.APPPort)
	r.Run(":" + cfg.APPPort)
}
