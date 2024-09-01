package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"test_task_stat4market/internal/envents"
	"test_task_stat4market/internal/mocks"
	"testing"
	"time"
)

func TestGetEventByTypeAndDate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	u := UseCase{Repo: mockRepo}

	date := "2024-08-30"
	eventType := "type1"
	eventID := envents.EventId(1)
	eventID2 := envents.EventId(2)
	UserId := envents.EventUserID(1)
	UserId2 := envents.EventUserID(2)
	payload := "payload"

	eventDate, _ := time.Parse("2006-01-02", date)
	startOfDay := eventDate
	endOfDay := startOfDay.Add(24 * time.Hour)

	expectedEvents := []envents.Event{
		{EventID: &eventID, EventType: &eventType, UserID: &UserId, EventTime: &eventDate, Payload: &payload},
		{EventID: &eventID2, EventType: &eventType, UserID: &UserId2, EventTime: &eventDate, Payload: &payload},
	}

	mockRepo.EXPECT().GetEventByTypeAndDateRepo(startOfDay, endOfDay, eventType).Return(expectedEvents, nil)

	result, err := u.GetEventByTypeAndDate(date, eventType)
	assert.NoError(t, err)
	assert.Equal(t, expectedEvents, result)
}

func TestGetEventByTypeAndDate_InvalidDate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	u := UseCase{Repo: mockRepo}

	date := "invalid-date"
	eventType := envents.EventType("type1")

	result, err := u.GetEventByTypeAndDate(date, eventType)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetEventByDay(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	u := UseCase{Repo: mockRepo}

	day := "2024-08-30"

	eventID := envents.EventId(1)
	eventID2 := envents.EventId(2)
	eventType := "type1"
	userID := envents.EventUserID(1)
	userID2 := envents.EventUserID(2)
	payload := "payload"

	eventDate, _ := time.Parse("2006-01-02", day)
	startOfDay := eventDate
	endOfDay := startOfDay.Add(24 * time.Hour)

	expectedEvents := []envents.Event{
		{EventID: &eventID, EventType: &eventType, UserID: &userID, EventTime: &startOfDay, Payload: &payload},
		{EventID: &eventID2, EventType: &eventType, UserID: &userID2, EventTime: &startOfDay, Payload: &payload},
	}

	mockRepo.EXPECT().GetEventByDayRepo(startOfDay, endOfDay).Return(expectedEvents, nil)

	result, err := u.GetEventByDay(day)
	assert.NoError(t, err)
	assert.Equal(t, expectedEvents, result)
}

func TestGetEventByDay_InvalidDate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	u := UseCase{Repo: mockRepo}

	day := "invalid-date"

	result, err := u.GetEventByDay(day)
	assert.Error(t, err)
	assert.Nil(t, result)
}
