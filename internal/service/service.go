package service

import (
	log "github.com/rs/zerolog"
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
