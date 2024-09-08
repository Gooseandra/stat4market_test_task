package delivery

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	"log"
	"test_task_stat4market/internal/envents"
	"test_task_stat4market/models"
	"test_task_stat4market/restapi/operations"
)

type (
	event struct {
		useCase envents.UseCase
	}
	GetEventByValue struct{ event }
	GetEventByDay   struct{ event }
	GetUserByValue  struct{ event }
	GetEvent        struct{ event }
	NewEvent        struct{ event }
)

// Функции для создания объектов функций Handle

func NewEventByValueHandler(UC envents.UseCase) GetEventByValue {
	return GetEventByValue{event{useCase: UC}}
}

func NewGetEventByDayHandler(UC envents.UseCase) GetEventByDay {
	return GetEventByDay{event{useCase: UC}}
}

func NewGetUserByValueHandler(UC envents.UseCase) GetUserByValue {
	return GetUserByValue{event{useCase: UC}}
}

func NewGetEventHandler(UC envents.UseCase) GetEvent {
	return GetEvent{event{useCase: UC}}
}

func NewCreateEventHandler(UC envents.UseCase) NewEvent {
	return NewEvent{event{useCase: UC}}
}

// Handle обрабатывает запрос GET /events/value/{value}.
func (h GetEventByValue) Handle(params operations.GetEventByValueParams) middleware.Responder {
	res, err := h.useCase.GetEventTypesByEventValue(int(params.Value))
	if err != nil {
		log.Println(err.Error())
		return operations.NewGetEventByValueInternalServerError().WithPayload(
			&models.ErrorResponse{
				Success: false,
				Error:   err.Error(),
			})
	}

	return operations.NewGetEventByValueOK().WithPayload(
		&models.GetEventByValueResponse{
			Success: true,
			Types:   res,
		})
}

// Handle обрабатывает запрос GET /events/day/{day}.
func (h GetEventByDay) Handle(params operations.GetEventByDayParams) middleware.Responder {
	res, err := h.useCase.GetEventByDay(params.Day)
	if err != nil {
		log.Println(err.Error())
		return operations.NewGetEventByDayInternalServerError().WithPayload(
			&models.ErrorResponse{
				Success: false,
				Error:   err.Error(),
			})
	}

	return operations.NewGetEventByDayOK().WithPayload(
		&models.GetEventByDayResponse{
			Success: true,
			Events:  res,
		})
}

// Handle обрабатывает запрос GET /user/{value}.
func (h GetUserByValue) Handle(params operations.GetUserByValueParams) middleware.Responder {
	res, err := h.useCase.GetUserByUniqueEventTypesValue(int(params.Value))
	if err != nil {
		log.Println(err.Error())
		return operations.NewGetUserByValueInternalServerError().WithPayload(
			&models.ErrorResponse{
				Success: false,
				Error:   err.Error(),
			})
	}
	return operations.NewGetUserByValueOK().WithPayload(
		&models.GetUserByValueResponse{
			Success:  true,
			UsersIds: res,
		})
}

// Handle обрабатывает запрос GET /event.
func (h GetEvent) Handle(params operations.GetEventParams) middleware.Responder {
	res, err := h.useCase.GetEventByTypeAndDate(params.Type, params.Date)
	if err != nil {
		log.Println(err.Error())
		return operations.NewGetEventInternalServerError().WithPayload(
			&models.ErrorResponse{
				Success: false,
				Error:   err.Error(),
			})
	}
	return operations.NewGetEventOK().WithPayload(
		&models.GetEventResponse{
			Success: true,
			Events:  res,
		})
}

// Handle обрабатывает запрос POST /event/new.
func (h NewEvent) Handle(params operations.NewEventParams) middleware.Responder {
	if params.Body == nil {
		return operations.NewNewEventBadRequest().WithPayload(
			&models.ErrorResponse{
				Success: false,
				Error:   errors.New("body is null").Error(),
			})
	}
	err := h.useCase.NewEvent(*params.Body)
	if err != nil {
		log.Println(err.Error())
		return operations.NewNewEventInternalServerError().WithPayload(
			&models.ErrorResponse{
				Success: false,
				Error:   err.Error(),
			})
	}
	return operations.NewNewEventOK().WithPayload(
		&models.NewEventResponse{
			Success: true,
		})
}
