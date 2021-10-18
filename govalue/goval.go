package main

import (
	"fmt"
	"log"

	"github.com/Knetic/govaluate"
)

func test_equ() {
	expr := "ip_waf_hit_score < 4 ? (ip_hit_waf_days_score > 0 || ip_hit_waf_wins_score > 0) : true"
	expression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		log.Fatalln("err:", err.Error())
	}

	parameters := map[string]interface{}{
		"ip_waf_hit_score":      4,
		"ip_hit_waf_days_score": 0,
		"ip_hit_waf_wins_score": 0,
	}
	result, err := expression.Evaluate(parameters)
	if err != nil {
		log.Fatalln("err:", err.Error())
	}

	log.Println("test equ result:", result)

}

func test_in_array() {
	expr := "ip_hit_waf in (1, 2, 3, 'tet')"
	expression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		log.Fatalln("err:", err.Error())
	}

	parameters := map[string]interface{}{
		"ip_hit_waf": "tet",
	}
	result, err := expression.Evaluate(parameters)
	if err != nil {
		log.Fatalln("err:", err.Error())
	}

	log.Println("in test result:", result)

}

func test_basic() {
	expr := "ip_hit_waf < 0.01 && ip_hit_waf > 0"
	expression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		log.Fatalln("err:", err.Error())
	}

	parameters := map[string]interface{}{
		"ip_hit_waf": 0.001,
	}
	result, err := expression.Evaluate(parameters)
	if err != nil {
		log.Fatalln("err:", err.Error())
	}

	log.Println("basic result:", result)
}

var functions map[string]govaluate.ExpressionFunction

func init() {
	functions = make(map[string]govaluate.ExpressionFunction, 0)
	functions["F_field_days"] = func(args ...interface{}) (interface{}, error) {
		var lvls []string

		switch args[1].(type) {
		case []string:
			lvls = append(lvls, args[1].([]string)...)
		default:
		}
		fmt.Println("func: F_field_days")
		fmt.Println("args:", args)
		fmt.Println("args[0]:", args[0])
		fmt.Println("args[1]:", args[1])
		fmt.Println("args[1]:", args[1])
		fmt.Println("args[2]:", args[2])
		fmt.Println("args[3]:", args[3])
		fmt.Println("args[4]:", args[4])
		length := len(args[0].(string))

		return (float64)(length), nil
	}
}

func test_function() error {

	//eva F_field_score
	expString := "F_field_days('ip_hit_waf',('L3', 'L2', 'L1'), (7,4,2), hoststat,fieldstat)"

	expression, err := govaluate.NewEvaluableExpressionWithFunctions(expString, functions)
	if err != nil {
		return err
	}
	parameters := map[string]interface{}{
		"hoststat":  4,
		"fieldstat": 0,
	}
	result, err := expression.Evaluate(parameters)
	if err != nil {
		return err
	}
	fmt.Println("function result:", result)

	return nil
}

func main() {
	test_basic()
	test_function()
	test_in_array()
	test_equ()
}
