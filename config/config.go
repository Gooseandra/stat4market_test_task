package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Config struct {
	Service struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"service"`
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(os.Getenv("CONFIG_PATH"))
	v.SetConfigName(os.Getenv("CONFIG_NAME"))
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Println(err.Error())
		time.Sleep(time.Minute)
		log.Fatalf("unable to decode config into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
