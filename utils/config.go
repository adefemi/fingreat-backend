package utils

import "github.com/spf13/viper"

type Config struct {
	DBdriver    string `mapstructure:"DB_DRIVER"`
	DB_name     string `mapstructure:"DB_NAME"`
	DB_username string `mapstructure:"DB_USERNAME"`
	DB_password string `mapstructure:"DB_PASSWORD"`
	DB_host     string `mapstructure:"DB_HOST"`
	DB_port     string `mapstructure:"DB_PORT"`
	Signing_key string `mapstructure:"SIGNING_KEY"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
