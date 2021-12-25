package model

type TodoData struct {
	ID     string `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Status bool   `json:"status" db:"status"`
}
