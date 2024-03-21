package db

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	models2 "project/db/models"
)

var DB *gorm.DB

func Connect() {
	v := viper.GetViper()
	dbAddr := v.GetString("db_addr")
	connection, err := gorm.Open(postgres.Open(dbAddr), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection
	err = connection.AutoMigrate(&models2.User{}, &models2.Course{}, &models2.Feedback{})
	if err != nil {
		panic("pppp")
	}
}
