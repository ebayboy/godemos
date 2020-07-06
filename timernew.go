package main

import (
	"git.jd.com/fanpengfei1/riskstats/pkg/mod/golang.org/x/tools@v0.0.0-20200513154647-78b527d18275/go/analysis/passes/printf"
	"fmt"
	"sync"
	"time"
)

type job struct {
	Msg       string
	StartTime int64
}

var jobs = make([]job, 1024)

func main() {

	for i := 0; i < 5; i++ {
		jobstr := printf("Msg:[%v]", i)
		curTime := time.Unix()
		jobs = append(jobs, {jobstr, curTime})
	}

	log.Println("Jobs:", jobs)

	tiker := time.NewTicker(time.Second)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(<-tiker.C)
		}
		wg.Done()
	}()

	wg.Wait()
}
