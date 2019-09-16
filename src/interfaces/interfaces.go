package main

import "fmt"

func init() {
	fmt.Println("I will learning about Interfaces in Golang.")
}

type car struct {
	name string
	color string
	yearOfFabrication int
	numberOfDoors int
}

type ferrari struct {
	car
	horsePower int
	rearMotor bool
	classic bool
}

type sedan struct {
	car
	luxury bool
}
// Empty interfaces: all the types implements a interface with zero methods
type allCars interface {
	// This interface can assume all the types that have startEngine method
	startEngine() bool
	describe()
}

// Methods for car type
func (c car) startEngine() bool {
	fmt.Println("Brrrr....")
	return true
}
func (c car) describe() {
	fmt.Printf("This car is a %v, painted with %v color and has %v doors \n", c.name, c.color, c.numberOfDoors)
}

// Methods for sedan type
func (s sedan) startEngine() bool {
	fmt.Println("Does not make any sound.")
	return true
}
func (s sedan) racing() {
	fmt.Println("Brrrr.... 0-60 in 2.5 seconds")
}

func allKindsOfTypes(i interface{}) {
	fmt.Printf("This function received %T type with value: %v \n", i, i)
}

func main() {
	var a allCars
	car1 := car{
		name:              "Ka",
		color:             "Red",
		yearOfFabrication: 2018,
		numberOfDoors:     5,
	}
	car1.startEngine()
	car1.describe()

	sedan1 := sedan{
		car:    car{
			name: "Corolla",
			color: "silver",
			yearOfFabrication: 2019,
			numberOfDoors: 5,
		},
		luxury: true,
	}
	sedan1.startEngine()
	sedan1.describe()
	sedan1.racing()

	a = car1
	fmt.Printf("type of interface: %T\n", a)

	a = sedan1
	fmt.Printf("type of interface: %T\n", a)

	// Underlying value
	var cars allCars
	cars = car1
	fmt.Println(cars)
	fmt.Println(cars.(car))

	// Some data types
	var i int = 4
	var f float64 = 3.5
	var xi []int = []int{1,2,3,4,5}
	var xs []string = []string{"Marcel", "Rocha", "Fonteles", "Vieira"}
	var xxi [][]int = [][]int{{1,2,3,4,5}, {6,7,8,9,10}}
	allKindsOfTypes(i)
	allKindsOfTypes(f)
	allKindsOfTypes(xi)
	allKindsOfTypes(xs)
	allKindsOfTypes(xxi)
	// A empty interface implements all kind of data types
}