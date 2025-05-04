package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	AppName   string `mapstructure:"appname"`
	Port      string `mapstructure:"port"`
	JWTSecret string `mapstructure:"jwt_secret"`

	DB struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
		User string `mapstructure:"user"`
		Pass string `mapstructure:"pass"`
		Name string `mapstructure:"name"`
	} `mapstructure:"db"`

	Redis struct {
		Addr string `mapstructure:"addr"`
	} `mapstructure:"redis"`

	RabbitMQ struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"rabbitmq"`

	S3 struct {
		Region    string `mapstructure:"region"`
		Bucket    string `mapstructure:"bucket"`
		AccessKey string `mapstructure:"access_key"`
		SecretKey string `mapstructure:"secret_key"`
	} `mapstructure:"s3"`
}

var AppConfig Config

func LoadConfig(path string) {
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Unmarshal config error: %v", err)
	}

	log.Println("✅ Loaded config from", path)
}

// package config

// import (
//     "log"

//     "github.com/spf13/viper"
// )

// type Config struct {
//     AppName string `mapstructure:"appname"`
//     Port    string `mapstructure:"port"`
//     JWTSecret string `mapstructure:"JWT_SECRET"`

//     DB struct {
//         Host string `mapstructure:"host"`
//         Port int    `mapstructure:"port"`
//         User string `mapstructure:"user"`
//         Pass string `mapstructure:"pass"`
//         Name string `mapstructure:"name"`
//     } `mapstructure:"db"`

//     Redis struct {
//         Addr string `mapstructure:"addr"`
//     } `mapstructure:"redis"`

//     RabbitMQ struct {
//         URL string `mapstructure:"url"`
//     } `mapstructure:"rabbitmq"`

//     S3 struct {
//         Region     string `mapstructure:"region"`
//         Bucket     string `mapstructure:"bucket"`
//         AccessKey  string `mapstructure:"access_key"`
//         SecretKey  string `mapstructure:"secret_key"`
//     } `mapstructure:"s3"`
// }

// var AppConfig Config

// func LoadConfig(path string) {
//     viper.SetConfigFile(path)

//     err := viper.ReadInConfig()
//     if err != nil {
//         log.Fatalf("Config error: %v", err)
//     }

//     err = viper.Unmarshal(&AppConfig)
//     if err != nil {
//         log.Fatalf("Unmarshal config error: %v", err)
//     }

//     log.Println("✅ Loaded config from", path)
// }
