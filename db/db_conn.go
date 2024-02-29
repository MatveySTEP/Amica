package db

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"project/models"
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
	err = connection.AutoMigrate(&models.User{})
	if err != nil {
		panic("pppp")
	}
}
