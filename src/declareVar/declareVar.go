package main

import "fmt"

var globalString string = "Global String"

func main() {
	var str1, str2 string

	str1, str2 = shortDeclarationString()
	fmt.Println(str1, str2)

	var int1, int2 int

	int1, int2 = shortDeclarationInteger()
	fmt.Println(int1, int2)

	fmt.Println(anotherWayToDeclare())

	printGlobalString()

	// Zero value
	var intZero int
	fmt.Print("Unitialized variables assumes zero values: ")
	fmt.Println(intZero)

	// String inside string
	stringInside := `String inside "String"`

	fmt.Printf("%s \n", stringInside)

}

func shortDeclarationString() (string, string) {

	str1 := "String 1"
	str2 := "String 2"

	return str1, str2
}

func shortDeclarationInteger() (int, int) {

	int1 := 33
	int2 := 42

	return int1, int2
}

func anotherWayToDeclare() (string, int) {

	var str1 string = "string 3"
	var int1 int    = 42

	return str1, int1
}

func printGlobalString() {
	fmt.Println(globalString)
}

