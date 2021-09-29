package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	var age, static_rate, max float64

	flag.Float64Var(&age, "a", 39, "age:")
	flag.Float64Var(&max, "m", 0, "max:")
	flag.Float64Var(&static_rate, "s", 70, "static_rate:")
	flag.Parse()

	fmt.Println("age:", age, " static_rate:", static_rate)

	sort_keys := []string{"E", "M", "T", "I"}
	inters := map[string][]float64{
		"E": []float64{0.59, 0.74},
		"M": []float64{0.75, 0.84},
		"T": []float64{0.83, 0.88},
		"I": []float64{0.975, 1.00},
	}

	if max == 0 {
		max = 220 - age
	}
	for _, k := range sort_keys {
		if v, ok := inters[k]; !ok {
			fmt.Println("Error: key not exist!", k)
			return
		} else {
			start := (max-static_rate)*v[0] + static_rate
			end := (max-static_rate)*v[1] + static_rate
			fmt.Printf("%s:[%d,%d]\n", k, int(math.Floor(start)), int(math.Ceil(end)))
		}
	}
}
