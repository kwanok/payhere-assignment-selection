package config

import (
	"github.com/go-redis/redis/v8"
	"os"
)

var JwtRedis *redis.Client

func InitRedis() {
	//Initializing redis
	jwtDsn := os.Getenv("REDIS_JWT_DSN")
	if len(jwtDsn) == 0 {
		jwtDsn = "localhost:6379"
	}
	JwtRedis = redis.NewClient(&redis.Options{
		Addr: jwtDsn, //redis port
	})
}
