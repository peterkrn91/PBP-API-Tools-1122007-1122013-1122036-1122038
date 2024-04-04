package controllers

import (
	"context"
	"fmt"
	"log"

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

	// Example: Add tasks to the list
	addTask(redisClient, "Task 1")
	addTask(redisClient, "Task 2")
	addTask(redisClient, "Task 3")

	// Example: Retrieve all tasks
	tasks, err := getAllTasks(redisClient)
	if err != nil {
		log.Fatal("Error retrieving tasks:", err)
	}
	fmt.Println("All tasks:")
	for _, task := range tasks {
		fmt.Println(task)
	}
}

func addTask(client *redis.Client, task string) {
	client.LPush(ctx, "tasks", task)
}

func getAllTasks(client *redis.Client) ([]string, error) {
	return client.LRange(ctx, "tasks", 0, -1).Result()
}
