package main

import "fmt"

func init() {
	fmt.Println("####################")
	fmt.Println("### Bit Shifting ###")
	fmt.Println("####################")
}

func main() {
	pow2 := 2
	var i uint
	for i = 1; i <= 10; i++ {
		fmt.Printf("%d\t", pow2)
		pow2 = 2 << i
	}

	fmt.Printf("\n####################\n")
}
/*
	10
	100
	1000
	10000
*/