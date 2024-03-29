package eventshandler

import (
	"fmt"
	"github.com/mel3kings/event-driven-architecture/events"
)

type ApplicationStartedEventHandler struct {
	event     events.Event
	eventRepo events.EventRepo
}

func (handler ApplicationStartedEventHandler) HandleEvent() {
	fmt.Print("Application Started Event Handled", handler.event)
	fmt.Print(handler.event.Metadata)
	handler.eventRepo.SetEventHandled(handler.event.GUID)
}
