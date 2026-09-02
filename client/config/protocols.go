package config

import (
	"fmt"
	"os"
)

// defaultPayloadType is used when `type` is omitted; 0 (with message) must be explicit.
const defaultPayloadType uint8 = 1

type ProtocolConfig struct {
	ID          uint8  `toml:"id"`
	APIURL      string `toml:"api_url"`
	APIEndpoint string `toml:"api_endpoint"` // temporary to avoid a breaking change
	Type        *uint8 `toml:"type"`         // nil when omitted; 0 is a valid, distinct type
}

// PayloadType returns the submitSignatures payload type to send.
func (cfg ProtocolConfig) PayloadType() uint8 {
	if cfg.Type == nil {
		return defaultPayloadType
	}
	return *cfg.Type
}

// validate rejects a payload type the encoder cannot produce.
func (cfg ProtocolConfig) validate(name string) error {
	if t := cfg.PayloadType(); t != 0 && t != 1 {
		return fmt.Errorf("protocol.%s: type %d is not a submitSignatures payload type (0 or 1)", name, t)
	}
	return nil
}

func (cfg ProtocolConfig) XAPIKey() string {
	envVar := fmt.Sprintf("PROTOCOL_X_API_KEY_%d", cfg.ID)
	return os.Getenv(envVar)
}

// BaseURL returns the provider's base URL, preferring api_url.
func (cfg ProtocolConfig) BaseURL() string {
	if cfg.APIURL != "" {
		return cfg.APIURL
	}
	return cfg.APIEndpoint
}
