package service

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/rs/zerolog"
)

var (
	screaming = false
	bot       *tgbotapi.BotAPI
	updates   tgbotapi.UpdatesChannel
)

type Service struct {
	version string
	addr    string
	Logger  *log.Logger
}

func New(logger *log.Logger, version, addr string) (*Service, error) {
	return &Service{
		version: version,
		addr:    addr,
		Logger:  logger,
	}, nil
}

func (s *Service) GetAddr() string {
	return s.addr
}

func (s *Service) Log(data []byte) error {
	msg := Msg{}
	err := Decode(data, &msg)
	if err != nil {
		return err
	}

	s.Logger.Info().Int("code", msg.Code).Str("method", msg.Method).Str("url", msg.URL).Str("addr", msg.Address).Send()

	return nil
}

func (s *Service) Notif(data []byte) error {
	msg := Msg{}
	err := Decode(data, &msg)
	if err != nil {
		return err
	}

	if bot == nil {
		bot, err = tgbotapi.NewBotAPI("")
		if err != nil {
			return err
		}
	}

	bot.Debug = false

	msgTg := tgbotapi.NewMessage(123, fmt.Sprintf("%v\n%v\n%v\n%v", msg.Code, msg.Method, msg.URL, msg.Address))
	_, err = bot.Send(msgTg)
	if err != nil {
		return err
	}

	return nil
}
