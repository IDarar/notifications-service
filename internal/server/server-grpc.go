package server

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	"github.com/IDarar/notifications-service/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
}

var (
	crtFile = "server.crt"
	keyFile = "server.key"
	caFile  = "ca.crt"
)

func InitServer(cfg *config.Config, cert tls.Certificate) {
	certificate, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append ca certificate")
	}
	opts := []grpc.ServerOption{
		grpc.Creds(
			credentials.NewTLS(&tls.Config{
				ClientAuth:   tls.RequireAndVerifyClientCert,
				Certificates: []tls.Certificate{certificate},
				ClientCAs:    certPool,
			},
			)),
	}
	s := grpc.NewServer(opts...)
	proto.RegisterProductInfoServer(s, &server{})
	lis, err := net.Listen("tcp", cfg.GRPC.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
