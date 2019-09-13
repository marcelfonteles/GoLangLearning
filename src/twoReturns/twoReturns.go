package main

import "fmt"

func main() {
	p, s := size()
	fmt.Println(p, s)

}

func size() (primeiro int, segundo int) {
	primeiro = 1
	segundo = 2
	return
}