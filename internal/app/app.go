package app

import (
	"context"
	"log"

	"github.com/IDarar/hub/pkg/logger"
	"github.com/IDarar/notifications-service/internal/config"
	"github.com/IDarar/notifications-service/internal/repository"
	"github.com/IDarar/notifications-service/internal/server"
	"github.com/IDarar/notifications-service/internal/service"
	v1 "github.com/IDarar/notifications-service/internal/transport/grpc/v1"
	"github.com/IDarar/notifications-service/pkg/database/mongodb"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		logger.Error(err)
		return
	}

	logger.Info(cfg)

	mongoClient, err := mongodb.NewClient(cfg.Mongo.URI, cfg.Mongo.User, cfg.Mongo.Password)
	if err != nil {
		logger.Error(err)
		return
	}
	ctx := context.Background()

	db := mongoClient.Database(cfg.Mongo.Name)

	defer db.Client().Disconnect(ctx)
	repos := repository.NewRepositories(db)

	srvcs := service.NewServices(service.Deps{
		Repos: repos,
	})
	notsServer := v1.NewNotificationsServer(srvcs)

	log.Fatal(server.RunServer(cfg, notsServer))

}
