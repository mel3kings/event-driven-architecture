package events

import "github.com/pborman/uuid"

type EventRepo interface {
	CreateEvent(event Event)
	GetNextEvent() (event Event)
	SetEventHandled(eventID uuid.UUID)
}

// this is just pseudo code, in real application we would have a persistent storage to handle the events
type InMemoryEventRepo struct {
	currentEvent *Event
}

func (repo InMemoryEventRepo) CreateEvent(event Event) {
	repo.currentEvent = &event
}

func (repo InMemoryEventRepo) GetNextEvent() (event Event) {
	if repo.currentEvent != nil{
		return *repo.currentEvent
	}
	return Event{}
}

func (repo InMemoryEventRepo) SetEventHandled(eventID uuid.UUID) {
	repo.currentEvent = nil
}
