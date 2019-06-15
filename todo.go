package main

import "time"

type Todo struct {
	ID          int16     `json:"id"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	Due         time.Time `json:"due"`
}

type Todos []Todo
