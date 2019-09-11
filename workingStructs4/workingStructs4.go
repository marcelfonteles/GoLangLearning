package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func init() {
	fmt.Println("Working with structs, that last one (i hope)")
}

func main() {
	ka := vehicle{
		doors: 5,
		color: "red",
	}
	corolla := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "silver",

		},
		luxury:  true,
	}
	carretaFuracao := truck{
		vehicle:   vehicle{
			doors: 8,
			color: "almost all color of universe",
		},
		fourWheel: true,
	}

	fmt.Println(ka)
	fmt.Println(corolla)
	fmt.Println(carretaFuracao)
	// The variable name was inspired in movie directed by Steven Spielberg, Duel.
	// I will create a anonymous type wich will simbolize the truck of film.
	spilbergTruck := struct {
		name string
		message string
		color string
		doors int
		wheels int
		isClose bool
	}{
		name: "Devil's Truck",
		message: "I'm gonna kill you",
		color: "Black Manta",
		doors: 2,
		wheels: 16,
		isClose: true,
	}

	fmt.Println(spilbergTruck)
}