package config

import (
	"fmt"
	"os"
)

type ProtocolConfig struct {
	Id          uint8  `toml:"id"`
	ApiEndpoint string `toml:"api_endpoint"`
}

func (cfg ProtocolConfig) XApiKey() string {
	envVar := fmt.Sprintf("PROTOCOL_X_API_KEY_%d", cfg.Id)
	return os.Getenv(envVar)
}
