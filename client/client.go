package main

import (
	"log"

	"github.com/IDarar/hub/pkg/logger"
	"github.com/IDarar/notifications-service/internal/config"
	"github.com/IDarar/notifications-service/pkg/tlscredentials"
	"google.golang.org/grpc"
)

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
}
