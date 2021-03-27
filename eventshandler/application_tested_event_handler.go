package eventshandler

import (
	"fmt"
	"github.com/mel3kings/event-driven-architecture/events"
)

type ApplicationTestedEventHandler struct {
	event     events.Event
	eventRepo events.EventRepo
}

func (handler ApplicationTestedEventHandler) HandleEvent() {
	fmt.Print("Application TESTED Event Handled", handler.event)
	fmt.Print(handler.event.Metadata)
	handler.eventRepo.SetEventHandled(handler.event.GUID)
}
