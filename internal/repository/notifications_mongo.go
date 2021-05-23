package repository

import (
	"context"
	"errors"

	"github.com/IDarar/hub/pkg/logger"
	"github.com/IDarar/notifications-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	amount      = 20 //get nots by one request
	forum       = "forum"
	friend      = "friend"
	firstFriend = "firstFriend"
	ban         = "ban"
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

func (r *NotificationsRepo) GetForUser(uID int, offset int, notType string) ([]*domain.Notification, error) {
	if notType == "" {
		cur, err := r.db.Aggregate(context.Background(), mongo.Pipeline{
			bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: 124}}}},
			bson.D{{Key: "$skip", Value: 0}},
			bson.D{{Key: "$limit", Value: amount}}})
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		res := []*domain.Notification{}
		err = cur.All(context.Background(), &res)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return res, nil
	} else if notType == friend {
		cur, err := r.db.Aggregate(context.Background(), mongo.Pipeline{
			bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: 124}}}},
			bson.D{{Key: "$match", Value: bson.D{{Key: "type", Value: bson.D{{Key: "$eq", Value: friend}}}}}},
			bson.D{{Key: "$skip", Value: 0}},
			bson.D{{Key: "$limit", Value: amount}}})
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		res := []*domain.Notification{}
		err = cur.All(context.Background(), &res)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return res, nil
	} else if notType == forum {
		cur, err := r.db.Aggregate(context.Background(), mongo.Pipeline{
			bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: 124}}}},
			bson.D{{Key: "$match", Value: bson.D{{Key: "type", Value: bson.D{{Key: "$eq", Value: forum}}}}}},
			bson.D{{Key: "$skip", Value: 0}},
			bson.D{{Key: "$limit", Value: amount}}})
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		res := []*domain.Notification{}
		err = cur.All(context.Background(), &res)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return res, nil
	} else {
		return nil, errors.New("invalid type")
	}
}
