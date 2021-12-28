package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database  *DatabaseConfig
	Port      string
	JWTSecret string
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
	SSLMode  string
	Migrate  bool
}

func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", d.User, d.Password, d.Host, d.Port, d.Name)
}

func NewConfig() *Config {
	return &Config{
		Database: &DatabaseConfig{
			Host:     viper.GetString("database_host"),
			User:     viper.GetString("database_user"),
			Password: viper.GetString("database_password"),
			Name:     viper.GetString("database_name"),
			Port:     viper.GetString("database_port"),
			SSLMode:  viper.GetString("database_ssl"),
			Migrate:  viper.GetBool("with_migration"),
		},
		Port:      viper.GetString("port"),
		JWTSecret: viper.GetString("jwt_secret"),
	}
}
