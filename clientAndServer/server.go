package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

var runningTime = time.Now()
var activeConn int // zero value = 0
var reqNum int

func main() {

	fmt.Println("Lauching server...")

	// Listen on all interfaces
	ln, _ := net.Listen("tcp", ":3000")

	// run forever
	for {
		// Accept Connection on port
		conn, _ := ln.Accept()
		// Increase number of connections
		activeConn += 1
		fmt.Println("One client has started a new connection")
		// Go routine for multiple conections
		go comunication(conn)

		//fmt.Println(time.Since(start))

	}

}

func comunication(conn net.Conn) {
	message := "Open"
	for message != "\\close\n" {
		message, _ = bufio.NewReader(conn).ReadString('\n')

		if message == "\\close\n" {
			conn.Close()
			break
		}
		if message == "\\actconn\n" {
			fmt.Print("Active connections:", activeConn, "\n")
			reqNum += 1
			conn.Write([]byte(strconv.Itoa(activeConn)+" active connection(s)\n"))
			continue
		} else if message == "\\uptime\n" {
			reqNum += 1
			conn.Write([]byte(time.Since(runningTime).String() + "\n"))
			continue
		} else if message == "\\reqnum\n" {
			reqNum += 1
			conn.Write([]byte(strconv.Itoa(reqNum) + " requests\n"))

		} else {
			fmt.Print("Message Received:", string(message))

			newMessage := strings.ToUpper(message)
			reqNum += 1
			conn.Write([]byte(newMessage+"\n"))
		}



	}
	fmt.Println("One client was disconected")
	activeConn -= 1
}


// DEPRECATED //
/////////////////////////////////////

func connection(c chan net.Conn) {
	// Listen on all interfaces
	ln, _ := net.Listen("tcp", ":3000")

	// accept connection on port
	for {
		conn, _ := ln.Accept()
		c <- conn
	}

}
