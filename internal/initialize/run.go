package initialize

import (
	"go-backend-api/internal/routers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() (*viper.Viper, *gin.Engine) {
	//load config
	config := LoadConfig()

	db, err := InitDB()

	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	router := routers.InitRouter(db)

	return config, router
}
