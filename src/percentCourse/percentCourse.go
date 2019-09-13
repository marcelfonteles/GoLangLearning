package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func init() {
	_, width := terminalSize()
	organizeTerminal(width, "#")
	fmt.Println("Calculate percentage of Udemy Course")
	organizeTerminal(width, "#")
}

func main() {
	_, width := terminalSize()
	option := "init"
	for option != "exit" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Digite o total de itens do curso: ")
		total, _ := reader.ReadString('\n')
		total = strings.TrimSuffix(total, "\n")
		option = total
		totalInt, _ := strconv.Atoi(total)
		for option != "exit" && total != "exit" {
			organizeTerminal(width, "-")
			fmt.Printf("0 - Para ver os checkpoints do curso \n1 - Para escolher um percentual \n2 - Para escolher um item \n")
			fmt.Print("Opção: ")
			option, _ = reader.ReadString('\n')
			option = strings.TrimSuffix(option, "\n")
			organizeTerminal(width, "-")
			if option == "0" {
				checkpoint := math.Ceil(float64(totalInt) * 0.25)
				fmt.Printf("25%%: %v itens conclúidos. \n", checkpoint)

				checkpoint = math.Ceil(float64(totalInt) * 0.50)
				fmt.Printf("50%%: %v itens conclúidos. \n", checkpoint)

				checkpoint = math.Ceil(float64(totalInt) * 0.75)
				fmt.Printf("75%%: %v itens conclúidos. \n", checkpoint)
			} else if option == "1" {
				fmt.Print("Digite uma porcetagem a ser atingida: ")
				percent, _ := reader.ReadString('\n')
				percent = strings.TrimSuffix(percent, "\n")
				percentInt, _ := strconv.Atoi(percent)
				objetivo := math.Ceil( float64(totalInt) * (float64(percentInt)/100) )
				fmt.Printf("É necessário concluir %v itens para atingir %v%% do curso.\n", objetivo, percentInt)

			} else if option == "2" {
				fmt.Print("Digite a quantidade de itens concluídos: ")
				items, _ := reader.ReadString('\n')
				items = strings.TrimSuffix(items, "\n")
				itemsInt, _ := strconv.Atoi(items)
				percent := (float64(itemsInt) / float64(totalInt)) * 100.00

				fmt.Printf("Ao concluir %v items você terá concluído %.2f%% do curso.\n", itemsInt, percent)

			} else if option != "exit" {
				fmt.Println("Opção Inválida")
			}
		}
	}
}

func terminalSize() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	indexHeight := strings.Index(string(out), " ")
	indexWidth := strings.Index(string(out), "\n")

	heightInt, _ := strconv.Atoi(string(out[0:indexHeight]))
	widthInt, _  := strconv.Atoi(string(out[indexHeight+1:indexWidth]))

	return heightInt, widthInt
}

func organizeTerminal(width int, char string) {
	for i := 1; i <= width; i++ {
		fmt.Print(char)
	}
	fmt.Print("\n")
}

func centerTerminalText(text string, width int) {
	sizeText := len(text)
	fmt.Println(sizeText)
}
























