package main

import (
	"log"

	"github.com/kouxi08/Artfolio/config"
	"github.com/kouxi08/Artfolio/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//jsonファイルのデコード
	config, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	//サーバ起動
	server(config)
}

func server(config *config.Config) {
	//インスタンス作成
	e := echo.New()

	//ミドルウェア設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("config", config)
			return next(c)
		}
	})
	//レコード追加処理にルーティング
	e.GET("/addrecode", handler.AddDnsHandler)
	// e.GET("/", showDnsHandler)

	e.Logger.Fatal(e.Start(":8088"))
}
