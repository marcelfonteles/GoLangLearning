package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 32,
		"Miss Moneypenny": 27,
		"Marcel": 24,
	}

	fmt.Println(m)
	fmt.Println(m["Marcel"])

	if value, ok := m["Barnabas"]; ok {
		fmt.Println(value)
	}

	if value, ok := m["Marcel"]; ok {
		fmt.Println(value)
	}
}