package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
本质上是生产者消费者模型
可以有效控制goroutine数量，防止暴涨

需求：
计算一个数字的各个位数之和，例如数字123，结果为1+2+3=6
随机生成数字进行计算

解决方案：
+ 创建协程池(N)
+ 协程获取rand指定位数上的数字bNum, sum += bNum
*/

// 算法：
// 234 % 100 = 34
// 234 - 34 = 200
// 200/100 = 2

// 34 % 10 = 4
// 34 - 4 = 30
// 30/10 = 3

// 4

func GetIntByteNum(r int) (ret []int, rlen int) {
	// 获取数字r的长度
	rStr := fmt.Sprintf("%d", r)
	rlen = len(rStr)
	posNum := r
	for posLen := rlen; posLen > 0; posLen-- {
		divNum := 1
		modNum := 0
		subNum := 0
		byteNum := 0
		for i := 1; i < posLen; i++ {
			divNum *= 10
		}
		modNum = posNum % divNum
		subNum = posNum - modNum
		byteNum = subNum / divNum
		ret = append(ret, byteNum)
		posNum = modNum
	}
	return ret, rlen
}

func main() {
	sum := 0
	r := rand.Int()
	nums, rlen := GetIntByteNum(r)
	fmt.Printf("rand:%d sum:%d numsb:%d rlen:%d", r, sum, nums, rlen)

	//创建rlen个协程，对nums进行累加

	wg := sync.WaitGroup{}

	//初始化通道
	outCh := make(chan int, 10)

	inCh := make([]chan int, rlen)
	for i := 0; i < rlen; i++ {
		inCh[i] = make(chan int, 10)
	}

	//初始化协程, 对inCh输入的数据+1后写入到outCh
	for i := 0; i < rlen; i++ {
		wg.Add(1)
		go func(c chan int, o chan int) {
			defer wg.Done()
			for {
				num, ok := <-c
				if !ok {
					break
				}
				num += 1
				o <- num
			}
		}(inCh[i], outCh)
	}

	//启动outChan通道读取协程， 读取数据并累加
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			data, ok := <-outCh
			if !ok {
				break
			}
			sum += data
		}
	}()

	//使用协程对nums进行计算
	for k, v := range nums {
		inCh[k] <- v //写入数据
	}

	// 收到信号后关闭通道 -> 子协程退出 ->  wg.Wait完成 -> 主进程退出
	sigCh := make(chan os.Signal)
	defer close(sigCh)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			s := <-sigCh
			switch s {
			case syscall.SIGINT, syscall.SIGTERM:
				close(outCh)
				for _, ch := range inCh {
					close(ch)
				}
				return
			}
		}
	}()

	wg.Wait()

	fmt.Println("sum:", sum)
}
