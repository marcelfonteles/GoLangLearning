package main

import (
	"github.com/kataras/iris"
	_ "github.com/kataras/iris/middleware/logger"
	_ "github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.Default()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello World</h1>")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello World!"})
	})

	app.Run(iris.Addr(":3000"))
}
