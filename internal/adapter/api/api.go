package api

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/joseCarlosAndrade/rinha-backend-2025-jc/internal/core/domain"
	"go.uber.org/zap"
)

func (r* Repository) Do(ctx context.Context, method, url string, queryParams map[string]string, body []byte) (*domain.HTTPResponse, error) {
	bodyReader := bytes.NewBuffer(body)
	
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		zap.L().Error("could not create get request with context", zap.Error(err))
		return nil, err
	}

	// assembling url query params
	q := req.URL.Query()

	for key, value := range queryParams {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		zap.L().Error("could not execute http get request", zap.String("url", req.URL.String()), zap.Error(err))
		return nil, err
	}

	// getting response body
	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		zap.L().Error("could not read response body", zap.Error(err))
		return nil, err
	}

	apiReponse := domain.HTTPResponse{
		Status: res.StatusCode,
		Body: resBytes,
	}

	return &apiReponse, nil
}
