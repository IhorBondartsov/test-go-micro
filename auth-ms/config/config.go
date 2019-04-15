package config

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)



type AuthMSCfg struct{
	GRPCAddress string `required:"true"`
}

func ResolveConfig() (*AuthMSCfg, error) {
	var (
		grpcAddress string
		envFile     string
	)

	flag.StringVar(&grpcAddress, "grpc-address", "", "Address to run server on")
	flag.StringVar(&envFile, "env", "", "External environment file")
	flag.Parse()

	// Load env. from file if provided
	if envFile != "" {
		if err := godotenv.Load(envFile); err != nil {
			return nil, fmt.Errorf("failed to load environment from [%s], %s", envFile, err)
		}
	}

	config := &AuthMSCfg{}
	if err := envconfig.Process("AUTH_MS", config); err != nil {
		return nil, err
	}

	return config, nil
}