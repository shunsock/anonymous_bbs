package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	// Connect to the database
	db, err := sql.Open("postgres", "postgres://root:password@postgres_db:5432/bbs?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create Echo instance
	e := echo.New()

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})
	e.POST("/threads", CreateThreadHandler(db))

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
