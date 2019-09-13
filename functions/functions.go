package main

import "fmt"

func init() {
	fmt.Println("Again, i'm learning about functions in Golang.")
}
func main() {
	texto, booleano := funcao("Vieira", "Marcel")

	fmt.Print(texto)
	fmt.Println(booleano)
	xi := []int{1,2,3,4,5}
	sum := funcaoVariadica(xi...)

	fmt.Println(sum)


}

func funcao(lastName string, firstName string) (string, bool) {
	returnString := fmt.Sprintln(firstName, lastName)
	returnBool := false
	return returnString, returnBool
}

func funcaoVariadica(x ...int) int {
	fmt.Println(x)
	var sum int
	for _, value := range x {
		sum += value
	}

	return sum

}