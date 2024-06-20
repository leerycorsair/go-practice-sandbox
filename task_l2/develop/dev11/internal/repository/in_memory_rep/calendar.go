package inmemoryrep

import (
	"develop/dev11/internal/models"
	"develop/dev11/internal/storage"
	"fmt"
	"log"
	"time"
)

type InMemoryRep struct {
	db *storage.CalendarStorage
}

func NewInMemoryRep(db *storage.CalendarStorage) (*InMemoryRep, error) {
	return &InMemoryRep{db: db}, nil
}

func (r *InMemoryRep) CreateEvent(event models.EventT) error {
	event.Id = r.db.Len() + 1
	r.db.Set(event.Id, event)
	log.Print(r.db.Len())
	return nil
}

func (r *InMemoryRep) UpdateEvent(event models.EventT) error {
	if _, ok := r.db.Get(event.Id); ok {
		r.db.Set(event.Id, event)
	} else {
		return fmt.Errorf("event with id=%v is not defined", event.Id)
	}
	return nil
}

func (r *InMemoryRep) DeleteEvent(id int) error {
	if _, ok := r.db.Get(id); ok {
		r.db.Set(id, models.EventT{})
	} else {
		return fmt.Errorf("event with id=%v is not defined", id)
	}
	return nil
}

func (r *InMemoryRep) GetEventsForDay(userId int, date time.Time) ([]models.EventT, error) {
	result := make([]models.EventT, 0)
	r.db.Range(func(key int, event models.EventT) bool {
		if event.Date == date && event.UserId == userId {
			result = append(result, event)
		}
		return true
	})
	return result, nil
}

func (r *InMemoryRep) GetEventsForWeek(userId int, date time.Time) ([]models.EventT, error) {
	result := make([]models.EventT, 0)
	year, week := date.ISOWeek()
	r.db.Range(func(key int, event models.EventT) bool {
		yearTemp, weekTemp := event.Date.ISOWeek()
		if yearTemp == year && weekTemp == week && event.UserId == userId {
			result = append(result, event)
		}
		return true
	})
	return result, nil
}

func (r *InMemoryRep) GetEventsForMonth(userId int, date time.Time) ([]models.EventT, error) {
	result := make([]models.EventT, 0)
	month, year := date.Month(), date.Year()
	r.db.Range(func(key int, event models.EventT) bool {
		yearTemp, monthTemp := event.Date.Year(), event.Date.Month()
		if yearTemp == year && monthTemp == month && event.UserId == userId {
			result = append(result, event)
		}
		return true
	})
	return result, nil
}
