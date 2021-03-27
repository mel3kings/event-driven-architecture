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

func (Repo *InMemoryEventRepo) CreateEvent(event Event) {
	Repo.currentEvent = &event
}

func (Repo *InMemoryEventRepo) GetNextEvent() (event Event) {
	if Repo.currentEvent != nil {
		return *Repo.currentEvent
	}
	return Event{}
}

func (Repo *InMemoryEventRepo) SetEventHandled(eventID uuid.UUID) {
	Repo.currentEvent = nil
}

func (Repo *InMemoryEventRepo) VerifyEvent() {
	if Repo.currentEvent != nil {
		println(Repo.currentEvent.Metadata)
	} else {
		println("Event not found")
	}

}
