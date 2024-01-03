package config

type ProtocolConfig struct {
	Id          uint8  `toml:"id"`
	ApiEndpoint string `toml:"api_enpoint"`
}
