package initialize

import (
	"go-backend-api/internal/global"
	"go-backend-api/internal/models"
	"log"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=" + global.DBConfig.Host +
		" port=" + strconv.Itoa(global.DBConfig.Port) +
		" user=" + global.DBConfig.User +
		" password=" + global.DBConfig.Password +
		" dbname=" + global.DBConfig.DBName +
		" sslmode=" + global.DBConfig.SSLMode

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get sql.DB: %v", err)
	}
	// AutoMigrate tự động tạo bảng theo struct
	if err := DB.AutoMigrate(models.AllModels...); err != nil {
		log.Fatal(err)
	}

	sqlDB.SetMaxIdleConns(global.DBConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.DBConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(global.DBConfig.ConnMaxLifetime) * time.Minute)

	log.Println("✅ Database connection established")
	return DB, err
}
