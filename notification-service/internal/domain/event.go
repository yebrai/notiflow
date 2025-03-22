package domain

import (
	"github.com/google/uuid"
	"time"
)

type EventType string

const (
	UserRegistered EventType = "USER_REGISTERED"
	OrderCompleted EventType = "ORDER_COMPLETED"
	PaymentFailed  EventType = "PAYMENT_FAILED"
)

type Event struct {
	ID        string    `json:"id"`
	Type      EventType `json:"type"`
	Payload   []byte    `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}

func NewEvent(eventType EventType, payload []byte) Event {
	return Event{
		ID:        uuid.New().String(),
		Type:      eventType,
		Payload:   payload,
		CreatedAt: time.Now(),
	}
}
