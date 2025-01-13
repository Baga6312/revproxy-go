package configs

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type resource struct {
	Name            string
	Endpoint        string
	Destination_url string
}

type configuratrions struct {
	server struct {
		Host        string
		Listen_Port string
	}
	Resources []resource
}

var configs *configuratrions

func NewConfigs() (*configuratrions, error) {
	viper.addConfigPath("data")
	viper.setConfigName("config")
	viper.setConfigType("yaml")
	viper.automaticEnv()
	viper.setEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Error loading files %s", err)
	}

	err = viper.Unmarshal(&configs)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %s", err)
	}
	return configs, nil
}
