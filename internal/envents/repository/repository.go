package repository

import (
	"database/sql"
	"test_task_stat4market/internal/envents"
	"time"
)

type EventsRepository struct {
	db *sql.DB
}

func NewClickHouseRepository(db *sql.DB) envents.Repository {
	return &EventsRepository{
		db: db,
	}
}

//	я решил, что стоит сделать систему более гибкой и добавить возможность сдлеть более точные запросы, также добавил
//	возможность получать эти данные по api исходные запросы для решения поставленных задач:
//1)
//	SELECT eventType, COUNT(*) AS event_count
//	FROM events
//	GROUP BY eventType
//	HAVING event_count > 1000;
//2)
//	SELECT eventID, eventType, userID, eventTime, payload
//	FROM events
//	WHERE toStartOfMonth(eventTime) = toDate(eventTime)
//3)
//	SELECT userID, COUNT(DISTINCT eventType) AS unique_event_types
//	FROM events
//	GROUP BY userID
//	HAVING unique_event_types > 3;

func (r EventsRepository) GetEventTypesByEventValueRepo(value int) ([]envents.EventTypesValueResponse, error) {
	query := `SELECT eventType, COUNT(*) AS event_count
		FROM events GROUP BY eventType
        HAVING event_count > $1`
	rows, err := r.db.Query(query, value)
	if err != nil {
		return nil, err
	}
	var temp envents.EventTypesValueResponse
	var result []envents.EventTypesValueResponse
	for rows.Next() {
		err = rows.Scan(&temp.Types, &temp.Value)
		if err != nil {
			return nil, err
		}
		result = append(result, temp)
	}
	return result, nil
}

func (r EventsRepository) GetEventByDayRepo(dayStart time.Time, dayEnd time.Time) ([]envents.Event, error) {
	query := `SELECT eventID, eventType, userID, eventTime, payload
        FROM events
        WHERE eventTime >= $1 AND eventTime < $2`
	rows, err := r.db.Query(query, dayStart, dayEnd)
	if err != nil {
		return nil, err
	}
	var temp envents.Event
	var result []envents.Event
	for rows.Next() {
		err = rows.Scan(&temp.EventID, &temp.EventType, &temp.UserID, &temp.EventTime, &temp.Payload)
		if err != nil {
			return nil, err
		}
		result = append(result, temp)
	}
	return result, nil
}

func (r EventsRepository) GetUserByUniqueEventTypesValueRepo(value int) ([]envents.UserByTypesValueResponse, error) {
	query := `SELECT userID, COUNT(DISTINCT eventType) AS unique_event_types
        FROM events
        GROUP BY userID
        HAVING unique_event_types > $1`
	rows, err := r.db.Query(query, value)
	if err != nil {
		return nil, err
	}
	var temp envents.UserByTypesValueResponse
	var result []envents.UserByTypesValueResponse
	for rows.Next() {
		err = rows.Scan(&temp.User, &temp.Value)
		if err != nil {
			return nil, err
		}
		result = append(result, temp)
	}
	return result, nil
}

func (r EventsRepository) NewEventRepo(event envents.Event) error {
	query := `INSERT INTO events(eventID, eventType, userID, eventTime, payload)
		VALUES ($1,$2,$3,$4,$5)`
	_, err := r.db.Exec(query,
		event.EventID,
		event.EventType,
		event.UserID,
		event.EventTime,
		event.Payload)
	if err != nil {
		return err
	}
	return nil
}

func (r EventsRepository) GetEventByTypeAndDateRepo(dayStart envents.EventTime, dayEnd envents.EventTime,
	eventType envents.EventType) ([]envents.Event, error) {
	query := `SELECT eventID, eventType, userID, eventTime, payload
    FROM events
    WHERE eventTime >= $1 AND eventTime < $2
      AND eventType = $3`
	rows, err := r.db.Query(query, dayStart, dayEnd, eventType)
	if err != nil {
		return nil, err
	}
	var temp envents.Event
	var result []envents.Event
	for rows.Next() {
		err = rows.Scan(&temp.EventID, &temp.EventType, &temp.UserID, &temp.EventTime, &temp.Payload)
		if err != nil {
			return nil, err
		}
		result = append(result, temp)
	}
	return result, nil
}
