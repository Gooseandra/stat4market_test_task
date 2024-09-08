package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Service struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"service"`
	Db struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
	}
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./config"
	}
	configName := os.Getenv("CONFIG_NAME")
	if configName == "" {
		configName = "config"
	}

	v.AddConfigPath(configPath)
	v.SetConfigName(configName)
	v.SetConfigType("yml")

	err := v.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %s", err)
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
