package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "go-backend-api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server struct implement SimpleServer interface
type grpcServer struct {
	pb. UnimplementedSimpleServer
}

// Implement SayHello - Unary RPC
func (s *grpcServer) SayHello(ctx context. Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received SayHello request from: %s", req.GetName())
	
	message := fmt.Sprintf("Hello %s! Welcome to gRPC server.", req.GetName())
	
	return &pb.HelloReply{
		Message: message,
	}, nil
}

// Implement StreamHello - Server Streaming RPC
func (s *grpcServer) StreamHello(req *pb.HelloRequest, stream pb.Simple_StreamHelloServer) error {
	log.Printf("Received StreamHello request from:  %s", req.GetName())
	
	for i := 1; i <= 5; i++ {
		message := fmt. Sprintf("Hello %s! This is message #%d from stream", req.GetName(), i)
		
		if err := stream. Send(&pb.HelloReply{
			Message: message,
		}); err != nil {
			log.Printf("Error sending stream message: %v", err)
			return err
		}
		
		log.Printf("Sent message #%d to %s", i, req.GetName())
		time.Sleep(1 * time.Second)
	}
	
	log.Printf("Completed streaming to %s", req.GetName())
	return nil
}

func RunGRPCServer() {
	lis, err := net.Listen("tcp", ":12051")
	if err != nil {
		log.Fatalf("Failed to listen:  %v", err)
	}

	s := grpc.NewServer(
		grpc.MaxConcurrentStreams(100),
	)
	
	pb.RegisterSimpleServer(s, &grpcServer{})
	reflection.Register(s)

	log.Println("===========================================")
	log.Println("gRPC Server is running on port:  12051")
	log.Println("===========================================")
	
	if err := s. Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}