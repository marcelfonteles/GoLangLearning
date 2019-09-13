package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type caixa struct {
	valor int
	point *caixa
}

func init() {
	fmt.Println("Welcome to linked list")
}

func main() {
	var primeraCaixa *caixa
	var atual *caixa
	var proximo *caixa

	first := true

	option := "start"

	for option != "exit" {
		leitor := bufio.NewReader(os.Stdin)
		fmt.Print("Insert a value: ")
		numberString, _ := leitor.ReadString('\n')
		numberString = strings.TrimSuffix(numberString, "\n")

		numberInt, err := strconv.Atoi(numberString)
		if err != nil {
			fmt.Println("Invalid number.")
			break
		}

		if first {
			caixa := caixa{numberInt, nil}
			primeraCaixa = &caixa
			atual = &caixa
			proximo = &caixa
		} else {
			caixa := caixa{numberInt, nil}
			proximo = &caixa
			atual.point = proximo
			atual = proximo
		}
		first = false

		var ponteiro *caixa
		ponteiro = primeraCaixa

		for ponteiro != nil {
			fmt.Print(ponteiro.valor, " -> ")
			ponteiro = ponteiro.point
		}
		fmt.Printf("End \n")

	}
}