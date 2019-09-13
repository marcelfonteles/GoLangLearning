package main

import (
	"bufio"
	"log"
	"net"
	"fmt"
	"os"
	"strings"
)

func main() {
	hostName := "localhost"
	portNum := "6000"

	service := hostName + ":" + portNum

	RemoteAddr, err := net.ResolveUDPAddr("udp", service)

	conn, err := net.DialUDP("udp", nil, RemoteAddr)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Established connection to %s \n", service)
	log.Printf("Remote UDP address : %s \n", conn.RemoteAddr().String())
	log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())

	defer conn.Close()


	message := []byte("Hello UDP server!")
	_, err = conn.Write(message)

	if err != nil {
		log.Println(err)
	}

	buffer := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buffer)

	fmt.Println("UDP Server : ", addr)
	fmt.Println("Received from UDP server : ", string(buffer[:n]))

	for {
		fmt.Print("mensagem: ")
		leitor := bufio.NewReader(os.Stdin)
		texto, _ := leitor.ReadString('\n')
		texto = strings.TrimSuffix(texto, "\n")
		if texto == "\\close" {
			fmt.Println("Closing connection...")
			conn.Close()
			break
		}
		message = []byte(texto)

		conn.Write(message)

		if texto == "\\uptime" {
			n, _, _ := conn.ReadFromUDP(buffer)
			fmt.Println("Running time: ", string(buffer[:n]))
		} else if texto == "\\reqnum" {
			n, _, _ := conn.ReadFromUDP(buffer)
			fmt.Println("Requests: ", string(buffer[:n]))
		}
	}


}