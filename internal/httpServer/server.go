package httpServer

import (
	"fmt"
	gojson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
	"log"
	"test_task_stat4market/config"
)

type Server struct {
	fiber *fiber.App
	cfg   *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		fiber: fiber.New(fiber.Config{
			DisableStartupMessage: true,
			JSONEncoder:           gojson.Marshal,
			JSONDecoder:           gojson.Unmarshal,
		}),
		cfg: cfg,
	}
}

func (s *Server) Run() error {
	if err := s.MapHandlers(s.fiber); err != nil {
		log.Println("Cannot map handlers: %s", err)
	}
	log.Println("Starting server on " + s.cfg.Service.Host + ":" + s.cfg.Service.Port)
	if err := s.fiber.Listen(fmt.Sprintf("%s:%s", s.cfg.Service.Host, s.cfg.Service.Port)); err != nil { // надо было в конфиг
		log.Fatalf("Error starting Server: %s", err)
	}

	return nil
}
