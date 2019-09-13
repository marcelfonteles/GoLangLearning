 package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func init() {
	awesomePrint()
	fmt.Println("Main program")
	awesomePrint()
}
func awesomePrint() {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	//fmt.Println(string(out))
	outString := strings.Fields(string(out))
	//fmt.Println(outString)
	width, _ := strconv.Atoi(outString[1])
	//fmt.Println(width)
	for i := 1; i <= width; i++ {
		fmt.Print("#")
	}
	fmt.Printf("\n")
}

func main() {
	number := 100
	for number = 100; number <= 1000; number++ {
		if number % 4 == 0 {
			fmt.Println(number)
		}
		//if number > 1000 {
		//	break
		//}
	}
}