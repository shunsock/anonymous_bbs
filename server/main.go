package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	// read the environment variables
	url := os.Getenv("DATABASE_URL")
	log.Println("DATABASE_URL: ", url)
	// Connect to the database
	db, err := sql.Open("postgres", url)
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
	e.GET("/threads", GetThreadsHandler(db))
	e.POST("/threads", CreateThreadHandler(db))
	e.POST("/comments", CreateCommentHandler(db))

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
