package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello World!</h1> <br> <a href='/ping'>Pong page</a>")
	})

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.HTML("<h1>Pong</h1")
	})

	app.Handle("GET", "/vieira", func(ctx iris.Context) {
		ctx.HTML("<h1>Marcel Rocha Fonteles Vieira</h1>")
	})

	app.Run(iris.Addr(":3000"), iris.WithoutServerError(iris.ErrServerClosed))
}