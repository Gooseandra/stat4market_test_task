package envents

import (
	"test_task_stat4market/models"
)

//интерфейс UseCase

type UseCase interface {
	GetEventTypesByEventValue(value int) ([]*models.EventType, error)
	GetEventByDay(day string) ([]*models.EventDetail, error)
	GetUserByUniqueEventTypesValue(value int) ([]*models.UserDetail, error)
	GetEventByTypeAndDate(date string, eventType string) ([]*models.EventDetail, error)

	NewEvent(event models.NewEventRequest) error
}
