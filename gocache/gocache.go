// Package main provides ...
package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

type Student struct {
	Name string
	Age  int
}

func main() {

	c := cache.New(5*time.Minute, 10*time.Minute)
	gob.Register(&Student{})
	cacheFile := ".cachedata"

	if _, err := os.Stat(cacheFile); err == nil {
		if err := c.LoadFile(cacheFile); err != nil {
			fmt.Println("error: load cache!", err.Error())
			return
		}
		fmt.Println("load cache count:", c.ItemCount())
	}

	if c.ItemCount() == 0 {
		fmt.Println("set cache foo, baz!")
		stu1 := Student{Name: "fanpf", Age: 36}
		stu2 := Student{Name: "rose", Age: 35}
		c.Set("fanpf", &stu1, cache.DefaultExpiration)
		c.Set("rose", &stu2, cache.NoExpiration)
	} else {
		fmt.Println("not set cache!")
	}

	stu, found := c.Get("fanpf")
	if found {
		fmt.Println("get fanpf:", stu)
	}

	if err := c.SaveFile(cacheFile); err != nil {
		fmt.Println("error: save cache!", err.Error())
	}

	time.Sleep(time.Second * 3)
}
