package service

import (
	"context"

	"github.com/IDarar/notifications-service/internal/domain"
	"github.com/IDarar/notifications-service/internal/repository"
)

type Notifications interface {
	Create(ctx context.Context, not domain.Notification) error
	Delete()
	GetByUserID(uID int) ([]domain.Notification, error)
}
type Services struct {
	Notifications Notifications
}
type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	notificationService := NewNotificationsService(deps.Repos.Notifications)
	return &Services{
		Notifications: notificationService,
	}
}
