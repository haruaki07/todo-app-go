package datastore

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Xanvial/todo-app-go/model"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type DBStore struct {
	db *sql.DB
}

func NewDBStore() *DBStore {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		model.DBHost, model.DBPort, model.DBUser, model.DBPassword, model.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB Successfully connected!")

	return &DBStore{
		db: db,
	}
}

func (ds *DBStore) GetCompleted(w http.ResponseWriter, r *http.Request) {
	var completed []model.TodoData

	query := `
		SELECT id, title, status
		FROM todo
		WHERE status = true
	`

	rows, err := ds.db.Query(query)
	if err != nil {
		log.Println("error on getting todo:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	defer rows.Close()

	for rows.Next() {
		var data model.TodoData
		if err := rows.Scan(&data.ID, &data.Title, &data.Status); err != nil {
			log.Println("error on getting todo:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		completed = append(completed, data)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(completed)
}

func (ds *DBStore) GetIncomplete(w http.ResponseWriter, r *http.Request) {
	var incomplete []model.TodoData

	query := `
		SELECT id, title, status
		FROM todo
		WHERE status = false
	`

	rows, err := ds.db.Query(query)
	if err != nil {
		log.Println("error on getting todo:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	defer rows.Close()

	for rows.Next() {
		var data model.TodoData
		if err := rows.Scan(&data.ID, &data.Title, &data.Status); err != nil {
			log.Println("error on getting todo:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		incomplete = append(incomplete, data)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(incomplete)
}

func (ds *DBStore) CreateTodo(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")

	query := `
		INSERT INTO todo (title, status)
		VALUES ($1, false)
		RETURNING id, title, status
	`

	log.Println("DBStore | title:", title)

	var todo model.TodoData

	err := ds.db.QueryRow(query, &title).Scan(&todo.ID, &todo.Title, &todo.Status)
	if err != nil {
		log.Println("error on insert todo:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (ds *DBStore) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	status, _ := strconv.ParseBool(r.FormValue("status"))

	query := `
		UPDATE todo
		SET status = $1
		WHERE id = $2
	`
	log.Println("DBStore | id:", id)
	log.Println("DBStore | status:", status)

	_, err := ds.db.Exec(query, &status, &id)
	if err != nil {
		log.Println("error on update todo:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	io.WriteString(w, "success")
}

func (ds *DBStore) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	query := `
		DELETE FROM todo
		WHERE id = $1
	`

	log.Println("DBStore | id:", id)

	_, err := ds.db.Exec(query, &id)
	if err != nil {
		log.Println("error on delete todo:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	io.WriteString(w, "success")
}
