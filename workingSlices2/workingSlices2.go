package main

import "fmt"

func main() {
	var slice1 []int
	slice1 = append(slice1, 1)
	fmt.Println(slice1)
	slice1 = append(slice1, 2)
	fmt.Println(slice1)
	slice1 = append(slice1, slice1...)
	fmt.Println(slice1)

	var slice2 = []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(slice2)

	for index, value := range slice2 {
		fmt.Println("index:",index, ", value:", value)
	}

	slice2 = append(slice2[:4], slice2[6:]...)
	fmt.Println(slice2)
}