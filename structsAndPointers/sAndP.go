package main

import "fmt"

type box struct {
	value string
	p *box
}

func init() {
	fmt.Println("I will use structs and pointers")
}

func main() {

	firstBox := box{
		value: "Marcel",
	}

	firstBox.p = &firstBox
	fmt.Println(firstBox)
	fmt.Println(*firstBox.p)

	secondBox := box{
		value: "Teste",
		p:     nil,
	}
	firstBox.p = &secondBox

	thirdBox := box{
		value: "Armando",
		p:     nil,
	}
	secondBox.p = &thirdBox

	forthBox := box{
		value: "Letícia",
		p:     nil,
	}
	thirdBox.p = &forthBox

	var pointer *box
	pointer = &firstBox
	fmt.Println("---------------")
	for pointer != nil {
		fmt.Println(*pointer)
		pointer = pointer.p
	}



}