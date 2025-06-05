# Task Manager API

A simple Task Manager REST API written in Go using [`httprouter`](https://github.com/julienschmidt/httprouter) with JSON file persistence.

---

## Features

- Create, read, update (mark done), and delete tasks
- Tasks stored persistently in a local `tasks.json` file
- Simple and minimal dependencies

---

## Requirements

- Go 1.21 or later

---

## Installation & Running

1. Clone the repository or copy the source files.

2. Initialize and install dependencies:

```bash
go mod tidy
```

3. (Optional) Create an empty tasks.json file in the project directory if it doesn't exist:
```bash
echo "[]" > tasks.json
```
4. Run the server.
```
go run main.go
```

---
## API Endpoints

This application provides the following RESTful API endpoints:

* **`GET /health_check`**: Checks the health and status of the API.
* **`POST /api/tasks`**: Creates a new task. Requires a JSON request body with `title` and `details` fields.
* **`GET /api/tasks`**: Retrieves a list of all tasks. Optionally accepts a query parameter `done` (e.g., `?done=true` or `?done=false`) to filter tasks by their completion status.
* **`PUT /api/tasks/:id`**: Updates an existing task identified by its ID.
* **`DELETE /api/tasks/:id`**: Deletes a task identified by its ID.