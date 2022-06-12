package models

import (
	"fmt"

	enviroment "github.com/matiaslamela/go-ecom/src/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func queryStringBuilder() string {
	return "host=" + enviroment.DB_HOST + " user=" + enviroment.DB_USER + " password=" + enviroment.DB_PASSWORD + " dbname=" + enviroment.DB_NAME + " port=" + enviroment.DB_PORT + " sslmode=disable"
}

func Connect() {
	dsn := queryStringBuilder()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	migrateError := db.AutoMigrate(&User{})
	if migrateError != nil {
		fmt.Println(err.Error())
		panic("error in migration")
	}
	DB = db

}
