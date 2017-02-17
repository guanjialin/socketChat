package main

import (
	"log"
	"net"
	"time"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:5000")
	if err != nil {
		log.Fatalln("udpaddr", err)
	}
	udpConn, err := net.DialUDP("udp4", nil, udpAddr)
	if err != nil {
		log.Fatalln("client udpConn", err)
	}
	defer udpConn.Close()

	for {
		udpConn.Write([]byte("hello world"))
		time.Sleep(1 * time.Second)
	}
}
