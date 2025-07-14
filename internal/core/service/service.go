package service

import (
	"context"
	"time"

	"github.com/joseCarlosAndrade/rinha-backend-2025-jc/internal/core/domain"
	"github.com/joseCarlosAndrade/rinha-backend-2025-jc/internal/core/domain/port"
	"go.uber.org/zap"
)


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

func (s* Service) GeneratePayment(ctx context.Context, correlationId string, amount float32) error {

	return nil
}

func (s* Service) GeneratePaymentsSummary(ctx context.Context, from, to string) (domain.PaymentsSummary, error) {

	fromTime, toTime, err := handleTimePeriod(from, to)
	if err != nil {
		return domain.PaymentsSummary{}, err
	}

	// implement redis check

	return domain.PaymentsSummary{}, nil
}

// takes a from and to string, tries to convert it to time and returns a nulltime type,
// which facilitates when querying this parameters
func handleTimePeriod( from, to string) (domain.NullTime, domain.NullTime, error) {
	var fromTime domain.NullTime

	if from != "" {
		time, err := handleTimeString(from)
		if err != nil {
			return domain.NullTime{}, domain.NullTime{}, err
		}

		fromTime.Time = &time
		fromTime.IsValid = true
	}

	var toTime domain.NullTime

	if to != "" {
		time, err := handleTimeString(to)
		if err != nil {
			return domain.NullTime{}, domain.NullTime{}, err
		}

		toTime.Time = &time
		toTime.IsValid = true
	}

	return fromTime, toTime, nil
}


// format the time string param from RFC3339 2020-07-10T12:34:56.000Z to time.Time
func handleTimeString(param string) (time.Time, error) {
	time, err := time.Parse(time.RFC3339, param)

	if err != nil {
		zap.L().Error("could not format time", zap.String("string", param), zap.Error(err))
		return time, err
	}

	zap.L().Debug("formated time", zap.Time(param, time) )
	return time, nil
}