package notification

import "context"

type Repository interface {
	Save(ctx context.Context, notification *Notification) error
	FindById(ctx context.Context, notificationId string) (*Notification, error)
}
