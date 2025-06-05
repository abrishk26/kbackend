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

func loadTasks() {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return // file doesn't exist yet
		}
		log.Fatalf("Failed to read data file: %v", err)
	}

	var taskList []Task
	if err := json.Unmarshal(file, &taskList); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	for _, t := range taskList {
		tasks[t.ID] = t
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}
}

func main() {
	loadTasks()

	// Router
	router := httprouter.New()

	router.GET("/health_check", healthCheck)
	router.POST("/api/tasks", createTask)
	router.GET("/api/tasks", listTasks)

	// Server
	server := &http.Server {
		Addr: ":8080",
		Handler: router,
	}

	log.Println("✅ Server is running at http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}

