package grpc

import (
	"context"

	pb "github.com/IhorBondartsov/test-k8s/auth-ms/pb"
	"github.com/IhorBondartsov/test-k8s/auth-ms/service"
)

type greeterService struct {
	auth service.Authenticator
}

func (g *greeterService) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	answer := g.auth.Greet(req.GetName())
	return &pb.HelloResponse{
		Greeting: answer,
	}, nil
}

func NewGreeterService(auth service.Authenticator) pb.GreeterServer {
	return &greeterService{
		auth: auth,
	}
}
