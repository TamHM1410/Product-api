package global

import "github.com/redis/go-redis/v9"

type GlobalConfigType struct {
	AppName string
	Version string
	Port    int
	Host    string
}

type DatabaseConfigType struct {
	User            string
	Password        string
	DBName          string
	Host            string
	Port            int
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int // in minutes
}
type RedisConfigType struct {
	Host     string
	Port     int
	Password string
	DB       int
}

var Config GlobalConfigType
var DBConfig DatabaseConfigType
var RedisConfig RedisConfigType
var Rdb *redis.Client
