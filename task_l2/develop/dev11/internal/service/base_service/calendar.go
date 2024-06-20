package base_service

import (
	"develop/dev11/internal/dto"
	"develop/dev11/internal/models"
	"develop/dev11/internal/repository"
	"develop/dev11/internal/request"
	"sort"
)

type CalendarService struct {
	repo repository.CalendarRepository
}

func NewCalendarService(repo repository.CalendarRepository) (*CalendarService, error) {
	return &CalendarService{repo: repo}, nil
}

func (s *CalendarService) CreateEvent(event request.CreateEventRequest) error {
	return s.repo.CreateEvent(models.EventT{UserId: event.UserId, Name: event.Name, Date: event.Date.Time})
}

func (s *CalendarService) UpdateEvent(event request.UpdateEventRequest) error {
	return s.repo.UpdateEvent(models.EventT{Id: event.Id, UserId: event.UserId, Name: event.Name, Date: event.Date.Time})
}

func (s *CalendarService) DeleteEvent(id int) error {
	return s.repo.DeleteEvent(id)
}

func (s *CalendarService) GetEvents(req interface{}) (dto.DTOGetEvents, error) {
	var events []models.EventT
	var err error
	switch req.(type) {
	case request.GetEventsForDayRequest:
		events, err = s.repo.GetEventsForDay(
			req.(request.GetEventsForDayRequest).UserId,
			req.(request.GetEventsForDayRequest).Date)
		break
	case request.GetEventsForWeekRequest:
		events, err = s.repo.GetEventsForWeek(
			req.(request.GetEventsForWeekRequest).UserId,
			req.(request.GetEventsForWeekRequest).Date)
		break
	case request.GetEventsForMonthRequest:
		events, err = s.repo.GetEventsForMonth(
			req.(request.GetEventsForMonthRequest).UserId,
			req.(request.GetEventsForMonthRequest).Date)
		break

	}
	if err != nil {
		return dto.DTOGetEvents{}, err
	}
	sort.Sort(models.EventSortT(events))
	var resp dto.DTOGetEvents
	for _, e := range events {
		resp.Events = append(resp.Events, s.EntityToMap(e))
	}
	return resp, nil
}

func (s *CalendarService) EntityToMap(e models.EventT) map[string]interface{} {
	event := map[string]interface{}{}
	event[`id`] = e.Id
	event[`user_id`] = e.UserId
	event[`name`] = e.Name
	event[`date`] = e.Date
	return event
}
