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
	println("application doing something")
	memoryRepo := events.InMemoryEventRepo{}
	applicationStartedEvent := events.Event{
		UserID:      0,
		UserIP:      "",
		GUID:        nil,
		RequestID:   nil,
		EventType:   events.ApplicationTested,
		EventStatus: 0,
		Metadata:    "TestedEvent",
	}
	memoryRepo.CreateEvent(applicationStartedEvent)
	memoryRepo.VerifyEvent()
	eventPoller := eventshandler.EventPoller{
		PolledEvent: events.Event{},
		EventRepo:   &memoryRepo,
		EventPool:   *eventshandler.NewEventHandlerPool(5),
	}
	eventPoller.Start()

	// graceful shutdown here
	incomingSignalsChannel := make(chan os.Signal, 1)
	// register events for channel
	signal.Notify(incomingSignalsChannel, os.Interrupt)
	signal.Notify(incomingSignalsChannel, syscall.SIGTERM)
	fmt.Println("Server waiting on signal.")
	go func() {
		time.Sleep(20 * time.Second)
		memoryRepo.CreateEvent( events.Event{
			UserID:      0,
			UserIP:      "",
			GUID:        nil,
			RequestID:   nil,
			EventType:   events.ApplicationStarted,
			EventStatus: 0,
			Metadata:    "StartedEvent",
		})
		println("inserting new event")
	}()
	<-incomingSignalsChannel
	delayedContext, signalAppToStop := context.WithCancel(context.Background())
	//common.LogInfo("Shutting down HTTP Server..")
	//if err := server.Shutdown(delayedContext); err != nil {
	//	common.LogErrorFormatted(fmt.Sprintf("HTTP server Shutdown"), err)
	//}
	// do some extra manual handling from here on if needed
	eventPoller.Stop()

	delayedContext.Done()
	defer signalAppToStop()
}
