package service

import (
	"Gl0ven/kata_projects/rates/internal/models"
	"Gl0ven/kata_projects/rates/internal/provider/garantex"
	"Gl0ven/kata_projects/rates/internal/storage"
	"context"

	"go.uber.org/zap"
)

type Service interface {
	GetRates(ctx context.Context) (models.Rates, error)
}

type ratesService struct {
	storage storage.Storage
	provider garantex.GarantexApi
	logger *zap.Logger
}

func NewService(strg storage.Storage, prv garantex.GarantexApi, lg *zap.Logger) Service {
	return &ratesService{
		storage: strg,
		provider: prv,
		logger: lg,
	}
}

func(s *ratesService) GetRates(ctx context.Context) (models.Rates, error) {
	rates, err := s.provider.GetRates()
	if err != nil {
		s.logger.Error("receiving rates from provider error", zap.Error(err))
		return models.Rates{}, err
	} 
	s.logger.Info("rates were successfully received from the provider", zap.Any("rates", rates))

	err = s.storage.SaveRates(ctx, rates)
	if err != nil {
		s.logger.Error("saving rates to the storage error", zap.Error(err))
		return models.Rates{}, err
	}

	s.logger.Info("rates were successfully saved to the storage")
	return rates, nil
}
