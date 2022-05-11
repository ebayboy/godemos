package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

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
