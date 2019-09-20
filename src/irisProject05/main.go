package main

import (
	"github.com/kataras/iris"
	_ "github.com/go-xorm/xorm"
	_ "github.com/kataras/iris/hero"
)

func main() {
	app := iris.New()
	// Register views
	app.RegisterView(iris.HTML("/home/marcelvieira/Documentos/GoLangLearning/src/irisProject05/app/views", ".html"))

	//rootHandler := hero.Handler(root)
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	app.Run(iris.Addr(":3000"))
}

func root(ctx iris.Context) {
	ctx.View("index.html")
}