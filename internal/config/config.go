package config

import (
	"github.com/SenselessA/w2w_backend/internal/db"
	"github.com/spf13/viper"
)

type Config struct {
	DB     db.ConfigDB  `mapstructure:"db"`
	HTTP   HTTPConfig   `mapstructure:"http"`
	Secret SecretConfig `mapstructure:"secret"`
}

type SecretConfig struct {
	Jwt   string `mapstructure:"jwt"`
	Kodik string `mapstructure:"kodik"`
}

type HTTPConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

func Init(folder, filename string) (*Config, error) {
	var cfg Config

	viper.AddConfigPath(folder)
	viper.SetConfigName(filename)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		// log.Fatalln("ALO")
		return nil, err
	}

	//if err := envconfig.Process("db", &cfg.DB); err != nil {
	//	return nil, err
	//}

	return &cfg, nil
}
