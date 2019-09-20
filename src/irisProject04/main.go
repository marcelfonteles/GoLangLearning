package main

import(
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

func main() {
	app := iris.New()

	formHandler := hero.Handler(form)
	app.Get("/", formHandler)

	formPostHandler := hero.Handler(formpost)
	app.Post("/formpost", formPostHandler)

	orm, err := xorm.NewEngine("sqlite3", "./test.db")	app.Run(iris.Addr(":3000"))
}

func form() string {
	return `
<html>
	<head>
		<title>Iris Learning 04</title>
	</head>
	<body>
		<form action='/formpost' method='post'>
			<p>First Name: <input type='text' name='fname'></p>
			<p>Last Name: <input type='text' name='lname'></p>
			<input type='submit' value='Submit'>
		</form>
	</body>
</html>
`
}

func formpost(ctx iris.Context) string {
	fname := ctx.PostValue("fname")
	lname := ctx.PostValue("lname")
	fmt.Println(fname)
	fmt.Println(lname)
	return `
<html>
	<head>
		<title>Iris Learning 04</title>
	</head>
	<body>
		<h1>Form Post Page</h1>
		<p>Hello, ` + fname + ` ` + lname + `</p>
	</body>
</html>
`
}