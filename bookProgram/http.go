package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Please visit http://127.0.0.1:12345/")
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		s := fmt.Sprint("Hello, world! Time: %s", time.Now().String)
		fmt.Fprintf(w, "%v", s)
	})

	if err := http.ListenAndServe(":12345", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
