package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGRPCServer() {
	lis, err := net.Listen("tcp", ":12051")
	if err != nil {
		log.Fatalf("Failed to listen:  %v", err)
	}

	s := grpc.NewServer(
		grpc.MaxConcurrentStreams(100),
	)

	// pb.RegisterSimpleServer(s, &GrpcServer{})
	reflection.Register(s)

	log.Println("===========================================")
	log.Println("gRPC Server is running on port:  12051")
	log.Println("===========================================")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
