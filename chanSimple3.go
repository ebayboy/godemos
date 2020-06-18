package main

import (
	"log"
	"time"
)

func sum() {
	log.Println("start sum....")
	time.Sleep(1 * time.Hour)
	log.Println("start over....")
}

func Test() {
	go sum()

	time.Sleep(3 * time.Second)
	log.Println("Test exit!")
}

func main() {
	log.Println("main start!")

	Test()

	//time.Sleep(1 * time.Hour)

	log.Println("main over!")
}
