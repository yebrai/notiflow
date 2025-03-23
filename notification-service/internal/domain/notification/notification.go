package notification

import (
	"time"

	"github.com/google/uuid"
)

type NotificationStatus string

const (
	Pending   NotificationStatus = "PENDING"
	Delivered NotificationStatus = "DELIVERED"
	Failed    NotificationStatus = "FAILED"
)

type NotificationType string

const (
	Email NotificationType = "EMAIL"
	SMS   NotificationType = "SMS"
	Push  NotificationType = "PUSH"
)

type Notification struct {
	ID        uuid.UUID          `json:"id"`
	EventID   uuid.UUID          `json:"event_id"`
	Type      NotificationType   `json:"type"`
	Recipient string             `json:"recipient"`
	Content   string             `json:"content"`
	Status    NotificationStatus `json:"status"`
	SentAt    *time.Time         `json:"sent_at,omitempty"`
	CreatedAt time.Time          `json:"created_at"`
}
