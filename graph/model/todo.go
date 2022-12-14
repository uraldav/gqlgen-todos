package model

import "time"

type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Done      bool      `json:"done"`
	UserID    string    `json:"userId"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}
