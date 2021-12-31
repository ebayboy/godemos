package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var expr string = "(\\d+)@qq.com"

//获取网站内容, 正则过滤获取@qq.com邮箱
func GetEmail() {
	fmt.Println("expr:", expr)

	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	defer resp.Body.Close()

	//读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error:%v", err.Error())
		return
	}
	page := string(pageBytes)
	fmt.Printf("page:%v\n", page)

	// MustCompile && MustCompilePosix区别：
	// 带POSIX后缀的不同点在于其使用POSIX语法，该语法使用最长最左方式搜索，而不带该后缀的方法是采用最左方式搜索
	//（如[a-z]{2,4}这样的正则表达式，应用于"aa09aaa88aaaa"这个文本串时，带POSIX后缀的将返回aaaa，不带后缀的则返回aa）。
	re := regexp.MustCompile(expr)

	// self.IPMatcher.FindStringSubmatch && re.FindAllStringSubmatch 区别
	results := re.FindAllStringSubmatch(page, -1)

	// 遍历结果
	for _, result := range results {
		fmt.Printf("email:%v\n", result[0]) //分组0: 整个匹配内容  qq号码@qq.com
		fmt.Printf("qq:%v\n", result[1])    //分组1: (\\d+)匹配的内容， qq号码
	}
}

func main() {
	GetEmail()
}
