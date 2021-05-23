package app

import (
	"context"
	"log"

	"github.com/IDarar/hub/pkg/logger"
	"github.com/IDarar/notifications-service/internal/config"
	"github.com/IDarar/notifications-service/internal/domain"
	"github.com/IDarar/notifications-service/internal/repository"
	"github.com/IDarar/notifications-service/internal/server"
	"github.com/IDarar/notifications-service/internal/service"
	v1 "github.com/IDarar/notifications-service/internal/transport/grpc/v1"
	"github.com/IDarar/notifications-service/pkg/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		logger.Error(err)
		return
	}

	logger.Info(cfg.GRPC.ClientCACertFile)

	mongoClient, err := mongodb.NewClient(cfg.Mongo.URI, cfg.Mongo.User, cfg.Mongo.Password)
	if err != nil {
		logger.Error(err)
		return
	}
	ctx := context.Background()

	db := mongoClient.Database(cfg.Mongo.Name)

	defer db.Client().Disconnect(ctx)
	cur, err := db.Collection("notifications").Aggregate(context.Background(), mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: 124}}}},
		bson.D{{Key: "$match", Value: bson.D{{Key: "type", Value: bson.D{{Key: "$eq", Value: "friend"}}}}}},
		bson.D{{Key: "$skip", Value: 0}},
		bson.D{{Key: "$limit", Value: 20}}})
	if err != nil {
		logger.Error(err)
		return
	}

	res := []*domain.Notification{}
	cur.Current.Elements()
	err = cur.All(context.Background(), &res)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(*res[0])

	repos := repository.NewRepositories(db)

	srvcs := service.NewServices(service.Deps{
		Repos: repos,
	})
	notsServer := v1.NewNotificationsServer(srvcs)

	log.Fatal(server.RunServer(cfg, notsServer))

}
