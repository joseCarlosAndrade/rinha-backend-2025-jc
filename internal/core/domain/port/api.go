package port

import (
	"context"

	"github.com/joseCarlosAndrade/rinha-backend-2025-jc/internal/core/domain"
)



type API interface {
	Do(ctx context.Context, method, url string, queryParams map[string]string, body []byte) (domain.HTTPResponse, error)
}