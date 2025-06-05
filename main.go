package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)


type Task struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
	Done    bool   `json:"done"`
}

var (
	tasks   = make(map[int]Task)
	nextID  = 1
	taskMux = sync.Mutex{}
	dataFile = "tasks.json"
)



func main() {
	loadTasks()

	// Router
	router := httprouter.New()

	router.GET("/health_check", healthCheck)
	router.POST("/api/tasks", createTask)
	router.GET("/api/tasks", listTasks)
	router.PUT("/api/tasks/:id", updateTask)
	router.DELETE("/api/tasks/:id", deleteTask)

	// Server
	server := &http.Server {
		Addr: ":8080",
		Handler: router,
	}

	log.Println("✅ Server is running at http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}

