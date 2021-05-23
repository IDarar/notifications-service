package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/IDarar/hub/pkg/logger"
	"github.com/IDarar/notifications-service/internal/domain"
	"github.com/IDarar/notifications-service/internal/repository"
	"github.com/IDarar/notifications-service/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//Notification types
const (
	forum       = "forum"
	friend      = "friend"
	firstFriend = "firstFriend"
	ban         = "ban"
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
	notifications, err := s.repo.GetForUser(uID, offset, notType)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	protoNots := []*pb.Notification{}

	if len(notifications) == 0 {
		return protoNots, nil
	}
	for i := range notifications {
		protoNots = append(protoNots, &pb.Notification{
			UserId:  int32(notifications[i].UserID),
			From:    int32(notifications[i].From),
			Topic:   notifications[i].Topic,
			Text:    notifications[i].Text,
			Time:    timestamppb.New(notifications[i].CreatedAt),
			Type:    notifications[i].Type,
			IsRead:  notifications[i].IsRead,
			Content: notifications[i].Content,
		})
	}

	return protoNots, nil
}
