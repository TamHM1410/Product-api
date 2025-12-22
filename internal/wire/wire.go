// internal/wire/wire.go
//go:build wireinject
// +build wireinject

package wire

import (
	"go-backend-api/cmd/initialize"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AppDependencies struct {
	DB                *gorm.DB
	Redis             *redis.Client
	ProductController *controller.ProductController
}

// Wire function
func InjectDependencies() *AppDependencies {
	wire.Build(
		initialize.LoadDbPostgres,
		initialize.LoadRedis,
		wire.Struct(new(AppDependencies), "*"), // inject tất cả field
	)
	return &AppDependencies{} // dummy return, Wire generate code sẽ replace
}
