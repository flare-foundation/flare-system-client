package config

import (
	"fmt"
	"os"
)

type ProtocolConfig struct {
	ID          uint8  `toml:"id"`
	APIUrl      string `toml:"api_url"`
	APIEndpoint string `toml:"api_endpoint"` // temporary to avoid a breaking change
	Type        uint8  `toml:"type"`
}

func (cfg ProtocolConfig) XApiKey() string {
	envVar := fmt.Sprintf("PROTOCOL_X_API_KEY_%d", cfg.ID)
	return os.Getenv(envVar)
}
