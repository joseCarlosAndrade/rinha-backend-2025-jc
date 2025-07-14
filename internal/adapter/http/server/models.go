package server

type PaymentsRequest struct {
	CorrelationID string  `json:"correlationId"`
	Amount        float32 `json:"amount"`
}

type PaymentsSummaryResponse struct {
	Default  Default  `json:"default"`
	Fallback Fallback `json:"fallback"`
}

type Default struct {
	TotalRequests int     `json:"totalRequests"`
	TotalAmount   float32 `json:"totalAmount"`
}

type Fallback struct {
	TotalRequests int     `json:"totalRequests"`
	TotalAmount   float32 `json:"totalAmount"`
}
