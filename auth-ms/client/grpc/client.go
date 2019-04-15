package grpc

import (
	"fmt"

	pb "github.com/IhorBondartsov/test-k8s/auth-ms/pb"

	"google.golang.org/grpc"
)

func NewDataManagerServiceClient(address string) (pb.GreeterClient, error) {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %s", err)
	}

	return pb.NewGreeterClient(conn), nil
}
