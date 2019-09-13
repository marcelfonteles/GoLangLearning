package main

import (
	"exercises/code"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the master program.")
	fmt.Println("In here will be called many packages of exercises.")
	fmt.Println("Type the number of exercise you want to see: ")
	code.Exercise01()
	code.Exercise02()
	code.Exercise03()

	fmt.Println(code.ReadInput())

	fmt.Println("End of Program. Successfully!")
}