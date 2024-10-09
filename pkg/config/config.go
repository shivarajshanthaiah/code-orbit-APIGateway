package config

import "github.com/spf13/viper"

type Config struct {
	APIPORT   string `mapstructure:"APIPORT"`
	SECRETKEY string `mapstructure:"SECRETKEY"`
	USERPORT  string `mapstructure:"USERPORT"`
	ADMINPORT string `mapstructure:"ADMINPORT"`
	RAZORPAY  string `mapstructure:"RAZORPAY"`
}

func LoadConfig() (*Config, error) {
	var config Config

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}