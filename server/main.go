package main

import (
	"fmt"
	"go-backend-api/internal/global"
	"go-backend-api/internal/initialize"
	"strconv"
)

func main() {
	// Start gRPC server in background
	go RunGRPCServer()
	
	// Load config  start HTTP server
	config, router := initialize.Run()
	fmt.Println("config:", strconv.Itoa(config.GetInt("port")), config.GetString("host"))
	router.Run(global.Config.Host + ":" + strconv.Itoa(global.Config.Port))
}