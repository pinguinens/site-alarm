package main

import (
	"github.com/rs/zerolog/log"

	"github.com/pinguinens/site-alarm/internal/api"
	"github.com/pinguinens/site-alarm/internal/service"
)

const (
	appVersion = "0.1"
)

func main() {
	appApi, err := api.New()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	svc, err := service.New(appApi)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	svc.Start()
}
