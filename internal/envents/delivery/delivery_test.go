package delivery

import (
	"errors"
	"github.com/go-openapi/strfmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"test_task_stat4market/internal/mocks"
	"test_task_stat4market/models"
	"test_task_stat4market/restapi/operations"
	"testing"
	"time"
)

func TestGetEventByValue_Handle_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mocks.NewMockUseCase(ctrl)
	handler := GetEventByValue{event{useCase: mockUseCase}}

	mockUseCase.EXPECT().
		GetEventTypesByEventValue(1).
		Return([]*models.EventType{{"type1", 2}, {"type3", 4}}, nil)

	params := operations.GetEventByValueParams{Value: 1}
	response := handler.Handle(params)

	okResponse := response.(*operations.GetEventByValueOK)
	assert.True(t, okResponse.Payload.Success)
	assert.Equal(t, []*models.EventType{{"type1", 2}, {"type3", 4}}, okResponse.Payload.Types)
}

func TestGetEventByValue_Handle_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mocks.NewMockUseCase(ctrl)
	handler := GetEventByValue{event{useCase: mockUseCase}}

	mockUseCase.EXPECT().
		GetEventTypesByEventValue(1).
		Return(nil, errors.New("some error"))

	params := operations.GetEventByValueParams{Value: 1}
	response := handler.Handle(params)

	errorResponse := response.(*operations.GetEventByValueInternalServerError)
	assert.False(t, errorResponse.Payload.Success)
	assert.Equal(t, "some error", errorResponse.Payload.Error)
}

func TestGetEventByDay_Handle_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mocks.NewMockUseCase(ctrl)
	handler := GetEventByDay{event{useCase: mockUseCase}}

	mockUseCase.EXPECT().
		GetEventByDay("2023-09-08").
		Return([]*models.EventDetail{{EventID: 1}, {EventID: 2}}, nil)

	params := operations.GetEventByDayParams{Day: "2023-09-08"}
	response := handler.Handle(params)

	okResponse := response.(*operations.GetEventByDayOK)
	assert.True(t, okResponse.Payload.Success)
	assert.Equal(t, []*models.EventDetail{{EventID: 1}, {EventID: 2}}, okResponse.Payload.Events)
}

func TestGetEventByDay_Handle_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mocks.NewMockUseCase(ctrl)
	handler := GetEventByDay{event{useCase: mockUseCase}}

	mockUseCase.EXPECT().
		GetEventByDay("2023-09-08").
		Return(nil, errors.New("some error"))

	params := operations.GetEventByDayParams{Day: "2023-09-08"}
	response := handler.Handle(params)

	errorResponse := response.(*operations.GetEventByDayInternalServerError)
	assert.False(t, errorResponse.Payload.Success)
	assert.Equal(t, "some error", errorResponse.Payload.Error)
}

func TestGetUserByValue_Handle_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mocks.NewMockUseCase(ctrl)
	handler := GetUserByValue{event{useCase: mockUseCase}}

	mockUseCase.EXPECT().
		GetUserByUniqueEventTypesValue(1).
		Return([]*models.UserDetail{{2, 2}, {1, 1}}, nil)

	params := operations.GetUserByValueParams{Value: 1}
	response := handler.Handle(params)

	okResponse := response.(*operations.GetUserByValueOK)
	assert.True(t, okResponse.Payload.Success)
	assert.Equal(t, []*models.UserDetail{{2, 2}, {1, 1}}, okResponse.Payload.UsersIds)
}

func TestGetUserByValue_Handle_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mocks.NewMockUseCase(ctrl)
	handler := GetUserByValue{event{useCase: mockUseCase}}

	mockUseCase.EXPECT().
		GetUserByUniqueEventTypesValue(1).
		Return(nil, errors.New("some error"))

	params := operations.GetUserByValueParams{Value: 1}
	response := handler.Handle(params)

	errorResponse := response.(*operations.GetUserByValueInternalServerError)
	assert.False(t, errorResponse.Payload.Success)
	assert.Equal(t, "some error", errorResponse.Payload.Error)
}

func TestGetEvent_Handle_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mocks.NewMockUseCase(ctrl)
	handler := GetEvent{event{useCase: mockUseCase}}

	mockUseCase.EXPECT().
		GetEventByTypeAndDate("type1", "2023-09-08").
		Return([]*models.EventDetail{{EventID: 1}, {EventID: 2}}, nil)

	params := operations.GetEventParams{Type: "type1", Date: "2023-09-08"}
	response := handler.Handle(params)

	okResponse := response.(*operations.GetEventOK)
	assert.True(t, okResponse.Payload.Success)
	assert.Equal(t, []*models.EventDetail{{EventID: 1}, {EventID: 2}}, okResponse.Payload.Events)
}

func TestGetEvent_Handle_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mocks.NewMockUseCase(ctrl)
	handler := GetEvent{event{useCase: mockUseCase}}

	mockUseCase.EXPECT().
		GetEventByTypeAndDate("type1", "2023-09-08").
		Return(nil, errors.New("some error"))

	params := operations.GetEventParams{Type: "type1", Date: "2023-09-08"}
	response := handler.Handle(params)

	errorResponse := response.(*operations.GetEventInternalServerError)
	assert.False(t, errorResponse.Payload.Success)
	assert.Equal(t, "some error", errorResponse.Payload.Error)
}

func TestNewEvent_Handle_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mocks.NewMockUseCase(ctrl)
	handler := NewEvent{event{useCase: mockUseCase}}

	body := &models.NewEventRequest{
		EventType: "type1",
		EventTime: strfmt.DateTime(time.Now()),
	}

	mockUseCase.EXPECT().
		NewEvent(*body).
		Return(nil)

	params := operations.NewEventParams{Body: body}
	response := handler.Handle(params)

	okResponse := response.(*operations.NewEventOK)
	assert.True(t, okResponse.Payload.Success)
}

func TestNewEvent_Handle_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mocks.NewMockUseCase(ctrl)
	handler := NewEvent{event{useCase: mockUseCase}}

	response := handler.Handle(operations.NewEventParams{Body: nil})

	badRequestResponse := response.(*operations.NewEventBadRequest)
	assert.False(t, badRequestResponse.Payload.Success)
	assert.Equal(t, "body is null", badRequestResponse.Payload.Error)
}

func TestNewEvent_Handle_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mocks.NewMockUseCase(ctrl)
	handler := NewEvent{event{useCase: mockUseCase}}

	body := &models.NewEventRequest{
		EventType: "type1",
		EventTime: strfmt.DateTime(time.Now()),
	}

	mockUseCase.EXPECT().
		NewEvent(*body).
		Return(errors.New("some error"))

	params := operations.NewEventParams{Body: body}
	response := handler.Handle(params)

	errorResponse := response.(*operations.NewEventInternalServerError)
	assert.False(t, errorResponse.Payload.Success)
	assert.Equal(t, "some error", errorResponse.Payload.Error)
}
