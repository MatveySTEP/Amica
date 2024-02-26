package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"project/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection
	connection.AutoMigrate(&models.User{})
}
