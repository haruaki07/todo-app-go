package datastore

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Xanvial/todo-app-go/model"
	"github.com/gorilla/mux"
	nanoid "github.com/matoous/go-nanoid/v2"
)

type MapStore struct {
	data map[string]model.TodoData
}

func NewMapStore() *MapStore {
	newData := make(map[string]model.TodoData, 0)

	return &MapStore{
		data: newData,
	}
}

func (ms *MapStore) GetCompleted(w http.ResponseWriter, r *http.Request) {
	completed := make([]model.TodoData, 0)
	for _, v := range ms.data {
		if v.Status {
			completed = append(completed, v)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(completed)
}

func (ms *MapStore) GetIncomplete(w http.ResponseWriter, r *http.Request) {
	incompleted := make([]model.TodoData, 0)
	for _, v := range ms.data {
		if !v.Status {
			incompleted = append(incompleted, v)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(incompleted)
}

func (ms *MapStore) CreateTodo(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")

	id, _ := nanoid.New()

	todo := model.TodoData{
		ID:     id,
		Title:  title,
		Status: false,
	}

	log.Println("MapStore | id:", id)
	log.Println("MapStore | title:", title)

	ms.data[id] = todo

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (ms *MapStore) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	status, err := strconv.ParseBool(r.FormValue("status"))
	if err != nil {
		panic(err)
	}

	log.Println("MapStore | id:", id)
	log.Println("MapStore | status:", status)

	if todo, ok := ms.data[id]; ok {
		todo.Status = status
		ms.data[id] = todo
	}
}

func (ms *MapStore) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Println("MapStore | id:", id)

	delete(ms.data, id)
}
