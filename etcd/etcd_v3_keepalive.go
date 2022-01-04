package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// etcd keepAlive

/*
Leases:列举所有etcd中的租约。
KeepAlive:自动定时的续约某个租约。
KeepAliveOnce:为某个租约续约一次。
Close:关闭当前客户端建立的所有租约。
*/

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect to etcd success.")
	defer cli.Close()

	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}

	_, err = cli.Put(context.TODO(), "/lmh/", "lmh", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}

	// the key '/lmh/' will be kept forever
	ch, kaerr := cli.KeepAlive(context.TODO(), resp.ID)
	if kaerr != nil {
		log.Fatal(kaerr)
	}
	for {
		ka := <-ch
		fmt.Println("ttl:", ka.TTL)
	}
}
