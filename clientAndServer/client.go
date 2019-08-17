package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func init() {
	fmt.Println("### Welcome ###")
	fmt.Println("# First a little help")
	fmt.Println("# if you don't enter an ip and port, the program will use this combination: localhost:3000")
	fmt.Println("# Server function: capitalize received string")
	fmt.Println("# Commands:")
	fmt.Println("  \\close    - close conection")
	fmt.Println("  \\reqnum   - server running time")
	fmt.Println("  \\actconn  - total of clients connected")

	fmt.Println("##############")
}

func main() {
	endereco := "192.168.100.14"
	porta := "3000"

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the ip address: ")
	enderecoDigitado, _ := reader.ReadString('\n')
	if enderecoDigitado != "\n"{
		endereco = strings.TrimSuffix(enderecoDigitado, "\n")
	}


	fmt.Print("Enter the port: ")
	portaDigitado, _ := reader.ReadString('\n')
	if portaDigitado != "\n" {
		porta = strings.TrimSuffix(portaDigitado, "\n")
	}

	conn, _ := net.Dial("tcp", endereco+":"+porta)

	for {
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		fmt.Fprint(conn, text + "\n")

		if text == "\\close\n" {
			break
		}

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)


	}
	conn.Close()
	fmt.Println("Gracefully shutdown client.")
}


