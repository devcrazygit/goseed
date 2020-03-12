package utils

import (
	"os"

	"github.com/joho/godotenv"
)

//EnvVar function is for read .env file
func EnvVar(key string) string {
	// fmt.Println("EnvVar called")
	godotenv.Load(".env")
	return os.Getenv(key)
}
