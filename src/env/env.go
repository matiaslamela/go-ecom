package enviroment

import (
	"os"

	"github.com/joho/godotenv"
)

// DB_HOST="localhost"
// DB_USER="postgres"
// DB_PASSWORD="1234"
// DB_NAME="go_prueba"
// DB_PORT="5432"
// PORT="3001"

var DB_HOST string
var DB_USER string
var DB_PASSWORD string
var DB_NAME string
var DB_PORT string
var PORT string

func GetEnvVars() {
	err := godotenv.Load(".env")

	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")
	PORT = os.Getenv("PORT")

	if err != nil {
		panic(err)
	}
}
