package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int       `json:"userId"`
}

var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

func GetEvents() []Event {
	return events
}
