package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	DBUrl        string `mapstructure:"DB_URL"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
        return config, fmt.Errorf("failed to read config: %w", err)
    }
	if err = viper.Unmarshal(&config); err != nil {
        return config, fmt.Errorf("failed to unmarshal config: %w", err)
    }
	return config,nil
}
