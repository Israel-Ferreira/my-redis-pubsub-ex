package config

import (
	"os"

	dotenv "github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func ConfigRedis() *redis.Client {

	dotenv.Load()

	envHost := os.Getenv("REDIS_HOST")

	rdb := redis.NewClient(&redis.Options{
		Addr:     envHost,
		Password: "",
		DB:       0,
	})

	return rdb
}
