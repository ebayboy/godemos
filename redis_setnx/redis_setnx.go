package main

import (
	"sync"
	"time"

	"github.com/go-redis/redis"
)

//10个goroutine并发操作redis，
//每次取出redis的counter，进行+1操作， 之后set回去

func inc() {
	//get redis cli
	cli := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "123456", DB: 0})

	lockKey := "counter_lock"
	counterKey := "counter"

	//lock
	resp := cli.SetNX(lockKey, 1, time.Second*5)
	lockRes, err := resp.Result()
	if err != nil || !lockRes {
		println("lock result:", lockRes)
		return
	}

	//counter ++
	getResp := cli.Get(counterKey)
	cntVal, err := getResp.Int64()
	if err == nil || err == redis.Nil {
		println("err:", err)
		cntVal++
		resp := cli.Set(counterKey, cntVal, 0)
		res, err := resp.Result()
		if err != nil {
			println("set value error!", err.Error())
			return
		} else {
			println("set redis result:", res)
		}
	}

	println("current counter is ", cntVal)

	delResp := cli.Del(lockKey)
	unlockRes, err := delResp.Result()
	if err == nil && unlockRes > 0 {
		println("unlock success!")
	} else {
		println("unlock failed!", err.Error())
	}

}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inc()
		}()
	}

	wg.Wait()
}
