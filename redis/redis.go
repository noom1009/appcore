package redis

import (
    "github.com/go-redis/redis/v8"
    "context"
    "log"
)

var client *redis.Client

func InitRedis(addr string) {
    client = redis.NewClient(&redis.Options{
        Addr: addr,
    })

    _, err := client.Ping(context.Background()).Result()
    if err != nil {
        log.Fatalf("Redis connection error: %v", err)
    }
    log.Println("✅ Redis connection established")
}

func GetClient() *redis.Client {
    return client
}
