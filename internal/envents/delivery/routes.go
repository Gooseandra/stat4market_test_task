package delivery

import "github.com/gofiber/fiber/v2"

func MapRoutes(router fiber.Router, h *Handlers) {
	router.Get("/events/value/:value", h.GetEventTypesByEventValueHandler())
	router.Get("/events/day/:day", h.GetEventByDayHandler())
	router.Get("/user/:value", h.GetUserByUniqueEventTypesValueHandler())
	router.Get("/event", h.GetEventByTypeAndDateRepo())
	router.Post("/event/new", h.NewEventHandler())
}
