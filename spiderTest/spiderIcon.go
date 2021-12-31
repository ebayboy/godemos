package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sync"
	"sync/atomic"
)

/*
// 并发爬思路：
+ 解析网页内容， 用正则找出所有图片
+ 初始化图片链接数个协程
+ 将每个图片链接写入到一个协程通道
+ 爬取协程下载图片保存到本地
*/

func GetIconURLs() (iconURLs []string, err error) {
	resp, err := http.Get(`https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/1.html`)
	if err != nil {
		return nil, err
	}

	byteBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//write to file
	err = os.WriteFile("body.html", byteBody, 0666)
	if err != nil {
		fmt.Printf("Error:%v\n", err.Error())
	}

	bodyStr := string(byteBody)

	//获取body内容
	//https://uploadfile.bizhizu.cn/up/ec/b4/55/ecb45559bda357aea824fde70cdde995.jpg
	//正则表达式扒取图片连接
	reImg := `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
	re := regexp.MustCompile(reImg)

	res := re.FindAllStringSubmatch(bodyStr, -1)
	for _, v := range res {
		iconURLs = append(iconURLs, v[0])
	}

	return iconURLs, nil
}

func main() {

	var workDone int64 = 0
	iconURLs, err := GetIconURLs()
	if err != nil {
		fmt.Printf("Error:%v\n", err.Error())
		return
	}
	fmt.Printf("len:[%d] iconURLs:%v\n", len(iconURLs), iconURLs)

	chs := make([]chan string, len(iconURLs))
	wg := sync.WaitGroup{}

	//1.1 init chans
	for i := 0; i < len(chs); i++ {
		chs[i] = make(chan string, 64)
		defer close(chs[i])
	}

	//1.2 初始化 爬虫写出 goroutines
	for k, ch := range chs {
		wg.Add(1)
		go func(ch chan string, idx int) {
			defer wg.Done()

			url := <-ch //读取管道连接
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Error: goroutine[%d] error:%v\n", idx, err.Error())
			} else {
				byteBody, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error: goroutine[%d] error:%v\n", idx, err.Error())
				} else {
					//下载图片, 保存成文件
					filename := fmt.Sprintf("%d.jpg", idx)
					err := os.WriteFile(filename, byteBody, 0666)
					if err != nil {
						fmt.Printf("Error: goroutine[%d] error:%v\n", idx, err.Error())
					} else {
						atomic.AddInt64(&workDone, 1)
						fmt.Printf("[%d] Save file to %v\n", idx, filename)
					}
				}
			}
		}(ch, k)
	}

	//向管道里写连接, 每个管道写入一个链接
	wg.Add(1)
	go func() {
		defer wg.Done()
		for k, url := range iconURLs {
			chs[k] <- url
		}
	}()

	wg.Wait()
}
