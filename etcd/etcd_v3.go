package main

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// etcd clientv3 put/get demo

func main() {

	//1. clientv3.New
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 3,
	})
	if err != nil {
		log.Fatalln("Error: clientv3.New", err.Error())
	}
	defer cli.Close()

	log.Println("connect to etcd success")

	// 3.1 Put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "lmh", "lmh")
	cancel()
	if err != nil {
		log.Fatalf("put to etcd failed, err:%v\n", err)
	}
	log.Printf("Put lmh:lmh\n")

	// 3.2 Get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "lmh")
	cancel()
	if err != nil {
		log.Fatalf("get from etcd failed, err:%v\n", err)
	}

	log.Printf("Get:")
	for _, kv := range resp.Kvs {
		log.Printf("%s:%s\n", kv.Key, kv.Value)
	}
}
