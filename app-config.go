package main

import (
	_ "embed"
	"fmt"
	configuration "login-backend/configuration"
	"os"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

var DefaultConfig = Config{
	Log: LogConfig{
		Level:      -1,
		EnableJSON: false,
	},
	App: AppConfig{
		GinRouter: configuration.RouterConfig{
			Router: "0.0.0.0:8085",
		},
		ServiceName: "login--backend",
		Auth: configuration.Param{
			TokenUrl:     "http://localhost:8443/realms/my-realm/protocol/openid-connect/token?",
			ClientID:     "my-client",
			ClientSecret: "TbcVdCDnuu2krqgN8yv3tGdrACIfaWT0",
		},
	},
}

type Config struct {
	Log LogConfig `yaml:"log"`
	App AppConfig `yaml:"config"`
}

type LogConfig struct {
	Level      int  `yaml:"level"`
	EnableJSON bool `yaml:"enablejson"`
}

type AppConfig struct {
	GinRouter   configuration.RouterConfig `yaml:"ginrouter" mapstructure:"ginrouter" json:"ginrouter"`
	ServiceName string                     `yaml:"service-name" mapstructure:"service-name" json:"service-name"`
	Auth        configuration.Param        `yaml:"auth" mapstructure:"auth" json:"auth"`
}

// Default config file.
//
//go:embed configuration/config.yaml
var projectConfigFile []byte

const ConfigFileEnvVar = "LOGIN_BACKEND_FILE_PATH"
const ConfigurationName = "login-backend"

func ReadConfig() (*Config, error) {

	configPath := os.Getenv(ConfigFileEnvVar)
	var cfgContent []byte
	var err error
	if configPath != "" {
		if _, err := os.Stat(configPath); err == nil {
			log.Info().Str("cfg-file-name", configPath).Msg("reading config")
			cfgContent, err = util.ReadFileAndResolveEnvVars(configPath)
			log.Info().Msg("++++CFG:" + string(cfgContent))
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("the %s env variable has been set but no file cannot be found at %s", ConfigFileEnvVar, configPath)
		}
	} else {
		log.Warn().Msgf("The config path variable %s has not been set. Reverting to bundled configuration", ConfigFileEnvVar)
		cfgContent = util.ResolveConfigValueToByteArray(projectConfigFile)
		return nil, fmt.Errorf("the config path variable %s has not been set; please set", ConfigFileEnvVar)
	}

	appCfg := DefaultConfig
	err = yaml.Unmarshal(cfgContent, &appCfg)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	if !appCfg.Log.EnableJSON {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	zerolog.SetGlobalLevel(zerolog.Level(appCfg.Log.Level))

	return &appCfg, nil
}
