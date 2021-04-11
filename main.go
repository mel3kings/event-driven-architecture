package main

import (
	"context"
	"fmt"
	"github.com/mel3kings/event-driven-architecture/events"
	"github.com/mel3kings/event-driven-architecture/eventshandler"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	memoryRepo := events.InMemoryEventRepo{}
	eventPoller := eventshandler.EventPoller{
		PolledEvent: events.Event{},
		EventRepo:   &memoryRepo,
		EventPool:   *eventshandler.NewEventHandlerPool(5),
	}
	eventPoller.Start()
	incomingSignalsChannel := make(chan os.Signal, 1)
	signal.Notify(incomingSignalsChannel, os.Interrupt)
	signal.Notify(incomingSignalsChannel, syscall.SIGTERM)
	fmt.Println("Server waiting on signal.")
	go func() {
		for i := 1; i < 5; i++ {
			createEvent(&memoryRepo, events.ApplicationStarted)
		}
	}()
	<-incomingSignalsChannel
	delayedContext, signalAppToStop := context.WithCancel(context.Background())
	eventPoller.Stop()
	delayedContext.Done()
	defer signalAppToStop()
}

func createEvent(memoryRepo *events.InMemoryEventRepo, eventType events.EventType) {
	time.Sleep(5 * time.Second)
	println("inserting new event")
	memoryRepo.CreateEvent(&events.Event{
		UserID:      0,
		UserIP:      "",
		GUID:        nil,
		RequestID:   nil,
		EventType:   eventType,
		EventStatus: 0,
		Metadata:    "StartedEvent META",
	})
}

// safe shutdown http server
//common.LogInfo("Shutting down HTTP Server..")
//if err := server.Shutdown(delayedContext); err != nil {
//	common.LogErrorFormatted(fmt.Sprintf("HTTP server Shutdown"), err)
//}
// do some extra manual handling from here on if needed