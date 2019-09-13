package main

import (
	"bufio"
	converting "converting/packages"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func init() {
	fmt.Println("### Programa de conversão ###")
	fmt.Println("### ##################### ###")
}

func main() {
	converting.ToBaseTwo(10)
	reader := bufio.NewReader(os.Stdin)
	option := "init"

	for option != "exit" {
		fmt.Print("Digite para qual base deseja transformar um número na base 10: ")
		text, err := reader.ReadString('\n')
		fmt.Println("--------------------------------------------------------------")
		if err != nil {
			fmt.Println(err)
		}
		option = strings.TrimSuffix(text, "\n")


		if option == "2" {
			baseTwo()
		} else if option == "16" {
			baseHex()
		} else if option != "exit" {
			fmt.Println("A conversão para essa base ainda não foi implementada. Por favor, selecione outra base.")
			fmt.Println("--------------------------------------------------------------")
		}



	}
	fmt.Println("### Programa encerrado com sucesso! ###")
	fmt.Println("### ############################### ###")
}

func baseTwo()  {
	fmt.Println("Convertendo para base 2.")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("0 - algoritmo do Go / 1 - Algoritmo do Marcel: ")
	option, _ := reader.ReadString('\n')
	option = strings.TrimSuffix(option, "\n")
	optionInt, _ := strconv.Atoi(option)
	fmt.Print("Digite um número na base 10: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	number, _ := strconv.Atoi(text)
	if optionInt == 0 {
		converting.ToBaseTwo(number)
		finalNumber := converting.ToBaseTwo(number)
		fmt.Printf("O número %d convertido para base 2 é %d. \n", number, finalNumber)
		fmt.Println("--------------------------------------------------------------")

	} else if optionInt == 1 {
		fmt.Printf("O número %d convertido para base 2 é %b. \n", number, number)
		fmt.Println("--------------------------------------------------------------")
	} else {
		fmt.Println("Opção inválida")
	}
}

func baseHex() {
	fmt.Println("Convertendo para base 16.")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite um número na base 10: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	number, _ := strconv.Atoi(text)

	fmt.Printf("O número %d convertido para base 16 é %#X. \n", number, number)
	fmt.Println("--------------------------------------------------------------")


}
















