package main

import (
	"log"
	"net/http"
	"sync"
	"os"
	"encoding/json"

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

func saveTasks() {
	var taskList []Task
	for _, t := range tasks {
		taskList = append(taskList, t)
	}
	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		log.Printf("Failed to encode tasks: %v", err)
		return
	}
	if err := os.WriteFile(dataFile, data, 0644); err != nil {
		log.Printf("Failed to write data file: %v", err)
	}
}

func jsonResponse(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func main() {
	// Router
	router := httprouter.New()

	router.GET("/health_check", healthCheck)
	router.POST("/api/tasks", createTask)

	// Server
	server := &http.Server {
		Addr: ":8080",
		Handler: router,
	}

	log.Println("✅ Server is running at http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}

