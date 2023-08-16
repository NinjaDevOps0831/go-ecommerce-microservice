package api

import (
	"fmt"
	"net"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/config"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/pb"
	"google.golang.org/grpc"
)

type ServiceServer struct {
	gs   *grpc.Server
	Lis  net.Listener
	Port string
}

func NewgrpcServer(c *config.Config, service pb.AuthServiceServer) (*ServiceServer, error) {
	grpcserver := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcserver, service)
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		return nil, err
	}
	return &ServiceServer{
		gs:   grpcserver,
		Lis:  lis,
		Port: c.Port,
	}, nil
}

func (s *ServiceServer) Start() error {
	fmt.Println("Authentication service listening on:", s.Port)
	return s.gs.Serve(s.Lis)
}
