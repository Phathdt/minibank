package main

import (
	"log"

	"minibank/config"
	"minibank/server"

	"github.com/spf13/viper"
)

func main() {
	config.Init()

	app, err := server.NewApp()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	if err := app.Run(viper.GetString("PORT")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
