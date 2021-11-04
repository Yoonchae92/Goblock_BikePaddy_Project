package models

type User struct {
	Id        string
	Password  string
	Name      string
	Created   string
	BoardUser []Board
}