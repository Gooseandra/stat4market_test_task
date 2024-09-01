package delivery

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"test_task_stat4market/internal/envents"
)

type Handlers struct {
	UseCase envents.UseCase
}

func NewHandlers(UC envents.UseCase) *Handlers {
	return &Handlers{
		UseCase: UC,
	}
}

func (h Handlers) GetEventTypesByEventValueHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		value := ctx.Params("value")
		res, err := h.UseCase.GetEventTypesByEventValue(value)
		if err != nil {
			log.Println(err.Error())
			return ctx.JSON(struct {
				Success bool   `json:"success"`
				Error   string `json:"error"`
			}{false, err.Error()})
		}
		return ctx.JSON(struct {
			Success bool                              `json:"success"`
			Types   []envents.EventTypesValueResponse `json:"types"`
			Error   *string                           `json:"error"`
		}{true, res, nil})
	}
}

func (h Handlers) GetEventByDayHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		day := ctx.Params("day")
		res, err := h.UseCase.GetEventByDay(day)
		if err != nil {
			log.Println(err.Error())
			return ctx.JSON(struct {
				Success bool   `json:"success"`
				Error   string `json:"error"`
			}{false, err.Error()})
		}
		return ctx.JSON(struct {
			Success bool            `json:"success"`
			Types   []envents.Event `json:"events"`
			Error   *string         `json:"error"`
		}{true, res, nil})
	}
}

func (h Handlers) GetUserByUniqueEventTypesValueHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		value := ctx.Params("value")
		res, err := h.UseCase.GetUserByUniqueEventTypesValue(value)
		if err != nil {
			log.Println(err.Error())
			return ctx.JSON(struct {
				Success bool   `json:"success"`
				Error   string `json:"error"`
			}{false, err.Error()})
		}
		return ctx.JSON(struct {
			Success bool                               `json:"success"`
			Types   []envents.UserByTypesValueResponse `json:"users_ids"`
			Error   *string                            `json:"error"`
		}{true, res, nil})
	}
}

func (h Handlers) NewEventHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params envents.Event
		if err := ctx.BodyParser(&params); err != nil {
			return err
		}
		err := h.UseCase.NewEvent(params)
		if err != nil {
			log.Println(err.Error())
			return ctx.JSON(struct {
				Success bool   `json:"success"`
				Error   string `json:"error"`
			}{false, err.Error()})
		}
		return ctx.JSON(struct {
			Success bool    `json:"success"`
			Error   *string `json:"error"`
		}{true, nil})
	}
}

func (h Handlers) GetEventByTypeAndDateRepo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		eventType := ctx.Query("type")
		date := ctx.Query("date")
		res, err := h.UseCase.GetEventByTypeAndDate(date, eventType)
		if err != nil {
			log.Println(err.Error())
			return ctx.JSON(struct {
				Success bool   `json:"success"`
				Error   string `json:"error"`
			}{false, err.Error()})
		}
		return ctx.JSON(struct {
			Success bool            `json:"success"`
			Types   []envents.Event `json:"events"`
			Error   *string         `json:"error"`
		}{true, res, nil})
	}
}
