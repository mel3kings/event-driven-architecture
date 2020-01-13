package eventshandler

import "sync"

type EventPool struct {
	handlerChan  chan EventHandler
	waitGroup    sync.WaitGroup
	shuttingDown bool
}

func NewEventHandlerPool(maxRoutines int) *EventPool {
	handlerPool := EventPool{handlerChan: make(chan EventHandler)}
	handlerPool.waitGroup.Add(maxRoutines)
	for i := 0; i < maxRoutines; i++ {
		go func() {
			for handler := range handlerPool.handlerChan {
				handler.HandleEvent()
			}
			if handlerPool.shuttingDown {
				handlerPool.waitGroup.Done()
			}
		}()
	}
	return &handlerPool
}

func (eventPool *EventPool) add(handler EventHandler) {
	if !eventPool.shuttingDown {
		eventPool.handlerChan <- handler
	}
}

func (eventPool *EventPool) shutDown() {
	eventPool.shuttingDown = true
	close(eventPool.handlerChan)
	eventPool.waitGroup.Wait()
}
