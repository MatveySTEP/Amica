package db

import (
	models2 "amica/db/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		panic("ошибка миграции")
	}
}
