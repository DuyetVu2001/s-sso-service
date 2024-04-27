package util

import "github.com/spf13/viper"

type Config struct {
	POSTGRES_URL   string `mapStructure:"POSTGRES_URL"`
	SERVER_ADDRESS string `mapStructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)

	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
