package handler

import (
	"develop/dev11/internal/dto"
	"develop/dev11/internal/middleware"
	"develop/dev11/internal/service"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type CalendarHandlers struct {
	service service.CalendarService
	mux     *http.ServeMux
}

func NewCalendarHandlers(mux *http.ServeMux, service service.CalendarService) (*CalendarHandlers, error) {
	return &CalendarHandlers{mux: mux, service: service}, nil
}

func (h *CalendarHandlers) InitRoutes() {
	h.mux.Handle(`POST /create_event`, middleware.Logging(http.HandlerFunc(h.CreateEvent)))
	h.mux.Handle(`POST /update_event`, middleware.Logging(http.HandlerFunc(h.UpdateEvent)))
	h.mux.Handle(`POST /delete_event`, middleware.Logging(http.HandlerFunc(h.DeleteEvent)))

	h.mux.Handle(`GET /events_for_day`, middleware.Logging(http.HandlerFunc(h.GetEventsForDay)))
	h.mux.Handle(`GET /events_for_week`, middleware.Logging(http.HandlerFunc(h.GetEventsForWeek)))
	h.mux.Handle(`GET /events_for_month`, middleware.Logging(http.HandlerFunc(h.GetEventsForMonth)))
}

func (h *CalendarHandlers) Run(host string) error {
	h.InitRoutes()
	return http.ListenAndServe(host, h.mux)
}

func (h *CalendarHandlers) handleError(w http.ResponseWriter, r *http.Request, err string, statusCode int) {
	log.Printf("%s %s %s %s", r.Method, r.RequestURI, "error:", err)
	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err)))
}

func (h *CalendarHandlers) buildResponseBody(events dto.DTOGetEvents) string {
	var builder strings.Builder
	builder.WriteString(`{"result" : [`)
	for i, event := range events.Events {
		builder.WriteString(`{`)
		counter := 0
		for key, value := range event {
			switch value.(type) {
			case int:
				builder.WriteString(fmt.Sprintf(`"%s" : %v`, key, value.(int)))
				break
			case time.Time:
				builder.WriteString(fmt.Sprintf(`"%s" : "%s"`, key, value.(time.Time).Format(`2006-01-02`)))
				break
			default:
				builder.WriteString(fmt.Sprintf(`"%s" : "%s"`, key, value.(string)))
			}
			if counter != len(event)-1 {
				builder.WriteString(`, `)
			}
			counter++
		}
		if i != len(events.Events)-1 {
			builder.WriteString(`}, `)
		} else {
			builder.WriteString(`}`)
		}
	}
	builder.WriteString(`]}`)
	return builder.String()
}
