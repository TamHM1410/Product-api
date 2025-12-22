package initialize

import (
	"go-backend-api/internal/global"
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	v := viper.New()

	v.AddConfigPath("./config")
	v.SetConfigName("local")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("❌ Cannot read config: %v", err)
	}
	/// add to global config
	global.Config.AppName = v.GetString("server.name")
	global.Config.Version = v.GetString("server.version")
	global.Config.Port = v.GetInt("server.port")
	global.Config.Host = v.GetString("server.host")

	global.DBConfig.User = v.GetString("database.user")
	global.DBConfig.Password = v.GetString("database.password")
	global.DBConfig.DBName = v.GetString("database.dbname")
	global.DBConfig.Host = v.GetString("database.host")
	global.DBConfig.Port = v.GetInt("database.port")
	global.DBConfig.SSLMode = v.GetString("database.sslmode")
	global.DBConfig.MaxOpenConns = v.GetInt("database.max_open_conns")
	global.DBConfig.MaxIdleConns = v.GetInt("database.max_idle_conns")
	global.DBConfig.ConnMaxLifetime = v.GetInt("database.conn_max_lifetime")

	global.RedisConfig.Host = v.GetString("redis.host")
	global.RedisConfig.Port = v.GetInt("redis.port")
	global.RedisConfig.Password = v.GetString("redis.password")
	global.RedisConfig.DB = v.GetInt("redis.db")

	/// redis config can be added here similarly
	log.Println("✅ Config loaded successfully")

	return v
}
