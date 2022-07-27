package main

import (
	"os"

	"example.com/m/src/controllers"
	"github.com/joho/godotenv"
)

const (
	envKey = "ENVIRONMENT"
)

func main() {
	env := os.Getenv(envKey)
	if env == "dev" {
		godotenv.Load()
	}

	controllers.App()
}
