package main

import "fmt"

type car struct {
	name string
	wheels int
	horsePower int
	doors int
	color string
}

type bus struct {
	car
	public bool
}

func init() {
	fmt.Println("This is Working With Structs 2")
}

func main() {

	car1 := car{
		name:       "Corvette",
		wheels:     4,
		horsePower: 350,
		doors:      2,
		color:      "red",
	}

	fmt.Println("I'm gonna talk about my car.")
	fmt.Println("Here some specification:")
	fmt.Printf("Name: %v ", car1.name)
	fmt.Printf("Color: %v ", car1.color)
	fmt.Printf("Number of doors: %v ", car1.doors)
	fmt.Printf("Number of wheels: %v ", car1.wheels)
	fmt.Printf("\n")

	bus1 := bus{
		car: car {
			name:"",
			wheels: 0,
			horsePower: 350,
			doors: 6,
			color: "white",
		},
		public: true,
	}

	fmt.Println(bus1)
}