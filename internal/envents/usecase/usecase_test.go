package usecase

import (
	"test_task_stat4market/internal/generated/mocks"
	"test_task_stat4market/internal/generated/models"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetEventTypesByEventValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	u := UseCase{Repo: mockRepo}

	value := 123
	expectedEventTypes := []*models.EventType{{"type1", 2}, {"type2", 1}}

	mockRepo.EXPECT().GetEventTypesByEventValueRepo(value).Return(expectedEventTypes, nil)

	result, err := u.GetEventTypesByEventValue(value)
	assert.NoError(t, err)
	assert.Equal(t, expectedEventTypes, result)
}

func TestGetEventByDay(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	u := UseCase{Repo: mockRepo}

	day := "2024-09-08"
	eventDate, _ := time.Parse("2006-01-02", day)
	startOfDay := eventDate
	endOfDay := startOfDay.Add(24 * time.Hour)

	expectedEvents := []*models.EventDetail{
		{EventID: 1, EventType: "type1", EventTime: strfmt.DateTime(startOfDay), Payload: "payload1", UserID: 1},
		{EventID: 2, EventType: "type2", EventTime: strfmt.DateTime(endOfDay), Payload: "payload2", UserID: 2},
	}

	mockRepo.EXPECT().GetEventByDayRepo(startOfDay, endOfDay).Return(expectedEvents, nil)

	result, err := u.GetEventByDay(day)
	assert.NoError(t, err)
	assert.Equal(t, expectedEvents, result)
}

func TestGetUserByUniqueEventTypesValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	u := UseCase{Repo: mockRepo}

	value := 123
	expectedUsers := []*models.UserDetail{
		{User: 1, Value: 1},
		{User: 2, Value: 2},
	}

	mockRepo.EXPECT().GetUserByUniqueEventTypesValueRepo(value).Return(expectedUsers, nil)

	result, err := u.GetUserByUniqueEventTypesValue(value)
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, result)
}

func TestNewEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	u := UseCase{Repo: mockRepo}

	newEvent := models.NewEventRequest{
		EventType: "type1",
		EventTime: strfmt.DateTime(time.Now()),
		Payload:   "payload1",
		UserID:    1,
	}

	mockRepo.EXPECT().NewEventRepo(newEvent).Return(nil)

	err := u.NewEvent(newEvent)
	assert.NoError(t, err)
}

func TestGetEventByTypeAndDate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	u := UseCase{Repo: mockRepo}

	date := "2024-09-08"
	eventType := "type1"
	eventDate, _ := time.Parse("2006-01-02", date)
	startOfDay := eventDate
	endOfDay := startOfDay.Add(24 * time.Hour)

	expectedEvents := []*models.EventDetail{
		{EventID: 1, EventType: "type1", EventTime: strfmt.DateTime(startOfDay), Payload: "payload1", UserID: 1},
		{EventID: 2, EventType: "type2", EventTime: strfmt.DateTime(endOfDay), Payload: "payload2", UserID: 2},
	}

	mockRepo.EXPECT().GetEventByTypeAndDateRepo(startOfDay, endOfDay, eventType).Return(expectedEvents, nil)

	result, err := u.GetEventByTypeAndDate(date, eventType)
	assert.NoError(t, err)
	assert.Equal(t, expectedEvents, result)
}
