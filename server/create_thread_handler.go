package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type CreateThreadRequest struct {
	Title    string `json:"title" validate:"required"`
	Username string `json:"username" validate:"required"`
}

func CreateThreadHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse the request body
		var req CreateThreadRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// Insert thread into the database
		query := `
			INSERT INTO bbs_threads (title, username, created_at, updated_at)
			VALUES ($1, $2, NOW(), NOW())
			RETURNING thread_id, title, username, created_at, updated_at
		`
		log.Printf("Executing query: %s with Title: %s, Username: %s", query, req.Title, req.Username)

		var thread Thread
		err := db.QueryRow(query, req.Title, req.Username).Scan(
			&thread.ThreadID,
			&thread.Title,
			&thread.Username,
			&thread.CreatedAt,
			&thread.UpdatedAt,
		)

		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create thread"})
		}

		// Return the created thread
		return c.JSON(http.StatusCreated, thread)
	}
}
