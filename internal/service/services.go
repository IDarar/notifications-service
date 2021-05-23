package service

import (
	"context"

	"github.com/IDarar/notifications-service/internal/domain"
	"github.com/IDarar/notifications-service/internal/repository"
	"github.com/IDarar/notifications-service/pb"
)

type Notifications interface {
	Create(ctx context.Context, not domain.Notification) error
	Delete()
	GetForUser(uID int, offset int, notType string) ([]*pb.Notification, error)
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
