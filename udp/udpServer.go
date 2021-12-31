package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	//listen
	var laddr net.UDPAddr

	laddr.IP = net.IPv4(0, 0, 0, 0)
	laddr.Port = 20001

	listener, err := net.ListenUDP("udp", &laddr)
	if err != nil {
		log.Fatalf("Error:", err.Error())
	}
	fmt.Println("net.ListenUDP 0.0.0.0:20001")

	data := make([]byte, 4096)
	for {
		//readfrom
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("Error:", err.Error())
			continue
		}
		fmt.Printf("n:%d remoteAddr:%v", n, remoteAddr.String())

		//write to
		n1, err1 := listener.WriteToUDP(data[:n], remoteAddr)
		if err1 != nil {
			fmt.Println("Error:", err1.Error())
		}
		fmt.Printf("write: %d:%v ok\n", n1, string(data[:n]))
	}
}
