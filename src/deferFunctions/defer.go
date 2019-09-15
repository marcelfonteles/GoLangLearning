package main

import "fmt"

func init() {
	fmt.Println("Learning about defer.")
}

func main() {
	// Defer executes when the program ends
	defer foo()
	bar()
}

func foo() {
	fmt.Println("this is foo.")
}
func bar() {
	fmt.Println("this is bar.")
}