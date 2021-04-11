package events

import (
	"github.com/pborman/uuid"
)

type EventRepo interface {
	CreateEvent(event *Event)
	GetNextEvent() (event *Event)
	SetEventHandled(eventID uuid.UUID)
}

// this is just pseudo code, in real application we would have a persistent storage to handle the events
type InMemoryEventRepo struct {
	CurrentEvent *Event
}

func (Repo *InMemoryEventRepo) CreateEvent(event *Event) {
	Repo.CurrentEvent = event
}

func (Repo *InMemoryEventRepo) GetNextEvent() (event *Event) {
	if Repo.CurrentEvent != nil {
		return Repo.CurrentEvent
	}
	return nil
}

func (Repo *InMemoryEventRepo) SetEventHandled(eventID uuid.UUID) {
	Repo.CurrentEvent = nil
}

func (Repo *InMemoryEventRepo) VerifyEvent() {
	if Repo.CurrentEvent != nil {
		println(Repo.CurrentEvent.Metadata)
	} else {
		println("Event not found")
	}

}
