package service

import (
	"fmt"

	log "github.com/rs/zerolog"

	"github.com/pinguinens/site-alarm/internal/messenger"
)

type Service struct {
	version   string
	addr      string
	Logger    *log.Logger
	Messenger *messenger.Messenger
}

func New(logger *log.Logger, msgr *messenger.Messenger, version, addr string) (*Service, error) {
	return &Service{
		version:   version,
		addr:      addr,
		Logger:    logger,
		Messenger: msgr,
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

	err = s.Messenger.Send(fmt.Sprintf("ℹ️Notification\n<b>Code</b>: %v\n<b>Method</b>: %v\n<b>URL</b>: %v\n<b>Address</b>: %v", msg.Code, msg.Method, msg.URL, msg.Address))

	return nil
}
