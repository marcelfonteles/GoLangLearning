package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the ip address: ")
	endereco, _ := reader.ReadString('\n')
	endereco = strings.TrimSuffix(endereco, "\n")
	fmt.Print("Enter the port: ")
	porta, _ := reader.ReadString('\n')
	porta = strings.TrimSuffix(porta, "\n")
	conn, _ := net.Dial("tcp", endereco+":"+porta)

	for {

		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprint(conn, text + "\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)

		//conn, _ = net.Dial("tcp", "127.0.0.1:3000")
		conn, _ = net.Dial("tcp", endereco+":"+porta)
	}
}
