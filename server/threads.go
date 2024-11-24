package main

import "time"

type Thread struct {
	ThreadID  int       `json:"thread_id"`
	Title     string    `json:"title"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

