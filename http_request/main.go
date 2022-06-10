package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func sendReques() {

	//Method
	method := http.MethodPost

	//URL
	request_path := "http://www.baidu.com/s"

	//Body
	data := url.Values{}
	data.Set("body-key1", "body-value1")
	data.Set("body-key2", "body-value2")
	req, err := http.NewRequest(method, request_path, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}

	//Header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")

	//Args
	q := req.URL.Query()
	q.Add("arg1", "value1")
	req.URL.RawQuery = q.Encode()

	//Method

	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("resp body:", string(body), " code:", resp.StatusCode)
}

func readLogs() error {

	of, err := os.Open("./waf.log")
	if err != nil {
		return err
	}

	bio_reader := bufio.NewReader(of)
	for {
		line, err := bio_reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		log.Println("line:", line)
	}
}

func loadConf() ([]string, error) {
	confName := "access_log.conf"
	conf, err := ioutil.ReadFile(confName)
	if err != nil {
		return nil, err
	}

	fields := strings.Split(string(conf), "\n")
	if fields[len(fields)-1] == "" {
		fields = fields[:len(fields)-1]
	}
	log.Printf("lend:%d fields:%v last:[%s]", len(fields), fields, fields[len(fields)-1])

	return fields, nil
}

func main() {
	fields, err := loadConf()
	if err != nil {
		log.Panic("err:", err.Error())
	}

	for k, v := range fields {
		fmt.Println(k, ":", v)
	}

	if err := readLogs(); err != nil {
		log.Panic("err:", err.Error())
	}

	//sendReques()
}
