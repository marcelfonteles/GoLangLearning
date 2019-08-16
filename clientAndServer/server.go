package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Lauching server...")
	c1 := make(chan net.Conn)

	go connection1(c1)

	var conn net.Conn

	// run forever
	for {
		conn = <- c1
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))

		newMessage := strings.ToUpper(message)

		conn.Write([]byte(newMessage+"\n"))
		fmt.Println(time.Since(start))

	}

}

func connection1(c chan net.Conn) {
	// Listen on all interfaces
	ln, _ := net.Listen("tcp", ":3000")

	// accept connection on port
	for {
		conn, _ := ln.Accept()
		c <- conn
	}

}
