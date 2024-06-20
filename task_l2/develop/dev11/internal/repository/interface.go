package repository

import (
	"develop/dev11/internal/models"
	"time"
)

type CalendarRepository interface {
	CreateEvent(event models.EventT) error
	UpdateEvent(event models.EventT) error
	DeleteEvent(id int) error
	GetEventsForDay(userId int, date time.Time) ([]models.EventT, error)
	GetEventsForWeek(userId int, date time.Time) ([]models.EventT, error)
	GetEventsForMonth(userId int, date time.Time) ([]models.EventT, error)
}
