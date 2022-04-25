package main

import (
	"io/ioutil"
	"log"

	"github.com/tidwall/sjson"
)

/*
{
    "name": {"first": "Tom", "last": "Anderson"},
    "age":37,
    "children": ["Sara","Alex","Jack"],
    "fav.movie": "Deer Hunter",
    "friends": [
        {"first": "James", "last": "Murphy"},
        {"first": "Roger", "last": "Craig"}
    ]
}
*/

func main() {
	content, err := ioutil.ReadFile("./in.json")
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	value, _ := sjson.Set(string(content), "name.last", "fanpengfei----")
	print(value)
	value, _ = sjson.Set(value, "age", 11)
	print(value)

	ioutil.WriteFile("out.json", []byte(value), 0666)
}
