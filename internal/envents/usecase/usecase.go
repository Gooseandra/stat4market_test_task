package usecase

import (
	"test_task_stat4market/config"
	"test_task_stat4market/internal/envents"
	"test_task_stat4market/internal/generated/models"
	"time"
)

type UseCase struct {
	Repo envents.Repository
	cfg  *config.Config
}

// NewUseCase создаёт новый объект UseCase
func NewUseCase(repo envents.Repository) envents.UseCase {
	return UseCase{Repo: repo}
}

func (u UseCase) GetEventTypesByEventValue(value int) ([]*models.EventType, error) {
	result, err := u.Repo.GetEventTypesByEventValueRepo(value) //логики тут нет, просто слой useCase
	return result, err
}

func (u UseCase) GetEventByDay(day string) ([]*models.EventDetail, error) {
	eventDate, err := time.Parse("2006-01-02", day)
	if err != nil {
		return nil, err
	}

	startOfDay := eventDate
	endOfDay := startOfDay.Add(24 * time.Hour) //формирование временного диапозона для поиска в бд

	result, err := u.Repo.GetEventByDayRepo(startOfDay, endOfDay)
	return result, err
}

func (u UseCase) GetUserByUniqueEventTypesValue(value int) ([]*models.UserDetail, error) {
	result, err := u.Repo.GetUserByUniqueEventTypesValueRepo(value) //логики тут нет, просто слой useCase
	return result, err
}

func (u UseCase) NewEvent(event models.NewEventRequest) error {
	err := u.Repo.NewEventRepo(event) //логики тут нет, просто слой useCase
	return err
}

func (u UseCase) GetEventByTypeAndDate(date string, eventType string) ([]*models.EventDetail, error) {
	eventDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	startOfDay := eventDate
	endOfDay := startOfDay.Add(24 * time.Hour) //формирование временного диапозона для поиска в бд

	result, err := u.Repo.GetEventByTypeAndDateRepo(startOfDay, endOfDay, eventType)
	return result, err
}
