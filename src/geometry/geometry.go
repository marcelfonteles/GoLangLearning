// geometry.go
package main

import (
	"fmt"
	rectangle "geometry/rectangle"
	geometry2 "geometry/geometry2"
)

func init() {
	fmt.Println("Geometrical shape properties.")
}

func main() {
	var rectLen, rectWidth float64 = 6,7
	// Area
	fmt.Printf("Area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	// Diagonal
	fmt.Printf("Diagonal of rectangle %.2f\n", rectangle.Diagonal(rectLen, rectWidth))

	// This package was only inserted for test, does not have any other use
	fmt.Println(geometry2.Geometry2)
}
