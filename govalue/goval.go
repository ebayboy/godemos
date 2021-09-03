package main

import (
	"log"

	"github.com/Knetic/govaluate"
)

func main() {

	expression, err := govaluate.NewEvaluableExpression("(mem_used / total_mem) * 100")
	if err != nil {
		log.Fatalln("err:", err.Error())
	}

	parameters := make(map[string]interface{}, 8)
	parameters["total_mem"] = 1024
	parameters["mem_used"] = 512

	result, err := expression.Evaluate(parameters)

	log.Println("result:", result)
}
