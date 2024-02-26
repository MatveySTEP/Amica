package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"project/api"
	"project/db"
)

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal().Err(err).Msg("Не удалось запустить рутовую команду")
	}
	db.Connect()
	app := gin.Default()
	app.Use(cors.New(cors.Config{

		AllowCredentials: true,
	}))
	cnf := api.Config{Addr: "localhost:8100"}
	api.NewApi(cnf)

}
