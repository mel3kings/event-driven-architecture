package eventshandler

import "github.com/mel3kings/event-driven-architecture/events"

type EventHandler interface {
	HandleEvent()
}

func NewEventHandler(polledEvent *events.Event) EventHandler {
	switch polledEvent.EventType {
	case events.ApplicationStarted:
		return ApplicationStartedEventHandler{
			event:     *polledEvent,
			eventRepo: &events.InMemoryEventRepo{},
		}
	case events.ApplicationTested:
		return ApplicationTestedEventHandler{
			event:     *polledEvent,
			eventRepo: &events.InMemoryEventRepo{},
		}

	}
	return nil
}
