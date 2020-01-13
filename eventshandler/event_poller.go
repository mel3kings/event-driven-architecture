package eventshandler

import (
	"fmt"
	"github.com/mel3kings/event-driven-architecture/events"
	"time"
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
			time.Sleep(time.Second * 10)
		}
	}()
}

func (handler *EventPoller) Stop() {
	handler.EventPool.shutDown()
}
