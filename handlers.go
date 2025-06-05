package main

import (
	"net/http"
	"encoding/json"

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

	var taskList []Task
	for _, task := range tasks {
		taskList = append(taskList, task)
	}
	jsonResponse(w, taskList, http.StatusOK)
}