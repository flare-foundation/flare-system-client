package config

import (
	"fmt"
	"os"
)

type ProtocolConfig struct {
	ID          uint8  `toml:"id"`
	ApiEndpoint string `toml:"api_endpoint"`
	Type        uint8  `toml:"type"`
}

func (cfg ProtocolConfig) XApiKey() string {
	envVar := fmt.Sprintf("PROTOCOL_X_API_KEY_%d", cfg.ID)
	return os.Getenv(envVar)
}
