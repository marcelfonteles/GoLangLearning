package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"strconv"
)

func main() {
	app := iris.New()
	// Root page
	rootHandler := hero.Handler(root)
	app.Get("/", rootHandler)

	usersHandler := hero.Handler(users)
	app.Get("/users/{id: uint}/{name: string}", usersHandler)

	// Querying Parameters
	queryParamsHandler := hero.Handler(params)
	app.Get("/query", queryParamsHandler)

	// Post method
	formHandler := hero.Handler(form)
	app.Post("/form", formHandler)

	// Without hero.Handler
	withoutHeroHandler := hero.Handler(withoutHero)
	app.Get("/without", withoutHeroHandler)

	app.Run(iris.Addr(":3000"))
}

func root() string {
	return `
			<h1>Hello World</h1>
			<p>This is Root Page</p>
			`
}

func users(id uint, name string) string {
	return `
			<h1>Users Page</h1>
			<p>User Id: ` + strconv.Itoa(int(id)) + `</p>
			<p>User Name: ` + name + `</p>
			`
}

func params(ctx iris.Context) string {
	firstName := ctx.URLParamDefault("firstname", "Guest")
	lastName := ctx.URLParam("lastname")
	fmt.Println(firstName, lastName) // this will print on terminal

	return `
			<h1>Querying Params</h1>
			<p>Hello, ` + firstName + ` ` + lastName + `</p>
			`
}

func form(ctx iris.Context) {
	message := ctx.FormValue("message")
	user := ctx.FormValueDefault("nick", "anonymous")

	ctx.JSON(iris.Map{
		"status": "posted",
		"message": message,
		"nick": user,
	})
}

func withoutHero(ctx iris.Context) string {
	return `<h1>This page was created by a function without hero handler</h1>`
}