package main

import (
	"net/http"
	"encoding/json"
	"strconv"
	"fmt"

	"github.com/julienschmidt/httprouter"
)


func healthCheck(w http.ResponseWriter, r *http.Request, _  httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func createTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var t Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	taskMux.Lock()
	t.ID = nextID
	nextID++
	tasks[t.ID] = t
	saveTasks()
	taskMux.Unlock()

	jsonResponse(w, t, http.StatusCreated)
}

func listTasks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	taskMux.Lock()
	defer taskMux.Unlock()

	// Get "done" query param (optional)
	doneParam := r.URL.Query().Get("done")

	var filtered []Task
	if doneParam == "" {
		// No filter, return all tasks
		for _, t := range tasks {
			filtered = append(filtered, t)
		}
	} else {
		// Parse doneParam to bool
		done, err := strconv.ParseBool(doneParam)
		if err != nil {
			http.Error(w, "Invalid done parameter, must be true or false", http.StatusBadRequest)
			return
		}
		for _, t := range tasks {
			if t.Done == done {
				filtered = append(filtered, t)
			}
		}
	}

	jsonResponse(w, filtered, http.StatusOK)
}

func updateTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	taskMux.Lock()
	defer taskMux.Unlock()

	task, exists := tasks[id]
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	task.Done = true
	tasks[id] = task
	saveTasks()

	jsonResponse(w, task, http.StatusOK)
}

func deleteTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	taskMux.Lock()
	defer taskMux.Unlock()

	if _, exists := tasks[id]; !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	delete(tasks, id)
	saveTasks()

	response := map[string]string{
		"message": fmt.Sprintf("Task with ID %d has been deleted", id),
	}
	jsonResponse(w, response, http.StatusOK)
}
