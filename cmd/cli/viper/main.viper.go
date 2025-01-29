package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database []struct{
		User string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host string `mapstructure:"host"`
		DbName string `mapstructure:"dbName"`
	} `mapstructure:"database"`
	Security struct { 
		Jwt struct {
			Key string `mapstructure:"key"`
		} `mapstructure:"jwt"`
	} `mapstructure:"security"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/") // path to config
	viper.SetConfigName("local") // Tên file
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w \n ", err))
	}
	// read server configuration
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("JWT Port::", viper.GetString("security.jwt.key"))

	// Configure structure
	var config Config
	// Nếu mà bị lỗi
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode unmarshal configuration: %v", err)
	}

	// In ra cái Port của config
	fmt.Println("Config Port::", config.Server.Port)


	// Bởi vì thằng Database là một mảng nên là cần phải lặp qua nó
	for _, db := range config.Database {

		fmt.Printf(("database User: %s, Password: %s, host: %s"), db.User, db.Password, db.Host)
		fmt.Println("Database User::", db.User)
		fmt.Println("Database Password::", db.Password)
		fmt.Println("Database Host::", db.Host)
		fmt.Println("Database DbName::", db.DbName)
	}
} 