package handler

import (
	"develop/dev11/internal/request"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func (h *CalendarHandlers) validateGetRequest(r *http.Request) (int, time.Time, error) {
	if err := r.ParseForm(); err != nil {
		return 0, time.Time{}, err
	}
	userIdString, dateString := r.Form.Get(`user_id`), r.Form.Get(`date`)
	if userIdString == "" || dateString == "" || len(r.Form) > 2 {
		return 0, time.Time{}, errors.New(`empty user_id or date fields`)
	}
	if matched, err := regexp.MatchString(`\d+`, userIdString); !matched || err != nil {
		return 0, time.Time{}, errors.New(`user_id validation error`)
	}
	userId, _ := strconv.Atoi(userIdString)
	date, err := time.Parse(`2006-01-02`, dateString)
	if err != nil {
		return 0, time.Time{}, errors.New(`date validation error`)
	}
	return userId, date, nil
}

func (h *CalendarHandlers) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	userId, date, err := h.validateGetRequest(r)
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	events, err := h.service.GetEvents(request.GetEventsForDayRequest{UserId: userId, Date: date})
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(h.buildResponseBody(events)))
}

func (h *CalendarHandlers) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	userId, date, err := h.validateGetRequest(r)
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	events, err := h.service.GetEvents(request.GetEventsForWeekRequest{UserId: userId, Date: date})
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(h.buildResponseBody(events)))
}

func (h *CalendarHandlers) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	userId, date, err := h.validateGetRequest(r)
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	events, err := h.service.GetEvents(request.GetEventsForMonthRequest{UserId: userId, Date: date})
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(h.buildResponseBody(events)))
}

func UnmarshallBody[T any](body []byte) (T, error) {
	var request T
	if err := json.Unmarshal(body, &request); err != nil {
		return request, err
	}
	return request, nil
}

func (h *CalendarHandlers) CreateEvent(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	request, err := UnmarshallBody[request.CreateEventRequest](body)
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateEvent(request); err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(`{"result":"ok"}`))
}

func (h *CalendarHandlers) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	request, err := UnmarshallBody[request.UpdateEventRequest](body)
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateEvent(request); err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(`{"result":"ok"}`))
}

func (h *CalendarHandlers) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	request, err := UnmarshallBody[request.DeleteEventRequest](body)
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteEvent(request.Id); err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(`{"result":"ok"}`))
}
