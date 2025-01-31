package initialize

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

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/") // path to config
	viper.SetConfigName("local")     // Tên file
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
	// Sẽ khai báo global để mà truy cập toàn thư mục để mà lấy được cấu hình -> Thì chúng ta nên set cho nó một cái biến global
	var config Config
	// Nếu mà bị lỗi
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode unmarshal configuration: %v", err)
	}
}