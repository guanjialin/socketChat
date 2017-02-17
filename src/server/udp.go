package main

import (
	"log"
	"net"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp4", ":5000")
	if err != nil {
		log.Fatalln("udpaddr", err)
	}
	udpConn, err := net.ListenUDP("udp4", udpAddr)
	if err != nil {
		log.Fatalln("server udplisten", err)
	}
	defer udpConn.Close()

	for {
		var msg = make([]byte, 1024)
		_, addr, _ := udpConn.ReadFromUDP(msg)
		log.Println(addr, string(msg))
	}
}

func UDPListener() {

}
