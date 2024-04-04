package main

import (
	"PBP-API-Tools-1122007-1122013-1122036-1122038/controllers"
)

func main() {
	// Initialize Redis client
	controllers.InitializeRedisClient()

	// 	// Start HTTP server and handle login/connect routes
	// 	//controllers.Token()

	// Start HTTP server in a separate goroutine
	// go func() {
	// 	router := mux.NewRouter()
	// 	router.HandleFunc("/login", controllers.CheckUserLogin).Methods("GET")
	// 	// router.HandleFunc("/tasks", ).Methods("GET")
	// 	fmt.Println("Connected to port 8888")
	// 	log.Println("Connected to port 8888")
	// 	log.Fatal(http.ListenAndServe(":8888", router))
	// }()

	// // Start goCRON scheduler
	// s := gocron.NewScheduler()
	// s.Every(10).Second().Do(func() {

	// 	result, err := queryDatabase()
	// 	if err != nil {
	// 		log.Println("Error querying database:", err)
	// 		return
	// 	}
	// 	fmt.Println("Query result:", result)
	// 	for _, task := range result {
	// 		go func(t m.Task) {
	// 				if time.Now().After(t.DueDate.Add(-10*time.Minute)) && time.Now().Before(t.DueDate) && t.Notified == 0 {
	// 					//fmt.Printf("Task " + t.Title + " is due in 10 minutes!")

	// 					sendMail(t.Email, t.Title, t.Details, t.DueDate)
	// 					controllers.UpdateNotified(t.ID)

	// 					// Query ke db update task notified++
	// 					return
	// 				} else if time.Now().After(t.DueDate.Add(-5*time.Minute)) && time.Now().Before(t.DueDate) && t.Notified == 1 {
	// 					//fmt.Printf("Task " + t.Title + " is due in 5 minutes!")

	// 					sendMail(t.Email, t.Title, t.Details, t.DueDate)
	// 					controllers.UpdateNotified(t.ID)

	// 					return
	// 				}
	// 		}(task)
	// 	}

	// })
	// <-s.Start() // This line will block indefinitely, so it's typically not used in a real application
}
