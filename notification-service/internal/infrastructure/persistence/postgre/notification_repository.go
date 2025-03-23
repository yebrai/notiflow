package postgre

import (
	"context"
	"errors"
	"notification-service/internal/domain/notification"
	"notification-service/internal/infrastructure/persistence/postgre/dao"

	"gorm.io/gorm"
)

const notificationsTable = "notifications"

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) notification.Repository {
	return &NotificationRepository{db: db}
}

func (n NotificationRepository) Save(ctx context.Context, notification *notification.Notification) error {
	var notificationDAO dao.NotificationDAO
	notificationDAO.FromDomain(*notification)

	tx := n.db.WithContext(ctx).Create(&notificationDAO)
	return tx.Error
}

func (n NotificationRepository) FindById(ctx context.Context, notificationId string) (*notification.Notification, error) {
	var notificationDAO dao.NotificationDAO

	tx := n.db.WithContext(ctx).Where("id = ?", notificationId).First(&notificationDAO)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("notification not found")
		}
		return nil, tx.Error
	}

	response, err := notificationDAO.ToDomain()
	if err != nil {
		return nil, err
	}
	return response, nil
}
