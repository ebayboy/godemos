package main

import (
	"io/ioutil"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	addr := "www.baidu.com:80"

	//Resolve
	log.Info("ResolveTCPAddr...")
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Error("ResolveTCPAddr:", err)
		return
	}

	//DialTcp
	log.Info("DialTCP...")
	myconn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer myconn.Close()
	if err != nil {
		log.Error("DialTCP:", err)
		return
	}

	//Write
	log.Info("Write...")
	if _, err := myconn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n")); err != nil {
		log.Error("Write:", err)
		return
	}

	//Read
	log.Info("ReadAll...")
	result, err := ioutil.ReadAll(myconn)
	if err != nil {
		log.Error("ReadAll:", err)
		return
	}
	log.Info("ReadAll result:", string(result))

	os.Exit(0)
}
