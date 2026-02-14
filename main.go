package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"fastcicd/api/bot"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	ServerPort string
}

func loadConfig() Config {
	return Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "fastcicd"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	godotenv.Load()

	cfg := loadConfig()
	log.Printf("Starting server with config: %+v\n", cfg)

	mux := http.NewServeMux()

	mux.HandleFunc("/telegramWebhook", bot.TelegramWebhookHandler)

	// отдаём frontend
	mux.Handle("/", http.FileServer(http.Dir("./web")))

	serverAddr := ":" + cfg.ServerPort
	log.Printf("Server starting on %s", serverAddr)
	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		log.Fatal(err)
	}
}
