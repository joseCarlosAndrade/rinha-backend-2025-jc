package config

type AppInfo struct {
	PaymentProcessorAddress         string `default:"http://localhost:8080"`
	PaymentProcessorFallbackAddress string `default:"http://localhost:8081"`
	Debug                           bool   `default:"true"`
}

var (
	App AppInfo
)
