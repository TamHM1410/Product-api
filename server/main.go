package main

import (
	"fmt"
	"go-backend-api/internal/global"
	"go-backend-api/internal/initialize"
	"go-backend-api/internal/routers"
	"log"
	"net"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Load config  start HTTP server
	config := initialize.LoadConfig()

	db, err := initialize.InitDB()

	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	router := routers.InitRouter(db)

	lis, err := net.Listen("tcp", ":12051")

	if err != nil {
		log.Fatalf("Failed to listen:  %v", err)
	}

	s := grpc.NewServer(
		grpc.MaxConcurrentStreams(100),
	)
	///grpc handler

	//connect to rabbitmq
	conn, err := amqp.Dial("amqp://root:password@localhost:5672/")

	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)

	}

	initialize.InitialProductGRPC(s, db, conn)

	fmt.Println("config:", strconv.Itoa(config.GetInt("port")), config.GetString("host"))
	// Start gRPC server in background

	// pb.RegisterSimpleServer(s, &GrpcServer{})
	reflection.Register(s)

	log.Println("===========================================")
	log.Println("gRPC Server is running on port:  12051")
	log.Println("===========================================")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	router.Run(global.Config.Host + ":" + strconv.Itoa(global.Config.Port))
}
