package main

import (
	"fmt"
	"log"

	"github.com/Knetic/govaluate"
)

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
	expr := "ip_hit_waf > 0 ? (ip_hit_waf + ip_4xx)/ip_total : 0"
	//expr := "(ip_waf_ratio >= 0.5 ? 4 : 0) + (ip_waf_ratio >= 0.3 && ip_waf_ratio < 0.5 ? 3:0) + (ip_waf_ratio >= 0.2 && ip_waf_ratio < 0.3 ? 2:0) + (ip_waf_ratio >= 0.1 && ip_waf_ratio < 0.2 ? 1:0)"

	expression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		log.Fatalln("err:", err.Error())
	}

	parameters := map[string]interface{}{
		"ip_hit_waf": 0,
		"ip_4xx":     5,
		"ip_total":   10,
	}
	result, err := expression.Evaluate(parameters)
	if err != nil {
		log.Fatalln("err:", err.Error())
	}

	log.Println("result:", result)

}

var functions map[string]govaluate.ExpressionFunction

func init() {
	functions = make(map[string]govaluate.ExpressionFunction, 0)
	functions["F_field_score"] = func(args ...interface{}) (interface{}, error) {
		fmt.Println("func: F_field_score")
		length := len(args[0].(string))
		return (float64)(length), nil

	}
	functions["F_field_days"] = func(args ...interface{}) (interface{}, error) {
		fmt.Println("func: F_field_days")
		length := len(args[0].(string))
		return (float64)(length), nil

	}
	functions["F_field_wins"] = func(args ...interface{}) (interface{}, error) {
		fmt.Println("func: F_field_wins")
		length := len(args[0].(string))
		return (float64)(length), nil
	}
}

func test_function() error {

	//eva F_field_score
	expString := "F_field_days('ip_hit_waf')"
	expression, err := govaluate.NewEvaluableExpressionWithFunctions(expString, functions)
	if err != nil {
		return err
	}

	result, err := expression.Evaluate(nil)
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
}
