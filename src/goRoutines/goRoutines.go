package main

import (
	"fmt"
	"time"
)

func main() {
	go up()
	go down()
	time.Sleep(time.Millisecond * 50)
	fmt.Println("This is Sparta!")
	time.Sleep(time.Millisecond * 220)
}

func up(){
	for i := 1; i <= 100; i++ {
		fmt.Printf("%v \n", i)
		time.Sleep(time.Millisecond * 1)
	}
}

func down() {
	for i := 100; i >= 0; i-- {
		fmt.Printf("\t %v\n", i)
		time.Sleep(time.Millisecond * 2)

	}
}