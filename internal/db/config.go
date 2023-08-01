package db

import "fmt"

type ConfigDB struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

func (cfg *ConfigDB) String() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.User, cfg.Password, cfg.Name)
}
