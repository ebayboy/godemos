package main

import "fmt"

func main() {
	//定义一个 [key:string]value:string类型的map
	countryCapitalMap := make(map[string]string)

	//len获取的是map的容量
	fmt.Println("len:", len(countryCapitalMap))
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Roma"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "Delhi"

	fmt.Println("len:", len(countryCapitalMap))

	for key, value := range countryCapitalMap {
		fmt.Println(key, "Capital is", value)
	}

	//检查元素是否在map中
	capital, ok := countryCapitalMap["American"]
	fmt.Println("capital:", capital, " ok:", ok)
	if ok {
		fmt.Println("Check American capital:", capital)
	} else {
		fmt.Println("Check American not exist!")
	}

	//map delete
	delete(countryCapitalMap, "France")
	fmt.Println("============After delete France===========")
	for i, i2 := range countryCapitalMap {
		fmt.Println(i, "Capital is", i2)
	}
}
