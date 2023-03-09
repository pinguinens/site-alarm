package service

import (
	"github.com/pinguinens/site-alarm/internal/api"
)

type Service struct {
	api *api.API
}

func New(api *api.API) (Service, error) {
	return Service{api: api}, nil
}

func (s *Service) Start() {

}
