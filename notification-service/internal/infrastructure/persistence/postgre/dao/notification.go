package dao

import (
	"fmt"
	"notification-service/internal/domain/notification"
	"time"

	"github.com/google/uuid"
)

type NotificationDAO struct {
	ID        string `gorm:"primaryKey;type:uuid"`
	EventID   string `gorm:"type:uuid;index"`
	Type      string `gorm:"type:varchar(20)"`
	Recipient string `gorm:"type:varchar(255)"`
	Content   string `gorm:"type:text"`
	Status    string `gorm:"type:varchar(20);index"`
	SentAt    *time.Time
	CreatedAt time.Time
}

func (dao *NotificationDAO) FromDomain(notification notification.Notification) {
	dao.ID = notification.ID.String()
	dao.EventID = notification.EventID.String()
	dao.Type = string(notification.Type)
	dao.Content = notification.Content
	dao.Status = string(notification.Status)
	dao.SentAt = notification.SentAt
	dao.CreatedAt = notification.CreatedAt
}

func (dao *NotificationDAO) ToDomain() (*notification.Notification, error) {
	id, err := uuid.Parse(dao.ID)
	if err != nil {
		return nil, fmt.Errorf("error parsing ID: %w", err)
	}

	idEvent, err := uuid.Parse(dao.EventID)
	if err != nil {
		return nil, fmt.Errorf("error parsing EventID: %w", err)
	}

	return &notification.Notification{
		ID:        id,
		EventID:   idEvent,
		Type:      notification.NotificationType(dao.Type),
		Content:   dao.Content,
		Status:    notification.NotificationStatus(dao.Status),
		SentAt:    dao.SentAt,
		CreatedAt: dao.CreatedAt,
	}, nil
}
