package port

import (
	"context"

	"github.com/joseCarlosAndrade/rinha-backend-2025-jc/internal/core/domain"
)

type Service interface {
	GeneratePayment(ctx context.Context, correlationId string, amount float32) error
	GeneratePaymentsSummary(ctx context.Context, from, to string) (domain.PaymentsSummary, error)
}