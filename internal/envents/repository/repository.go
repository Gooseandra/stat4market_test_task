package repository

import (
	"database/sql"
	"test_task_stat4market/internal/envents"
	"test_task_stat4market/internal/generated/models"
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

func (r EventsRepository) GetEventTypesByEventValueRepo(value int) ([]*models.EventType, error) {
	query := `SELECT eventType, COUNT(*) AS event_count
		FROM events GROUP BY eventType
        HAVING event_count > $1`
	rows, err := r.db.Query(query, value)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Закрываем rows после завершения работы с ним

	var result []*models.EventType
	for rows.Next() {
		temp := &models.EventType{} // Создаем новый экземпляр EventType для каждой строки
		err = rows.Scan(&temp.Types, &temp.Value)
		if err != nil {
			return nil, err
		}
		result = append(result, temp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r EventsRepository) GetEventByDayRepo(dayStart time.Time, dayEnd time.Time) ([]*models.EventDetail, error) {
	query := `SELECT eventID, eventType, userID, eventTime, payload
        FROM events
        WHERE eventTime >= $1 AND eventTime < $2`
	rows, err := r.db.Query(query, dayStart, dayEnd)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Закрываем rows, чтобы освободить ресурсы
	var result []*models.EventDetail
	for rows.Next() {
		temp := new(models.EventDetail) // Инициализируем temp
		err = rows.Scan(&temp.EventID, &temp.EventType, &temp.UserID, &temp.EventTime, &temp.Payload)
		if err != nil {
			return nil, err
		}
		result = append(result, temp)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (r EventsRepository) GetUserByUniqueEventTypesValueRepo(value int) ([]*models.UserDetail, error) {
	query := `SELECT userID, COUNT(DISTINCT eventType) AS unique_event_types
        FROM events
        GROUP BY userID
        HAVING unique_event_types > $1`
	rows, err := r.db.Query(query, value)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Закрываем rows, чтобы освободить ресурсы

	var result []*models.UserDetail
	for rows.Next() {
		temp := &models.UserDetail{} // Инициализируем temp для записи строки
		err = rows.Scan(&temp.User, &temp.Value)
		if err != nil {
			return nil, err
		}
		result = append(result, temp)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (r EventsRepository) GetEventByTypeAndDateRepo(dayStart time.Time, dayEnd time.Time, eventType string) ([]*models.EventDetail, error) {
	query := `SELECT eventID, eventType, userID, eventTime, payload
    FROM events
    WHERE eventTime >= $1 AND eventTime < $2
      AND eventType = $3`
	rows, err := r.db.Query(query, dayStart, dayEnd, eventType)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Закрываем rows, чтобы освободить ресурсы

	var result []*models.EventDetail
	for rows.Next() {
		temp := &models.EventDetail{} // Инициализируем temp
		err = rows.Scan(&temp.EventID, &temp.EventType, &temp.UserID, &temp.EventTime, &temp.Payload)
		if err != nil {
			return nil, err
		}
		result = append(result, temp)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (r EventsRepository) NewEventRepo(event models.NewEventRequest) error {
	query := `INSERT INTO events(eventID, eventType, userID, eventTime, payload)
		VALUES ($1,$2,$3,$4,$5)` // тут, надеюсь, комментарии излишни
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
