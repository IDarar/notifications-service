package repository

import (
	"context"

	"github.com/IDarar/hub/pkg/logger"
	"github.com/IDarar/notifications-service/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationsRepo struct {
	db *mongo.Collection
}

func NewNotificationsRepo(db *mongo.Database) *NotificationsRepo {
	return &NotificationsRepo{db: db.Collection(notificationsCollection)}
}

func (r *NotificationsRepo) Create(ctx context.Context, not domain.Notification) error {
	res, err := r.db.InsertOne(ctx, not)
	logger.Info(res.InsertedID)
	return err
}
func (r *NotificationsRepo) Delete() {}
func (r *NotificationsRepo) GetByUserID(uID int) ([]domain.Notification, error) {
	return nil, nil
}
