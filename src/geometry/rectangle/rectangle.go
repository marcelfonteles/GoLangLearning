// rectangle.go
package rectangle

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("This package are for Rectangles forms.")
}

func main() {

}

func Area(len float64, wid float64) float64 {
	return len * wid
}

func Diagonal(len float64, wid float64) float64 {
	return math.Sqrt((len * len) + (wid * wid))
}