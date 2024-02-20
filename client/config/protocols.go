package config

type ProtocolConfig struct {
	Id          uint8  `toml:"id"`
	ApiEndpoint string `toml:"api_endpoint"`
	XApiKey     string `toml:"x_api_key"` // Value of the X-API-KEY header
}
