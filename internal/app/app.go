package app

import (
	"context"
	"time"

	"github.com/IDarar/hub/pkg/logger"
	"github.com/IDarar/notifications-service/internal/config"
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
	time.Sleep(15 * time.Second)
}
