package domain

import "time"

type PaymentsSummary struct {
	Default  Default
	Fallback Fallback
}

type Default struct {
	TotalRequests int
	TotalAmount   float32
}

type Fallback struct {
	TotalRequests int
	TotalAmount   float32
}

type NullTime struct {
	Time    *time.Time
	IsValid bool
}

// outgoing api calls

// request of POST /payments
type PaymentsRequest struct {
	CorrelationID string  `json:"correlationId"`
	Amount        float32 `json:"amount"`
	RequestesAt   string  `json:"requestedAt"` // timestamp in rfc3339 format
}

// response from POST /payments
type PaymentsResponse struct {
	Message string `json:"message"`
}

// response from GET /payments/service-health
type ServiceHealthResponse struct {
	Failing           bool `json:"failing"`
	MinResponseTimeMs int  `json:"minResponseTime"`
}

// general response
type HTTPResponse struct {
	Status int
	Body []byte
}