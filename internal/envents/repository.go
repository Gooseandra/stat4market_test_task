package envents

import "time"

type Repository interface {
	GetEventTypesByEventValueRepo(value int) ([]EventTypesValueResponse, error)
	GetEventByDayRepo(dayStart time.Time, dayEnd time.Time) ([]Event, error)
	GetUserByUniqueEventTypesValueRepo(value int) ([]UserByTypesValueResponse, error)
	GetEventByTypeAndDateRepo(dayStart EventTime, dayEnd EventTime, eventType EventType) ([]Event, error)

	NewEventRepo(event Event) error
}
