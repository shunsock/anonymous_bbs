package main

import (
	"context" // OpenAI APIで使用
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"              // Echoフレームワーク
	openai "github.com/sashabaranov/go-openai" // OpenAIライブラリ
)

// コメント作成リクエスト用の構造体
type CreateCommentRequest struct {
	Comment            string `json:"comment" validate:"required"`
	CommenterIPAddress string `json:"commenter_ip_address" validate:"required"`
	ThreadID           int    `json:"thread_id" validate:"required"`
	Username           string `json:"username"`
}

// コメント用の構造体
type Comment struct {
	ID                 int    `json:"id"`
	Comment            string `json:"comment"`
	CommenterIPAddress string `json:"commenter_ip_address"`
	ThreadID           int    `json:"thread_id"`
	Username           string `json:"username"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}

func CreateCommentHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// リクエストを解析
		var req CreateCommentRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// OpenAI APIでコメントを変形
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Println("OPENAI_API_KEY is not set")
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "OpenAI API key not configured"})
		}

		client := openai.NewClient(apiKey)

		// ChatGPT APIリクエスト
		resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
			Model: "gpt-4",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "あなたは2chのなんJ民です。",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "あなたは2ちゃんねらーです。次のコメントを2ちゃんねるっぽくしてください: " + req.Comment,
				},
			},
		})
		if err != nil {
			log.Printf("OpenAI API error: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to transform comment"})
		}

		transformedComment := resp.Choices[0].Message.Content

		// クエリを準備
		query := `
			INSERT INTO bbs_comments (comment, commenter_ip_address, thread_id, username, created_at, updated_at)
			VALUES ($1, $2, $3, $4, NOW(), NOW())
			RETURNING id, comment, commenter_ip_address, thread_id, username, created_at, updated_at
		`
		log.Println("Executing query: ", query)

		var comment Comment
		err = db.QueryRow(query, transformedComment, req.CommenterIPAddress, req.ThreadID, req.Username).Scan(
			&comment.ID,
			&comment.Comment,
			&comment.CommenterIPAddress,
			&comment.ThreadID,
			&comment.Username,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)

		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create comment"})
		}

		// 作成したコメントを返す
		return c.JSON(http.StatusCreated, comment)
	}
}
