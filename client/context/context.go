package context

import (
	"flag"

	"github.com/flare-foundation/flare-system-client/client/config"
	globalConfig "github.com/flare-foundation/flare-system-client/config"

	"github.com/flare-foundation/go-flare-common/pkg/database"

	"gorm.io/gorm"
)

type ClientContext interface {
	Config() *config.Client
	DB() *gorm.DB
	Flags() *ClientFlags
}

type ClientFlags struct {
	ConfigFileName string
	// Add additional flags here
}

type clientContext struct {
	config *config.Client
	db     *gorm.DB
	flags  *ClientFlags
}

func BuildContext() (ClientContext, error) {
	flags := parseFlags()
	cfg, err := config.Build(flags.ConfigFileName)
	if err != nil {
		return nil, err
	}

	db, err := database.Connect(&cfg.DB)
	if err != nil {
		return nil, err
	}

	return &clientContext{
		config: cfg,
		db:     db,
		flags:  flags,
	}, nil
}

func (c *clientContext) Config() *config.Client { return c.config }

func (c *clientContext) DB() *gorm.DB { return c.db }

func (c *clientContext) Flags() *ClientFlags { return c.flags }

func parseFlags() *ClientFlags {
	cfgFlag := flag.String("config", globalConfig.ConfigFile, "Configuration file (toml format)")
	flag.Parse()

	return &ClientFlags{
		ConfigFileName: *cfgFlag,
	}
}
