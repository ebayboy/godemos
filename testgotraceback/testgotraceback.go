
package main

import (
    "syscall"
    "fmt"
    "time"
)

//生成core文件条件：

//1. 设置ulimit -c 
//2. 开启GOTRACEBACK： env GOTRACEBACK=crash ./testgotraceback

func initSetLimit(cpu_max uint64, core_max uint64) error {
    var rlimit syscall.Rlimit

    // 限制cpu个数
    rlimit.Cur = 1
    rlimit.Max = cpu_max
    syscall.Setrlimit(syscall.RLIMIT_CPU, &rlimit)
    err := syscall.Getrlimit(syscall.RLIMIT_CPU, &rlimit)
    if err != nil {
        return err
    }

    //set core limit
    rlimit.Cur = 100 //以字节为单位
    rlimit.Max = rlimit.Cur + core_max
    if err := syscall.Setrlimit(syscall.RLIMIT_CORE, &rlimit); err != nil {
        return err
    }
    if err := syscall.Getrlimit(syscall.RLIMIT_CORE, &rlimit); err != nil {
        return err
    }

    return nil
}

func saferoutine(c chan bool) {
    for i := 0; i < 10; i++ {
        fmt.Println("Count:", i)
        time.Sleep(1 * time.Second)
    }
    c <- true
}
func panicgoroutine(c chan bool) {
    time.Sleep(5 * time.Second)
    panic("Panic, omg ...")
    c <- true
}

func main() {
    /*
    if err :=  initSetLimit(2, 2*1024*1024); err != nil {
        panic(err)
    }
    */

    c := make(chan bool, 2)
    go saferoutine(c)
    go panicgoroutine(c)
    for i := 0; i < 2; i++ {
        <-c
    }
}

