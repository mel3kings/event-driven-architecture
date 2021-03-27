package eventshandler

import (
	"fmt"
	"time"

	"github.com/mel3kings/event-driven-architecture/events"
)

type EventPoller struct {
	PolledEvent events.Event
	EventRepo   events.EventRepo
	EventPool   EventPool
}

func (handler *EventPoller) Start() {
	go func() {
		for !handler.EventPool.shuttingDown {
			fmt.Println("polling event")
			polledEvent := handler.EventRepo.GetNextEvent()
			eventHandler := NewEventHandler(&polledEvent)
			handler.EventPool.add(eventHandler)
			time.Sleep(time.Second * 5)
		}
	}()
}

func (handler *EventPoller) Stop() {
	handler.EventPool.shutDown()
}
