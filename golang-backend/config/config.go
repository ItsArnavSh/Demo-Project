package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string `mapstructure:"server_address"`

	DBHost     string `mapstructure:"host"`
	DBPort     int    `mapstructure:"port"`
	DBUser     string `mapstructure:"user"`
	DBPassword string `mapstructure:"password"`
	DBName     string `mapstructure:"dbname"`
	DBSSLMode  string `mapstructure:"sslmode"`
}

func LoadConfig(path string) (Config, error) {
	var config Config

	viper.SetConfigName("config") // name of file (without extension)
	viper.SetConfigType("yaml")   // or viper.AutomaticEnv()
	viper.AddConfigPath(path)     // look for config in the path

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	// Nested configs (like server.address) can be flattened automatically
	config.ServerAddress = viper.GetString("server.address")

	config.DBHost = viper.GetString("database.host")
	config.DBPort = viper.GetInt("database.port")
	config.DBUser = viper.GetString("database.user")
	config.DBPassword = viper.GetString("database.password")
	config.DBName = viper.GetString("database.dbname")
	config.DBSSLMode = viper.GetString("database.sslmode")

	return config, nil
}
