package server

import (
	"fmt"
	"net"

	"github.com/IDarar/hub/pkg/logger"
	"github.com/IDarar/notifications-service/internal/config"
	"github.com/IDarar/notifications-service/pb"
	"github.com/IDarar/notifications-service/pkg/tlscredentials"
	"google.golang.org/grpc"
)

func RunServer(cfg *config.Config, notServer pb.NotificationsServiceServer) error {

	tlsCredentials, err := tlscredentials.LoadTLSCredentialsServer(cfg)
	if err != nil {
		return fmt.Errorf("cannot load TLS credentials: %w", err)
	}

	opts := []grpc.ServerOption{}
	opts = append(opts, grpc.Creds(tlsCredentials))

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterNotificationsServiceServer(grpcServer, notServer)

	listener, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		return fmt.Errorf("cannot run tcp server: %w", err)
	}
	logger.Info("Start GRPC server at ", listener.Addr().String())
	return grpcServer.Serve(listener)
}
