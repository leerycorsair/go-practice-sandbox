package models

import "time"

type EventT struct {
	Id     int       `json:"id" db:"id"`
	UserId int       `json:"user_id" db:"user_id"`
	Name   string    `json:"name" db:"name"`
	Date   time.Time `json:"date" db:"date"`
}

type EventSortT []EventT

func (e EventSortT) Len() int { return len(e) }

func (e EventSortT) Less(i, j int) bool {
	return e[i].Date.Before(e[j].Date)
}

func (e EventSortT) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
