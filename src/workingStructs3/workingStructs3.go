package main

import "fmt"

type person struct {
	firstName string
	lastName string
	favoriteFlavors []string
}
func init() {
	fmt.Println("Learning about structs")
}

func main() {
	person1 := person{
		firstName:       "Marcel",
		lastName:        "Vieira",
		favoriteFlavors: []string{"Vanilla", "Chocolate", "Mango"},
	}

	fmt.Println(person1)
	fmt.Println(person1.firstName, person1.lastName)
	fmt.Println(person1.favoriteFlavors)

	// Part Two
	mapPerson := map[string]person{}

	mapPerson[person1.lastName] = person1
	fmt.Println(mapPerson)

	person2 := person{
		firstName:       "Larissa",
		lastName:        "Fonteles",
		favoriteFlavors: []string{"Ninho"},
	}
	mapPerson[person2.lastName] = person2
	fmt.Println(mapPerson)


	for key, value := range mapPerson {
		fmt.Println(key,":",value)
	}




}