package main

import (
	"encoding/json"
	"log"
	"os"
	"net/http"
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