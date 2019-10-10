package main

import (
	"fmt"
	"log"
	"net"
)

var rec int

func main() {
	rec = 1
	server()

}

func server() error {
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
func handleUDPConnection(conn *net.UDPConn) {

	// here is where you want to do stuff like read or write to client

	buffer := make([]byte, 1024)

	//n, addr, err := conn.ReadFromUDP(buffer)
	_, addr, err := conn.ReadFromUDP(buffer)

	//fmt.Println("UDP client : ", addr)
	//fmt.Println("Received from UDP client :  ", string(buffer[:n]))

	fmt.Printf("\033[2K\r Received: %d", rec)
	if err != nil {
		log.Fatal(err)
	}
	rec++

	// NOTE : Need to specify client address in WriteToUDP() function
	//        otherwise, you will get this error message
	//        write udp : write: destination address required if you use Write() function instead of WriteToUDP()

	// write message back to client
	message := []byte("Hello UDP client!")
	_, err = conn.WriteToUDP(message, addr)

	if err != nil {
		log.Println(err)
	}

}
