package service

import "github.com/joseCarlosAndrade/rinha-backend-2025-jc/internal/core/domain/port"


type Service struct {
	storageRepository port.Storage
	apiRepository port.API
}

func NewService(
	storage port.Storage,
	api port.API) *Service {

	return &Service{
		storageRepository: storage,
		apiRepository: api,
	}
}
