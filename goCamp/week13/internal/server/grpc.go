package server

import (
	"google.golang.org/grpc"
	v1 "week13/api/user"
	"week13/internal/service"
)

func NewGRPCService(service *service.EmpService) *grpc.Server {
	var opts []grpc.ServerOption

	srv := grpc.NewServer(opts...)
	v1.RegisterEmployeeServer(srv, service)

	return srv
}
