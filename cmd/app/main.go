package main

import (
	"log"

	"minibank/config"
	"minibank/server"

	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app, err := server.NewApp()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	if err := app.Run(viper.GetString("PORT")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
