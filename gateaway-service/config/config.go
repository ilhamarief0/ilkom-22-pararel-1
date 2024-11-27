package config

import (
	"log"

	"google.golang.org/grpc"
)

func ConnectToGRPCServer(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	return conn
}
