package main

import (
	"context"
	"log"

	"github.com/IDarar/hub/pkg/logger"
	"github.com/IDarar/notifications-service/internal/config"
	"github.com/IDarar/notifications-service/internal/domain/pb"
	"github.com/IDarar/notifications-service/pkg/tlscredentials"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
)

var ctx = context.Background()

func main() {
	cfg, err := config.Init("configs/main")
	if err != nil {
		logger.Error(err)
		return
	}
	tlsCredentials, err := tlscredentials.LoadTLSCredentialsClient(cfg)
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	transportOption := grpc.WithTransportCredentials(tlsCredentials)

	conn, err := grpc.Dial("0.0.0.0:"+cfg.GRPC.Port, transportOption)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	defer conn.Close()
	c := pb.NewNotificationsServiceClient(conn)

	ts := &timestamp.Timestamp{Seconds: 10, Nanos: 10}

	res := &pb.Notification{To: "1241", From: "124", Time: ts, Where: "RTE", Content: "Hi", Type: "forum not"}

	r, err := c.NotificationCreate(ctx, res)

	if err != nil {
		logger.Error("could not perform req ", err)
		return
	}
	logger.Info(r)
} /*
{
	"to": "Hello",
	"from": "Hello",
	"time": {
	  "seconds": 20,
	  "nanos": 10
	},
	"where": "Hello",
	"content": "Hello",
	"type": "Hello"
  }*/
