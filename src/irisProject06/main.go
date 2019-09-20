package main

import (
	"bytes"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/kataras/iris/middleware/basicauth"
	"log"
	"os/exec"
	"strings"
	"time"
)

func newApp() *iris.Application {
	app := iris.New()

	// Some configuration
	app.RegisterView(iris.HTML("./app/views", ".html"))

	authConfig := basicauth.Config{
		Users:   map[string]string{"marcelvieira":"password"},
		Realm:   "Authorization Required",
		Expires: time.Duration(1) * time.Minute,
		OnAsk:   nil,
	}
	authentication := basicauth.New(authConfig)
	app.Use(authentication)

	return app
}

func main() {
	app := newApp()

	rootHandler := hero.Handler(root)
	app.Get("/", rootHandler)

	indexHandler := hero.Handler(index)
	app.Get("/index/html", indexHandler)

	postFormHandler := hero.Handler(postForm)
	app.Post("/api/post", postFormHandler)

	app.Run(iris.Addr(":3000"))
}

func root() string {
	return`
<html>
	<head>
		<title>Hello World</title>		
	</head>
	<body>
		<h1>Hello World!</h1>
		<p><a href='/index/html' class='btn btn-primary'>Some Form</a></p>
	</body>
</html>
`
}

func index(ctx iris.Context) {
	ctx.View("index.html")
}

func postForm(ctx iris.Context) {
	fname := ctx.FormValue("fname")
	lname := ctx.FormValue("lname")
	birthday := ctx.FormValue("birthday")
	age := ctx.FormValue("age")
	sex := ctx.FormValue("radioOptionsSex")

	ctx.Redirect("/index/html")

	ctx.JSON(iris.Map{
		"fname": fname,
		"lname": lname,
		"birthday": birthday,
		"age": age,
		"sex": sex,
	})
}

/* DEPRECATED */
func returnPath() string {
	cmd := exec.Command("pwd")
	cmd.Stdin = strings.NewReader("\n")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSuffix(out.String(),"\n") + "/"
}
