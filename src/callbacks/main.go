package main

import "fmt"

func init() {
	fmt.Println("Learning about callbacks - a function that has a function in your parameters")
}

func main() {
	xi := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}

	xiReturn := oddNumbers(xi...)
	fmt.Println(xiReturn)

	xiReturn2 := sumSlice(oddNumbers, xi...)
	fmt.Println(xiReturn2)

	xiReturn3 := oddNumberSlice(sum, xi...)
	fmt.Println(xiReturn3)
}

func oddNumbers(xi ...int) []int {
	var xiReturn []int
	for _, v := range xi {
		if v % 2 == 1 {
			xiReturn = append(xiReturn, v)
		}
	}
	return xiReturn
}

func sumSlice(f func(xi ...int) []int, xi ...int) int {
	funcReturn := f(xi...)
	sum := 0
	for _, v := range funcReturn {
		sum += v
	}

	return sum
}

func sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func oddNumberSlice(f func(xi ...int) int, xi ...int) int {
	var xiOdd []int
	for _, v := range xi {
		if v % 2 == 1 {
			xiOdd = append(xiOdd, v)
		}
	}
	sum := f(xiOdd...)
	return sum
}