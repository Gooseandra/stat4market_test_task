package httpServer

import (
	"crypto/tls"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"test_task_stat4market/internal/envents/delivery"
	"test_task_stat4market/internal/envents/repository"
	"test_task_stat4market/internal/envents/usecase"
)

func (s *Server) MapHandlers(app *fiber.App) error {
	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr:     []string{os.Getenv("DB_HOST")},
		Protocol: clickhouse.Native,
		TLS:      &tls.Config{},
		Auth: clickhouse.Auth{
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
		},
	})
	row := conn.QueryRow("SELECT 1")
	var col uint8
	if err := row.Scan(&col); err != nil {
		fmt.Printf("An error while reading the data: %s", err)
	} else {
		log.Println("DB connected seccessfully")
	}

	EventRepo := repository.NewClickHouseRepository(conn)
	EventUseCase := usecase.NewUseCase(EventRepo)
	EventHandlers := delivery.NewHandlers(EventUseCase)

	delivery.MapRoutes(app, EventHandlers)

	return nil
}
