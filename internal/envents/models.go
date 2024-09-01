package envents

import "time"

type Event struct {
	EventID   *EventId      `json:"event_id"`
	EventType *EventType    `json:"event_type"`
	UserID    *EventUserID  `json:"user_id"`
	EventTime *EventTime    `json:"event_time"`
	Payload   *EventPayload `json:"payload"`
}

type EventTypesValueResponse struct {
	Types EventType
	Value int
}

type UserByTypesValueResponse struct {
	User  EventUserID
	Value int
}

type (
	EventId      = int64
	EventType    = string
	EventUserID  = int64
	EventTime    = time.Time
	EventPayload = string
)
