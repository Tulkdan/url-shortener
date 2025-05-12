package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Tulkdan/url-shortener/internal"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	connValkey := fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "6379"))
	db, err := internal.NewValkey(connValkey)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Client.Close()

	port := getEnv("PORT", "8000")
	if err := internal.NewHttpServer(":"+port, db).ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
