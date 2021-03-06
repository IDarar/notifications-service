package repository

import (
	"context"

	"github.com/IDarar/notifications-service/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type Notifications interface {
	Create(ctx context.Context, not domain.Notification) error
	Delete()
	GetForUser(uID int, offset int, notType string) ([]*domain.Notification, error)
}
type Repositories struct {
	Notifications Notifications
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Notifications: NewNotificationsRepo(db),
	}
}
