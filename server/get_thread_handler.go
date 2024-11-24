package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetThreadsHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// SQLクエリを定義
		query := `
			SELECT thread_id, title, username, created_at, updated_at
			FROM bbs_threads
			ORDER BY created_at DESC
		`

		// クエリ実行
		rows, err := db.Query(query)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch threads"})
		}
		defer rows.Close()

		// 結果をスライスに変換
		var threads []Thread
		for rows.Next() {
			var thread Thread
			if err := rows.Scan(&thread.ThreadID, &thread.Title, &thread.Username, &thread.CreatedAt, &thread.UpdatedAt); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to scan thread"})
			}
			threads = append(threads, thread)
		}

		// クエリ結果を返す
		return c.JSON(http.StatusOK, threads)
	}
}
