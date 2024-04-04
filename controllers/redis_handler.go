package controllers

import (
	"context"
	"log"

	main "PBP-API-Tools-1122007-1122013-1122036-1122038/view"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func InitializeRedisClient() {
	// Initialize Redis client (replace with your actual Redis configuration)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis-17679.c1.ap-southeast-1-1.ec2.cloud.redislabs.com:17679", // Redis server address
		Password: "iArExsNJWgqcTQlEail1ae6oBct9o3VR",                              // Redis password
	})

	// Ping the Redis server to check connectivity
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}

	main.Menu(redisClient)
}
