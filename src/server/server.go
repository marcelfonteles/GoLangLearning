package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

var reqNum int
var runningTime = time.Now()

func handleUDPConnection(conn *net.UDPConn) {

	// here is where you want to do stuff like read or write to client

	buffer := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buffer)
	reqNum += 1
	fmt.Println("UDP client : ", addr)
	received := string(buffer[:n])
	fmt.Println("Received from UDP client :  ", string(buffer[:n]))

	if err != nil {
		log.Fatal(err)
	}

	if received == "\\uptime" {
		conn.WriteToUDP([]byte(time.Since(runningTime).String()), addr)
	} else if received == "\\reqnum" {
		conn.WriteToUDP([]byte(strconv.Itoa(reqNum)), addr)
    } else {
		message := []byte("Hello UDP client!")
		_, err = conn.WriteToUDP(message, addr)

		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	hostName := "localhost"
	portNum := "6000"
	service := hostName + ":" + portNum

	udpAddr, err := net.ResolveUDPAddr("udp4", service)

	if err != nil {
		log.Fatal(err)
	}

	// setup listener for incoming UDP connection
	ln, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("UDP server up and listening on port 6000")

	defer ln.Close()

	for {
		// wait for UDP client to connect
		handleUDPConnection(ln)
	}

}