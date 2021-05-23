package v1

import (
	"context"
	"strconv"

	"github.com/IDarar/hub/pkg/logger"
	"github.com/IDarar/notifications-service/internal/domain"
	"github.com/IDarar/notifications-service/internal/service"
	"github.com/IDarar/notifications-service/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

//Maybe wrap all proto services to common for use multiplexing
type NotificationServer struct {
	services *service.Services
}

func NewNotificationsServer(srvs *service.Services) *NotificationServer {
	return &NotificationServer{
		services: srvs,
	}
}

func (s *NotificationServer) NotificationCreate(ctx context.Context, not *pb.Notification) (*emptypb.Empty, error) {

	notification := &domain.Notification{UserID: int(not.UserId), From: int(not.From), Type: not.Type, Topic: not.Topic, Content: not.Content, CreatedAt: not.Time.AsTime(), IsRead: false}
	logger.Info(notification)
	empty := &emptypb.Empty{}

	if err := s.services.Notifications.Create(ctx, *notification); err != nil {
		return nil, err
	}

	return empty, nil
}

func (s *NotificationServer) NotificationsGet(ctx context.Context, not *pb.ReqNotifications) (*pb.RespNotifications, error) {
	uId, err := idConv(not.User)
	if err != nil {
		return nil, err
	}
	nots, err := s.services.Notifications.GetForUser(uId, int(not.Offset), not.Type)
	if err != nil {
		return nil, err
	}
	pbNots := &pb.RespNotifications{
		Notifications: nots,
	}
	return pbNots, nil
}

func (s *NotificationServer) NotificationsGetUnread(ctx context.Context, in *pb.ReqNotifications) (*pb.UnreadNumResponse, error) {
	return nil, nil
}

func (s *NotificationServer) NotificationMark(ctx context.Context, in *pb.Notification) (*emptypb.Empty, error) {
	return nil, nil
}

func idConv(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		logger.Error("invalid user ID ", err)
		return 0, err
	}
	return i, nil
}
