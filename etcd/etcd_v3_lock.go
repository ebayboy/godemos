package main

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

//基于etcd实现分布式锁

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	defer cli.Close()
	log.Println("connect ok!")

	//创建会话1
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatalln("err:", err.Error())
	}
	defer s1.Close()
	m1 := concurrency.NewMutex(s1, "/my-lock/")

	//创建会话2
	s2, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatalln("err:", err.Error())
	}
	defer s2.Close()
	m2 := concurrency.NewMutex(s2, "/my-lock/")

	//会话s1获取锁
	if err := m1.Lock(context.TODO()); err != nil {
		log.Fatalln(err)
	}
	log.Println("acquired lock for s1")

	//m2锁定后退出协程，关闭m2locked chan
	m2locked := make(chan struct{})
	go func() {
		defer close(m2locked)

		// 等待直到会话s1释放了/my-lock/的锁
		if err := m2.Lock(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("release lock for s1 after 3 Second...")
	time.Sleep(time.Second * 3)
	if err := m1.Unlock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	log.Println("release lock for s1")

	<-m2locked
	log.Println("acquire lock for s2")
}
