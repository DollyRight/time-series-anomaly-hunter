package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"server/config"
)

func Config() config.Configs {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error read config file: %w", err))
	}
	var config config.Configs
	err = v.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("fatal error unmarshal config file: %w", err))
	}
	return config
}
