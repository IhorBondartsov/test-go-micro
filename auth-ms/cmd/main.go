package main

import (
	"fmt"
	"os"

	"github.com/IhorBondartsov/test-k8s/auth-ms/config"
	"github.com/IhorBondartsov/test-k8s/auth-ms/server/grpc"
	authService "github.com/IhorBondartsov/test-k8s/auth-ms/service"
)

func main() {
	// Read cfg in .env file
	cfg, err := config.ResolveConfig()
	if err != nil {
		fmt.Printf("Cant read .env file. Err: %v", err)
		os.Exit(1)
	}

	// Create grpc server with business logic
	as := authService.NewAuthenticator()

	// Create and run grpc server
	grpcServer := grpc.NewGreeterGRPCServer(grpc.NewGreeterService(as))
	err = grpcServer.Serve(cfg.GRPCAddress)
	if err != nil {
		panic(err)
	}
}
