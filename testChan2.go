package main

import (
	"log"
	"sync"
)

func TestSum(wg *sync.WaitGroup, channel chan map[string]int) {
	TestInputWrite(channel)
	wg.Done()
}

func TestInputWrite(channel chan map[string]int) {

	log.Println("TestInputWrite...")
	keys := []string{"grpid",
		"host",
		"ts",
		"ip",
		"ip_reqs",
		"ip_static_reqs",
		"ip_dyn_reqs",
		"ip_same_ua_page",
		"ip_page_click_time_avg",
		"ip_reqs_referer",
		"ip_dyn_reqs_pages",
		"dev_with_ips",
		"ip_dyn_post_reqs"}

	m := make(map[string]int)

	for i, key := range keys {
		m[key] = i
	}
	log.Println("m:", m)
	channel <- m
}

func TestInputRead(channel chan map[string]int) {
	m := <-channel
	log.Println("TestInputRead:", m)
}

func main() {
	var wg sync.WaitGroup

	var LChan []chan map[string]int
	LChan = make([]chan map[string]int, 3)

	wg.Add(1)
	go TestSum(&wg, LChan[0])
	TestInputRead(LChan[0])

	wg.Wait()
}
