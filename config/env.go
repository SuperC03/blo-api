package config

import (
	"net/url"

	"github.com/caarlos0/env/v11"
)

type EnvConfig struct {
	Hostname   string `env:"HOSTNAME" envDefault:"localhost"`
	Port       int    `env:"PORT" envDefault:"8080"`
	Production bool   `env:"PRODUCTION" envDefault:"0"`

	PostgresUser     string `env:"POSTGRES_USERNAME,required"`
	PostgresPassword string `env:"POSTGRES_PASSWORD,required"`
	PostgresName     string `env:"POSTGRES_NAME,required"`
	PostgresHost     string `env:"POSTGRES_HOST,required"`
	PostgresPort     int    `env:"POSTGRES_PORT,required"`

	GotifyEnabled bool    `env:"GOTIFY_ENABLED" envDefault:"0"`
	GotifyURL     url.URL `env:"GOTIFY_URL"`
	GotifyToken   string  `env:"GotifyToken"`

	GoogleKeyPath string `env:"GOOGLE_KEY_PATH,required"`

	SwaggerHost     string `env:"SWAGGER_HOST,required"`
	SwaggerBasePath string `env:"SWAGGER_BASE_PATH,required"`

	BananalogDocID             string `env:"BANANALOG_DOC_ID,required"`
	BananalogOverviewDataRange string `env:"BANANALOG_OVERVIEW_DATA_RANGE,required"`
	// Need to figure out how to do this with Env variables...eventually
	// BananalogEnabledOffset     int    `env:"BANANALOG_ENABLED_OFFSET,required"`
}

func ParseEnv() (EnvConfig, error) {
	return env.ParseAs[EnvConfig]()
}
