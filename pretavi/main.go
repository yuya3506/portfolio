package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートの設定
	e.GET("/", hello)
	e.GET("/pretavi", pretavi)
	e.GET("/pretavi/trip-plans", createPlan)

	// HTTPサーバーの設定
	srv := &http.Server{
		Addr:    ":1323",
		Handler: e,
	}

	// シャットダウン用のチャンネル
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// サーバー起動
	go func() {
		e.Logger.Fatal(srv.ListenAndServe())
	}()

	// シグナルを待機
	<-stop

	// シャットダウン処理
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		e.Logger.Fatal("Server Shutdown:", err)
	}

	e.Logger.Info("Server exited properly")
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func pretavi(c echo.Context) error {
	return c.String(http.StatusOK, "プリタビへようこそ！")
}

func createPlan(c echo.Context) error {
	return c.String(http.StatusOK, "ここでプランを作成します")
}
