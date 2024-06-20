package service

import (
	"develop/dev11/internal/dto"
	"develop/dev11/internal/models"
	"develop/dev11/internal/request"
)

type CalendarService interface {
	CreateEvent(event request.CreateEventRequest) error
	UpdateEvent(event request.UpdateEventRequest) error
	DeleteEvent(id int) error

	GetEvents(req interface{}) (dto.DTOGetEvents, error)
	EntityToMap(e models.EventT) map[string]interface{}
}

type Service struct {
	CalendarService
}
