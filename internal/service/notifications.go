package service

import (
	"context"

	"github.com/IDarar/notifications-service/internal/domain"
	"github.com/IDarar/notifications-service/internal/repository"
)

type NotificationsService struct {
	repo repository.Notifications
}

func NewNotificationsService(repo repository.Notifications) *NotificationsService {
	return &NotificationsService{
		repo: repo,
	}

}
func (s *NotificationsService) Create(ctx context.Context, not domain.Notification) error {
	return nil
}
func (s *NotificationsService) Delete()
func (s *NotificationsService) GetByUserID(uID int) ([]domain.Notification, error)
