package main

import (
	"log"

	"github.com/zieckey/etcdsync"
)

//TODO:
//etcdsync.Lock failed client: response is invalid json. The endpoint is probably not valid etcd cluster endpoint.
func main() {
	//etcdsync.SetDebug(true)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	m, err := etcdsync.New("/etcdsync", 10, []string{"http://127.0.0.1:2379"})
	if m == nil {
		println("etcdsync.NewMutex failed", err.Error())
		return
	}

	err = m.Lock()
	if err != nil {
		println("etcdsync.Lock failed", err.Error())
		return
	} else {
		println("etcdsync.Lock OK")
	}

	println("Get the lock. Do something here.")

	err = m.Unlock()
	if err != nil {
		println("etcdsync.Unlock failed")

	} else {
		println("etcdsync.Unlock OK")
	}
}
