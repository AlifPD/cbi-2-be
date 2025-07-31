package handlers

import (
	"cbi-2-be/middleware"
	"cbi-2-be/models"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

var todos = []models.Todo{}
var todoID = 1
var todoMutex = &sync.Mutex{}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	username := middleware.GetUsername(r)

	filtered := []models.Todo{}
	for _, todo := range todos {
		if todo.Owner == username {
			filtered = append(filtered, todo)
		}
	}

	json.NewEncoder(w).Encode(filtered)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	username := middleware.GetUsername(r)

	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	todoMutex.Lock()
	todo.ID = todoID
	todoID++
	todo.Owner = username
	todos = append(todos, todo)
	todoMutex.Unlock()

	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	username := middleware.GetUsername(r)

	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	var input models.Todo
	json.NewDecoder(r.Body).Decode(&input)

	todoMutex.Lock()
	defer todoMutex.Unlock()

	for i, t := range todos {
		if t.ID == id && t.Owner == username {
			todos[i].Title = input.Title
			todos[i].Checked = input.Checked
			json.NewEncoder(w).Encode(todos[i])
			return
		}
	}

	http.NotFound(w, r)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	username := middleware.GetUsername(r)

	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	todoMutex.Lock()
	defer todoMutex.Unlock()

	for i, t := range todos {
		if t.ID == id && t.Owner == username {
			todos = append(todos[:i], todos[i+1:]...)
			w.Write([]byte("deleted"))
			return
		}
	}

	http.NotFound(w, r)
}
