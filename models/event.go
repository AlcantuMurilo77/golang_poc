package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time 
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() {
	//TODO: implement database logic
	events = append(events, e)
}

func GetAllEvents() []Event {
	//TODO: implement database logic
	return events

}
