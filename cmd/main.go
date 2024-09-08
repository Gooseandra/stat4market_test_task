package main

import (
	"crypto/tls"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/go-openapi/loads"
	"github.com/rs/cors"
	"log"
	"net/http"
	"test_task_stat4market/config"
	"test_task_stat4market/internal/envents/delivery"
	"test_task_stat4market/internal/envents/repository"
	"test_task_stat4market/internal/envents/usecase"
	"test_task_stat4market/restapi"
	"test_task_stat4market/restapi/operations"
)

func main() {
	// Загрузка конфигурации
	cfgFile, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Config load error: ", err)
	}
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatal("Config parse error: ", err)
	}

	// Настройка соединения с базой данных
	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr:     []string{cfg.Db.Host},
		Protocol: clickhouse.Native,
		TLS:      &tls.Config{},
		Auth: clickhouse.Auth{
			Username: cfg.Db.User,
			Password: cfg.Db.Password,
		},
	})

	row := conn.QueryRow("SELECT 1")
	var col uint8
	if err := row.Scan(&col); err != nil {
		log.Fatal("Database connection error:", err)
	}
	log.Println("DB connected successfully")

	// Загрузка спецификации Swagger
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatal("Error loading swagger spec:", err)
	}

	eventRepo := repository.NewClickHouseRepository(conn)
	eventUseCase := usecase.NewUseCase(eventRepo)

	// Настройка API
	api := operations.NewEventsAPIAPI(swaggerSpec)
	api.GetEventByDayHandler = delivery.NewGetEventByDayHandler(eventUseCase)
	api.GetEventByValueHandler = delivery.NewEventByValueHandler(eventUseCase)
	api.GetEventHandler = delivery.NewGetEventHandler(eventUseCase)
	api.GetUserByValueHandler = delivery.NewGetUserByValueHandler(eventUseCase)
	api.NewEventHandler = delivery.NewCreateEventHandler(eventUseCase)

	// Создание и настройка сервера
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.ConfigureFlags()
	server.ConfigureAPI()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "JWT"},
		AllowCredentials: true,
	}).Handler(server.GetHandler())

	serverPort := cfg.Service.Port
	serverAddress := cfg.Service.Host
	log.Printf("Serving mchs at http://%s\n", serverAddress+serverPort)
	if err := http.ListenAndServe(serverAddress+serverPort, corsHandler); err != nil {
		log.Fatalln(err)
	}
}
