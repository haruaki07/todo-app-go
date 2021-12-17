package datastore

import (
	"log"
	"net/http"

	"github.com/Xanvial/todo-app-go/model"
	gonanoid "github.com/matoous/go-nanoid/v2"
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

	id, err := gonanoid.New()
	if err != nil {
		panic(err)
	}

	log.Println("ArrayStore | title:", title)
	ms.data[id] = model.TodoData{
		Title: title,
	}
}
