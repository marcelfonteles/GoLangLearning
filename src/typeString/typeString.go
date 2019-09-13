package main

import (
	"bufio"
	"fmt"
	"os"
)

func init() {
	fmt.Println("The main program was initialized.")
}

func main() {
	fmt.Print("Digite algo: ")
	text := readUserInput()
	fmt.Print("User input: ", text)

	fmt.Print("Digite algo novamente: ")
	text = readUserInput()
	fmt.Print("User input again: ", text)


}

func readUserInput() (string) {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	return text
}