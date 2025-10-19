package config

type HTTPConfig struct {
	Port string
}

func NewHTTPConfig() *HTTPConfig {
	return &HTTPConfig{
		Port: "8080",
	}
}
