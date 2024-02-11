package utilities

import (
	"log"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading env file %s", err)
	}
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hash), err
}
