package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	// Echoインスタンスを作成
	e := echo.New()

	// ルートエンドポイントを定義
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	// サーバーを起動
	e.Logger.Fatal(e.Start(":8080"))
}

