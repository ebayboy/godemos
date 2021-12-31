package main

import (
	"log"
	"net"
)

func testUDP() {

	addr := "0.0.0.0"
	netIP := net.ParseIP(addr)
	srvAddr := net.UDPAddr{
		IP:   netIP,
		Port: 20001,
	}

	//dailUDP
	connector, err := net.DialUDP("udp", nil, &srvAddr)
	if err != nil {
		log.Fatalf("Error:net.DialUDP", err.Error())
	}
	defer connector.Close()
	log.Printf("DialUDP:%v:%v\n", addr, 20001)

	//writetoUDP
	//注意：此处不能使用WriteToUDP，因为Client写入的服务端只有一个， 不需要指定地址
	n, err := connector.Write(([]byte)("hello udp"))
	if err != nil {
		log.Fatalf("Error: connector.WriteToUDP %v\n", err.Error())
	}
	log.Printf("Write size:%v\n", n)

	//readfromUDP
	data := make([]byte, 4096)
	//注意： 读取的时候
	n, remoteAddr, err := connector.ReadFromUDP(data)
	if err != nil {
		log.Fatalf("Error: connector.ReadFromUDP!%v\n", err.Error())
	}
	log.Printf("n:%v data:%v remoteAddr:%v\n", n, string(data[:n]), remoteAddr)
}

func main() {
	testUDP()
}
