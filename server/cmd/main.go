package main

import (
	"app/internal/db"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// DB接続の初期化
	dbConn := db.NewDBConnection()

	// マイグレーション
	dbConn.AutoMigrate()

	// リポジトリとサービスの初期化

	// ハンドラの初期化

	// ルーティング
	e.GET("/hello", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
