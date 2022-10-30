package models

import "time"

type Content struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Article struct {
	ID string `json:"id"`
	Content
	Author    Person     `json:"author"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
