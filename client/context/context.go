package context

import (
	"flag"
	"flare-tlc/client/config"
	globalConfig "flare-tlc/config"
	"flare-tlc/database"

	"gorm.io/gorm"
)

type ClientContext interface {
	Config() *config.ClientConfig
	DB() *gorm.DB
	Flags() *ClientFlags
}

type ClientFlags struct {
	ConfigFileName string
	// Add additional flags here
}

type clientContext struct {
	config *config.ClientConfig
	db     *gorm.DB
	flags  *ClientFlags
}

func BuildContext() (ClientContext, error) {
	flags := parseFlags()
	cfg, err := config.BuildConfig(flags.ConfigFileName)
	if err != nil {
		return nil, err
	}
	globalConfig.GlobalConfigCallback.Call(cfg)

	db, err := database.ConnectAndInitialize(&cfg.DB)
	if err != nil {
		return nil, err
	}

	return &clientContext{
		config: cfg,
		db:     db,
		flags:  flags,
	}, nil
}

func (c *clientContext) Config() *config.ClientConfig { return c.config }

func (c *clientContext) DB() *gorm.DB { return c.db }

func (c *clientContext) Flags() *ClientFlags { return c.flags }

func parseFlags() *ClientFlags {
	cfgFlag := flag.String("config", globalConfig.CONFIG_FILE, "Configuration file (toml format)")
	flag.Parse()

	return &ClientFlags{
		ConfigFileName: *cfgFlag,
	}
}
