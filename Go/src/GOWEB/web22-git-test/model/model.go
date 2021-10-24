package model

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Ide       string    `json:"ide"`
	Pass      string    `json:"pass"`
	Birth     string    `json:"birth"`
	Sex       string    `json:"sex"`
	Name      string    `json:"name"`
	Phone2    string    `json:"phone2"`
	Email     string    `json:"email"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type DBHandler interface {
	GetTodos(sessionId string) []*Todo
	AddTodo(sessionId string, ide string, pass string, birth string, sex string, name string, phone2 string, email string) *Todo
	RemoveTodo(id int) bool
	CompleteTodo(id int, complete bool) bool
	Close()
}

func NewDBHandler(filepath string) DBHandler {
	//handler = newMemoryHandler()
	return newSqliteHandler(filepath)
}
