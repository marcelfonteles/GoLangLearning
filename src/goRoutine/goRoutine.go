package main

import (
	"fmt"
	"time"
)

func main() {
	go count()
	time.Sleep(time.Millisecond * 2)
	fmt.Println("Hello World with go routines!")
	time.Sleep(time.Millisecond * 5)
}

func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 1)
	}
}