package main

import (
	"flag"

	"github.com/rs/zerolog/log"

	"github.com/pinguinens/site-alarm/internal/api"
	"github.com/pinguinens/site-alarm/internal/config"
	"github.com/pinguinens/site-alarm/internal/service"
)

const (
	appVersion = "0.1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "", "Custom config path")
	flag.Parse()
}

func main() {
	appConfig, err := config.New(configPath)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	svc, err := service.New(&log.Logger, appVersion, appConfig.Listen.Address)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	_, err = api.New(svc)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
