package usecase

import (
	"strconv"
	"test_task_stat4market/config"
	"test_task_stat4market/internal/envents"
	"time"
)

type UseCase struct {
	Repo envents.Repository
	cfg  *config.Config
}

func NewUseCase(repo envents.Repository) envents.UseCase {
	return UseCase{Repo: repo}
}

func (u UseCase) GetEventTypesByEventValue(value string) ([]envents.EventTypesValueResponse, error) {
	intVal, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}
	result, err := u.Repo.GetEventTypesByEventValueRepo(intVal)
	return result, err
}

func (u UseCase) GetEventByDay(day string) ([]envents.Event, error) {
	eventDate, err := time.Parse("2006-01-02", day)
	if err != nil {
		return nil, err
	}

	startOfDay := eventDate
	endOfDay := startOfDay.Add(24 * time.Hour)

	result, err := u.Repo.GetEventByDayRepo(startOfDay, endOfDay)
	return result, err
}

func (u UseCase) GetUserByUniqueEventTypesValue(value string) ([]envents.UserByTypesValueResponse, error) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}
	result, err := u.Repo.GetUserByUniqueEventTypesValueRepo(intValue)
	return result, err
}

func (u UseCase) NewEvent(event envents.Event) error {
	err := u.Repo.NewEventRepo(event)
	return err
}

func (u UseCase) GetEventByTypeAndDate(date string, eventType envents.EventType) ([]envents.Event, error) {
	eventDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	startOfDay := eventDate
	endOfDay := startOfDay.Add(24 * time.Hour)

	result, err := u.Repo.GetEventByTypeAndDateRepo(startOfDay, endOfDay, eventType)
	return result, err
}
