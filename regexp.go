package main

import (
	"fmt"
	"regexp"
)

//** re.Match : 验证正则表达式是否匹配 []byte, 返回bool, 性能最高
//re.MatchString : 验证正则表达式是否匹配 string, 返回bool
//** re.FindAllSubmatch: 非贪婪模式，分组, 性能比FindAllStringSubmatch高 返回字符串数组分组
//re.FindAllStringSubmatch: 贪婪模式， 分组 , 返回字符串数组分组

func main() {
	//分组0: a(x*)b(y|z)c
	//分组1: (x*)
	//分组2: (y|z)

	re := regexp.MustCompile("a(x*)b(y|z)c")

	//Match []byte
	fmt.Printf("re.Match========\n")
	fmt.Printf("%v\n", re.Match([]byte("-axxxbyc-axxbyc-axbyc-")))
	fmt.Printf("%v\n", re.Match([]byte("-axxxbyc-axxbyc-axbyc-")))
	fmt.Printf("%v\n", re.Match([]byte("-axxxbyc-axxbyc-axbyc-")))
	fmt.Printf("%v\n", re.Match([]byte("-abzc-abzc-")))
	fmt.Printf("%v\n", re.Match([]byte("-abzc-abzc-")))
	fmt.Printf("%v\n", re.Match([]byte("-aczc-")))

	//MatchString
	fmt.Printf("re.MatchString========\n")
	fmt.Printf("%v\n", re.MatchString("-axxxbyc-axxbyc-axbyc-"))
	fmt.Printf("%v\n", re.MatchString(("-axxxbyc-axxbyc-axbyc-")))
	fmt.Printf("%v\n", re.MatchString(("-axxxbyc-axxbyc-axbyc-")))
	fmt.Printf("%v\n", re.MatchString(("-abzc-abzc-")))
	fmt.Printf("%v\n", re.MatchString(("-abzc-abzc-")))
	fmt.Printf("%v\n", re.MatchString(("-aczc-")))

	//["axxxbyc" "xxx" "y"]中:
	//分组0： axxxbyc
	//分组1: xxx
	//分组2: y
	fmt.Printf("re.FindAllSubmatch=======\n")
	fmt.Printf("%v\n", re.FindAllSubmatch([]byte("-axxxbyc-axxbyc-axbyc-"), -1)) //[["axxxbyc" "xxx" "y"] ["axxbyc" "xx" "y"] ["axbyc" "x" "y"]]
	fmt.Printf("%v\n", re.FindAllSubmatch([]byte("-axxxbyc-axxbyc-axbyc-"), 2))  //[["axxxbyc" "xxx" "y"] ["axxbyc" "xx" "y"]]
	fmt.Printf("%v\n", re.FindAllSubmatch([]byte("-axxxbyc-axxbyc-axbyc-"), 1))  //[["axxxbyc" "xxx" "y"]]
	fmt.Printf("%v\n", re.FindAllSubmatch([]byte("-abzc-abzc-"), -1))            //[["abzc" "" "z"] ["abzc" "" "z"]]
	fmt.Printf("%v\n", re.FindAllSubmatch([]byte("-abzc-abzc-"), 1))             //[["abzc" "" "z"]]
	fmt.Printf("%v\n", re.FindAllSubmatch([]byte("-aczc-"), -1))                 //[],整个都不匹配，更没有分组匹配，将返回空数组

	//贪婪匹配，分组
	fmt.Printf("=re.FindAllStringSubmatch======\n")
	fmt.Printf("%v\n", re.FindAllStringSubmatch("-axxxbyc-axxbyc-axbyc-", -1)) //[["axxxbyc" "xxx" "y"] ["axxbyc" "xx" "y"] ["axbyc" "x" "y"]]
	fmt.Printf("%v\n", re.FindAllStringSubmatch("-axxxbyc-axxbyc-axbyc-", 2))  //[["axxxbyc" "xxx" "y"] ["axxbyc" "xx" "y"]]
	fmt.Printf("%v\n", re.FindAllStringSubmatch("-axxxbyc-axxbyc-axbyc-", 1))  //[["axxxbyc" "xxx" "y"]]
	fmt.Printf("%v\n", re.FindAllStringSubmatch("-abzc-abzc-", -1))            //[["abzc" "" "z"] ["abzc" "" "z"]]
	fmt.Printf("%v\n", re.FindAllStringSubmatch("-abzc-abzc-", 1))             //[["abzc" "" "z"]]
	fmt.Printf("%v\n", re.FindAllStringSubmatch("-aczc-", -1))                 //[],整个都不匹配，更没有分组匹配，将返回空数组
}
