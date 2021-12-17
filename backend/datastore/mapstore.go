package datastore

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Xanvial/todo-app-go/model"
	"github.com/gorilla/mux"
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

	log.Println("ArrayStore | title:", title)
	ms.data[title] = model.TodoData{
		Title: title,
	}
}

func (ms *MapStore) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	status, err := strconv.ParseBool(r.FormValue("status"))
	if err != nil {
		panic(err)
	}

	log.Println("ArrayStore | title:", title)
	log.Println("ArrayStore | status:", status)

	if todo, ok := ms.data[title]; ok {
		todo.Status = status
		ms.data[title] = todo
	}
}

func (ms *MapStore) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	log.Println("ArrayStore | title:", title)

	delete(ms.data, title)
}
