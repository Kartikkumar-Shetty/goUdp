package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const message = ``

func main() {
	client()
}
func client() {
	hostName := "localhost"
	portNum := "6000"

	service := hostName + ":" + portNum

	RemoteAddr, err := net.ResolveUDPAddr("udp", service)

	//LocalAddr := nil
	// see https://golang.org/pkg/net/#DialUDP

	conn, err := net.DialUDP("udp", nil, RemoteAddr)

	// note : you can use net.ResolveUDPAddr for LocalAddr as well
	//        for this tutorial simplicity sake, we will just use nil

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Established connection to %s \n", service)
	log.Printf("Remote UDP address : %s \n", conn.RemoteAddr().String())
	log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())

	defer conn.Close()

	// write a message to server
	//message := []byte("Hello UDP server!")
	sent := 1
	for {
		time.Sleep(0.5 * time.Second)
		fmt.Printf("\033[2K\r Sent: %d", sent)
		_, err = conn.Write([]byte(message))
		sent++
		if sent == 10000 {
			break
		}
		if err != nil {
			log.Println(err)
		}
	}

	// receive message from server
	buffer := make([]byte, 1024)
	//n, addr, err := conn.ReadFromUDP(buffer)
	_, _, _ = conn.ReadFromUDP(buffer)

	//fmt.Println("UDP Server : ", addr)
	//fmt.Println("Received from UDP server : ", string(buffer[:n]))
}
