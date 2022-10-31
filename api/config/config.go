package config

import (
	"github.com/spf13/viper"
	"strings"
)

const (
	configPath    = "../../api/config"
	configName    = "config"
	configPathENV = "FX_TEST_CFG_PATH"
	configNameENV = "FX_TEST_CFG_NAME"
)

type APIConfig struct {
	Host  string `mapstructure:"host"`
	Route struct {
		Echo  string `mapstructure:"echo"`
		Hello string `mapstructure:"hello"`
	} `mapstructure:"route"`
}

func InitializeAPIConfig() (*APIConfig, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := new(APIConfig)
	err = viper.Unmarshal(c)

	return c, err
}

func InitializeAPIConfigFromENV() (*APIConfig, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigName(configNameENV)
	viper.AddConfigPath(configPathENV)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := new(APIConfig)
	err = viper.Unmarshal(c)

	return c, err
}
