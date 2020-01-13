package events

import "github.com/pborman/uuid"

type Event struct {
	UserID      int         `json:"-"`
	UserIP      string      `json:"userIP"`
	GUID        uuid.UUID   `json:"guidID"`
	RequestID   uuid.UUID   `json:"requestID"`
	EventType   EventType   `json:"eventType"`
	EventStatus EventStatus `json:"-"`
	Metadata    string      `json:"-"`
}

type EventType int
type EventStatus int

const (
	ApplicationStarted EventType = 0
	ApplicationTested  EventType = 1
)

const (
	EventCreated EventStatus = 0
	EventPolled  EventStatus = 1
	EventSuccess EventStatus = 2
	EventFailure EventStatus = 3
)

func NewEvent(userID int, requestID uuid.UUID, userIP string, eventType EventType) *Event {
	return &Event{
		GUID:      uuid.NewUUID(),
		UserID:    userID,
		RequestID: requestID,
		EventType: eventType,
		UserIP:    userIP,
	}
}
