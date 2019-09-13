package main

import "fmt"

type people struct {
	lastName string
	firstName string
	age int
}

type secretAgent struct {
	people
	ltk bool
}


func init() {
	fmt.Println("Struct was initialized!")
}

func main() {
	person1 := people{
		lastName: "Vieira",
		firstName: "Marcel",
		age: 24,
	}

	fmt.Println("Complete Struct:", person1)
	fmt.Println("First name:", person1.firstName)
	fmt.Println("Last name:", person1.lastName)
	fmt.Println("Age:", person1.age, "years")

	sAgent := secretAgent{
		people: people{
			lastName: "Bond",
			firstName: "James",
			age: 32,
		},
		ltk:    true,
	}

	fmt.Println(sAgent.lastName + ",",sAgent.firstName)
	if sAgent.ltk {
		fmt.Println("I have licence kill")
		fmt.Println("What your back, Motherfucker")
	}








}