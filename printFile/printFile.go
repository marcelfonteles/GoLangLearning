package main

import "fmt"

func main() {
	fmt.Println("Olá, olá! Como vai essa vida maravilhosa?")

	foo()

	fmt.Println("Another println in main function")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Print(i)
			fmt.Print(" ")
		}
	}
	fmt.Println(" ")

	fmt.Println("-----------------------")

	str1, str2, str3 := bar()

	fmt.Println(str1 + ", " + str2 + ", " + str3)
	fmt.Println(str1, str2, str3)
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() (string, string, string) {
	var str1, str2 string
	str1 = "String 1"
	str2 = "String 2"
	str3 := "String 3, do not declared explicit"

	return str1, str2, str3

}