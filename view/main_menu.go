package view

import (
	m "PBP-API-Tools-1122007-1122013-1122036-1122038/model"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/redis/go-redis/v9"
	gomail "gopkg.in/gomail.v2"
)

var tasks = make(map[int]m.Task)
var password = "uyakuyaoye"
var ctx = context.Background()

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
		SendMail("New Task Added", "Task Tittle : "+title)
		gocrons()
		Menu(redisClient)
	case 2:
		var id int
		fmt.Println("Enter task id to delete:")
		fmt.Scan(&id)
		deleteTask(redisClient, id)
		Menu(redisClient)
	case 3:
		return
	default:
		fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
	}

}

func printTask(client *redis.Client) {
	tasks, err := getAllTasks(client)
	if err != nil {
		log.Fatal("Error retrieving tasks:", err)
	}
	fmt.Println("All tasks:")
	for id, task := range tasks {
		fmt.Printf("ID: %d, Task: %s\n", id, task)

		go func() {
			SendMail("Task List", fmt.Sprintf("ID: %d, Task: %s\n", id, task))
		}()
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
	tasks, err := getAllTasks(client)
	if err != nil {
		log.Fatal("Error retrieving tasks:", err)
	}

	if id < 0 || id >= len(tasks) {
		fmt.Println("Invalid ID")
		return
	}

	err = client.LRem(ctx, "tasks", 0, tasks[id]).Err()
	if err != nil {
		log.Fatal("Error deleting task:", err)
	}

	fmt.Println("Task deleted successfully")
}

func getAllTasks(client *redis.Client) ([]string, error) {
	tasks, err := client.LRange(ctx, "tasks", 0, -1).Result()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func SendMail(subject string, body string) {
	abc := gomail.NewMessage()

	var email = "if-22007@students.ithb.ac.id"
	abc.SetHeader("From", email)
	abc.SetHeader("To", "waluyajuang330@gmail.com")
	abc.SetHeader("Subject", subject)
	abc.SetBody("text/html", body)

	a := gomail.NewDialer("smtp.gmail.com", 587, email, password)

	if err := a.DialAndSend(abc); err != nil {
		fmt.Println(err)
		panic(err)
	}

}

func gocrons() {
	local, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		local = time.UTC
	}
	s := gocron.NewScheduler(local)
	startTime := time.Now()
	s.Every(5).Minute().Do(func() {
		elapsed := time.Since(startTime)
		minuteCounter := 5 - int(elapsed.Minutes())
		SendMail("Reminder", fmt.Sprintf("%d Minutes left!", minuteCounter))
	})
	s.StartBlocking()
	time.Sleep(1 * time.Minute)
	s.Clear()
}
