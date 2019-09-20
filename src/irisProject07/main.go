package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/kataras/iris/middleware/basicauth"
	"time"
)

func newApp() *iris.Application {
	app := iris.New()

	// *** Some Configuration for new App ***
	//Register Views
	app.RegisterView(iris.HTML("./app/views", ".html"))

	// Auth
	authConfig := basicauth.Config{
		Users: map[string]string{ "marcelvieira":"password", "marcelfonteles":"password"},
		Realm: "Authorization Required",
		Expires: time.Duration(30) * time.Minute,
	}
	authentication := basicauth.New(authConfig)
	needAuth := app.Party("/", authentication)
	{
		needAuth.Get("/", root)
	}
	app.Use(authentication)
	return app
}

func main() {
	app := newApp()

	rootHandler := hero.Handler(root)
	app.Get("/", rootHandler)

	noAuthHandler := hero.Handler(noAuth)
	app.Get("/noAuth", noAuthHandler)

	app.Run(iris.Addr(":3000"))
}

func root(ctx iris.Context) {
	title := "Basic Auth"
	ctx.ViewData("title", title)
	ctx.View("index.html")
}

func noAuth(ctx iris.Context) {
	ctx.ViewData("title", "Basic Auth")
	ctx.View("index.html")
}

func h(ctx iris.Context) {
	username, password, _ := ctx.Request().BasicAuth()
	ctx.Writef("%s %s:%s", ctx.Path(), username, password)
}