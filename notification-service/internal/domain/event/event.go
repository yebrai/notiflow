package event

import (
	"time"

	"github.com/google/uuid"
)

type EventType string

const (
	UserRegistered EventType = "USER_REGISTERED"
	OrderCompleted EventType = "ORDER_COMPLETED"
	PaymentFailed  EventType = "PAYMENT_FAILED"
)

type Event struct {
	ID        uuid.UUID `json:"id"`
	Type      EventType `json:"type"`
	Payload   []byte    `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}

func NewEvent(eventType EventType, payload []byte) Event {
	return Event{
		ID:        uuid.New(),
		Type:      eventType,
		Payload:   payload,
		CreatedAt: time.Now(),
	}
}
