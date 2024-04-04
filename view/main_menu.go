package view

import (
	m "PBP-API-Tools-1122007-1122013-1122036-1122038/model"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var tasks = make(map[int]m.Task)
var ctx = context.Background()

func AddTask(id int, title string, details string, dueDate time.Time, startTask time.Time, email string) {
	task := m.Task{ID: id, Title: title, Details: details, DueDate: dueDate, StartTask: startTask, Email: email}
	tasks[id] = task
	fmt.Println("Task added successfully")
}

func Menu(redisClient *redis.Client) {
	printTask(redisClient)

	// menu.Menu()
	var choice int
	fmt.Println("\n1. Add Task")
	fmt.Println("2. Delete Task")
	fmt.Println("3. Exit")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		var title string
		fmt.Println("Enter task title:")
		fmt.Scan(&title)
		addTask(redisClient, title)

	case 2:
		var id int
		fmt.Println("Enter task id to delete:")
		fmt.Scan(&id)
		// DeleteTask(id)
		deleteTask(redisClient, id)
	case 3:
		return
	default:
		fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
	}

	Menu(redisClient)

}

func printTask(client *redis.Client) {
	tasks, err := getAllTasks(client)
	if err != nil {
		log.Fatal("Error retrieving tasks:", err)
	}
	fmt.Println("All tasks:")
	for id, task := range tasks {
		fmt.Printf("ID: %d, Task: %s\n", id, task)
	}
}

func addTask(client *redis.Client, task string) {
	err := client.LPush(ctx, "tasks", task).Err()
	if err != nil {
		log.Fatal("Error adding task:", err)
	}
	return
}

func deleteTask(client *redis.Client, id int) {
	// Get all tasks
	tasks, err := getAllTasks(client)
	if err != nil {
		log.Fatal("Error retrieving tasks:", err)
	}

	// Check if the ID is valid
	if id < 0 || id >= len(tasks) {
		fmt.Println("Invalid ID")
		return
	}

	// Remove the task from the list
	err = client.LRem(ctx, "tasks", 0, tasks[id]).Err()
	if err != nil {
		log.Fatal("Error deleting task:", err)
	}

	fmt.Println("Task deleted successfully")
}

func getAllTasks(client *redis.Client) ([]string, error) {
	return client.LRange(ctx, "tasks", 0, -1).Result()
}
