package config

import (
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func RandomString(n int) string {

	rand.Seed(time.Now().UnixNano())

	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func EnvFunc(env string) string {

	errenv := godotenv.Load()
	if errenv != nil {
		panic("Error loading .env file")
	}

	env = os.Getenv(env)

	return string(env)
}
