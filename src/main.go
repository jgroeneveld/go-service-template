package main

import (
	"fmt"
	"github.com/jgroeneveld/go-cloud-run-template/src/app"
	"os"
)

func main() {
	port := getEnvOrDefault("PORT", "8080")

	fmt.Printf("Starting server on port %s\n", port)

	if err := app.StartServer(port); err != nil {
		panic(err)
	}
}

func getEnvOrDefault(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}
	return value
}
