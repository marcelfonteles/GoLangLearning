package geometry2

import (
	"fmt"
	rectangle "geometry/rectangle"
	"log"
)

var rectLen, rectWidth float64 = -6, 8

func init() {
	fmt.Println("The main package was initialized")
	if rectLen <= 0 {
		log.Fatal("The length is less or equal than zero. That's not possible")
	}
	if rectWidth <= 0 {}
	log.Fatal("The width is less or equal than zero. That's not possible")
}

func Geometry2() {
	// Do the same thing that geometry.go does
	rectangle.Area(rectLen, rectWidth)
}