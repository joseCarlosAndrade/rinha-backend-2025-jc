package config

type AppInfo struct {
	PaymentProcessorAddress         string `default:"http://localhost:8080"`
	PaymentProcessorFallbackAddress string `default:"http://localhost:8081"`
	Debug                           bool   `default:"true"`
	Port                            string `default:"8080"`
}

var (
	App AppInfo
)
