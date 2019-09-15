package main

import "fmt"

func init() {
	fmt.Println("I'm learning about methods and interfaces in Golang.")
}

type person struct {
	firstName string
	lastName string
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println(sa.lastName + ",", sa.firstName, sa.lastName)
}

func (p person) speak() {
	fmt.Println("Please, do not kill me, i'm begging for my life!!")
}

type human interface {
	speak()
}

func receiveHuman(h human) {
	fmt.Println("I received this human:", h)
	fmt.Printf("I received this human, with type: %T\n", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			firstName: "James",
			lastName: "Bond",
		},
		ltk:    true,
	}
	sa1.speak()

	// Person p1 does not have access to method speak(), only secret agent has.
	p1 := person{
		firstName: "Marcel",
		lastName:  "Vieira",
	}

	p1.speak()

	receiveHuman(sa1)
	receiveHuman(p1)
}