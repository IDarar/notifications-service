package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/IDarar/notifications-service/internal/domain"
	"github.com/IDarar/notifications-service/internal/repository"
	"github.com/IDarar/notifications-service/pb"
)

//Notification types
const (
	forum       = "forum"
	friend      = "friend"
	firstFriend = "firstFriend"
	ban         = "ban"

	amount = 20 //get nots by one request
)

type NotificationsService struct {
	repo repository.Notifications
}

func NewNotificationsService(repo repository.Notifications) *NotificationsService {
	return &NotificationsService{
		repo: repo,
	}

}

//Checks type and formats content accodring to it
func (s *NotificationsService) Create(ctx context.Context, not domain.Notification) error {
	if not.Type == forum {
		not.Text = fmt.Sprintf(domain.ForumNotification, not.Content, not.Topic)
	} else if not.Type == friend {
		not.Text = fmt.Sprintf(domain.FirstFriendRequest)
	} else if not.Type == firstFriend {
		not.Text = firstFriend
	} else if not.Type == ban {
		//TODO, after impelemting full ban system, add ban time field to message in pb
		not.Text = fmt.Sprintf(domain.Ban, "2 hours", not.Content)
	} else {
		return errors.New("invalid notification type")
	}

	return s.repo.Create(ctx, not)
}

func (s *NotificationsService) Delete() {}

//Change seen to unseen and vice versa
func (s *NotificationsService) MarkAs() {}

func (s *NotificationsService) GetForUser(uID int, offset int, notType string) ([]*pb.Notification, error) {

	return nil, nil
}
