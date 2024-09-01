package main

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
	"test_task_stat4market/config"
	"test_task_stat4market/internal/httpServer"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfgFile, err := config.LoadConfig()
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Println(err.Error())
		time.Sleep(time.Minute)
		log.Fatal("Config parse error: ", err.Error())
	}

	s := httpServer.NewServer(cfg)
	if err = s.Run(); err != nil {
		log.Println(err.Error())
	}
}
