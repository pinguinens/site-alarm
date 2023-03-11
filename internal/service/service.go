package service

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/rs/zerolog"
	zLog "github.com/rs/zerolog/log"
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

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	if updates == nil {
		updates = bot.GetUpdatesChan(u)
	}
	//go receiveUpdates(ctx, updates)

	handleUpdate(<-updates, &msg)

	cancel()

	return nil
}

func handleUpdate(update tgbotapi.Update, data *Msg) {
	handleMessage(update.Message, data)
}

func handleMessage(message *tgbotapi.Message, data *Msg) {
	user := message.From
	text := message.Text

	if user == nil {
		return
	}

	// Print to console
	zLog.Info().Msg(fmt.Sprintf("%s wrote %s", user.FirstName, text))

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("%v\n%v\n%v\n%v", data.Code, data.Method, data.URL, data.Address))
	_, err := bot.Send(msg)
	if err != nil {
		zLog.Info().Msg(fmt.Sprintf("%s wrote %s", user.FirstName, text))
		return
	}
}
