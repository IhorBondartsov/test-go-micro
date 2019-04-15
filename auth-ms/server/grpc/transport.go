package grpc

import (
	"fmt"
	"net"

	pb "github.com/IhorBondartsov/test-k8s/auth-ms/pb"
	"google.golang.org/grpc"
)

type greeterServer struct {
	handler *grpc.Server
}

func (rcv *greeterServer) Serve(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to acquire address [%s]: %s", address, err)
	}

	return rcv.handler.Serve(lis)
}

func NewGreeterGRPCServer(server pb.GreeterServer) *greeterServer {
	grpcServer := &greeterServer{handler: grpc.NewServer()}
	pb.RegisterGreeterServer(grpcServer.handler, server)

	return grpcServer
}
