package models

import "time"

type Content struct {
	Title string //`json:"title"`
	Body  string //`json:"body"`
}

type Article struct {
	ID int //`json"id"`
	Content
	Author    Person
	CreatedAt *time.Time
}
