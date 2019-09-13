package main

import "fmt"

func init() {

}

func main() {
	m := map[string]int{
		"Marcel": 24,
		"Armando": 28,
		"Leticia": 23,
	}

	fmt.Println(m)
	fmt.Println(m["Marcel"])

	for key, value := range m {
		fmt.Println(key, value)
	}

	if value, ok := m["Glauton"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("Glauton não está presente, ele está tentando em outra oportunidade.")
	}
}