package main

import (
	"fmt"
	config "go-backend-api/cmd/cli/viper"
	"go-backend-api/internal/routers"
)

func main() {
	//#region config load
	cfg := config.LoadConfig()

	//start server
	server := routers.NewRouter()
	server.Run(fmt.Sprintf(":%d", cfg.GetInt("server.port")))

}
