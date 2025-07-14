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
