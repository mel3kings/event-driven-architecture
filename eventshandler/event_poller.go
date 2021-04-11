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
			time.Sleep(3 * time.Second)
			fmt.Println("polling event")
			polledEvent := handler.EventRepo.GetNextEvent()
			if polledEvent != nil {
				eventHandler := NewEventHandler(
					polledEvent, &handler.EventRepo)
				handler.EventPool.add(eventHandler)
				time.Sleep(time.Second * 5)
			} else {
				fmt.Println("no event to process")
			}
		}
	}()
}

func (handler *EventPoller) Stop() {
	handler.EventPool.shutDown()
}
