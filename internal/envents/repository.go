package envents

import (
	"test_task_stat4market/internal/generated/models"
	"time"
)

//интерфейс репозитория

type Repository interface {
	GetEventTypesByEventValueRepo(value int) ([]*models.EventType, error)
	GetEventByDayRepo(dayStart time.Time, dayEnd time.Time) ([]*models.EventDetail, error)
	GetUserByUniqueEventTypesValueRepo(value int) ([]*models.UserDetail, error)
	GetEventByTypeAndDateRepo(dayStart time.Time, dayEnd time.Time, eventType string) ([]*models.EventDetail, error)

	NewEventRepo(event models.NewEventRequest) error
}
