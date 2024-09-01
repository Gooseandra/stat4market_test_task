package envents

type UseCase interface {
	GetEventTypesByEventValue(value string) ([]EventTypesValueResponse, error)
	GetEventByDay(day string) ([]Event, error)
	GetUserByUniqueEventTypesValue(value string) ([]UserByTypesValueResponse, error)
	GetEventByTypeAndDate(date string, eventType EventType) ([]Event, error)

	NewEvent(event Event) error
}
