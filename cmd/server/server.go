package main

import (
	"log"
	"net"

	"github.com/vaguiiinho/grpc/pb"
	"github.com/vaguiiinho/grpc/services"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Cloud not connect: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Cloud not serve: %v", err)
	}
}
