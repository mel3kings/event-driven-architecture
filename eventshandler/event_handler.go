package eventshandler

import "github.com/mel3kings/event-driven-architecture/events"

type EventHandler interface {
	HandleEvent()
}

func NewEventHandler(polledEvent *events.Event, repo *events.EventRepo) EventHandler {
	switch polledEvent.EventType {
	case events.ApplicationStarted:
		return ApplicationStartedEventHandler{
			event:     *polledEvent,
			eventRepo: *repo,
		}
	}
	return nil
}
