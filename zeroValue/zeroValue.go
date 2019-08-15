package main

import "fmt"

func main() {
	// Zero value
	var str1 string
	var int1 int
	var flo1 float32

	fmt.Println("Printing the zero values:")
	fmt.Println("String:", str1)
	fmt.Println("Integer:", int1)
	fmt.Println("Float:", flo1)

	flo1 = 0.1
	fmt.Println("Float:", flo1)
}
