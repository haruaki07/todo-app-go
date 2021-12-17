package datastore

import (
	"log"
	"net/http"

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

func (ms *MapStore) CreateTodo(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")

	log.Println("ArrayStore | title:", title)
	ms.data[title] = model.TodoData{
		Title: title,
	}
}

func (ms *MapStore) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	delete(ms.data, title)
}
