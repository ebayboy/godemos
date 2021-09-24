package main

import (
	"flag"
	"fmt"
)

func main() {
	var age, static_rate float64

	flag.Float64Var(&age, "a", 39, "age:")
	flag.Float64Var(&static_rate, "s", 70, "static_rate:")
	flag.Parse()

	fmt.Println("age:", age, " static_rate:", static_rate)

	sort_keys := []string{"E", "M", "T", "I", "R"}

	inters := map[string][]float64{
		"E": []float64{0.59, 0.74},
		"M": []float64{0.75, 0.84},
		"T": []float64{0.83, 0.88},
		"I": []float64{0.975, 1.00},
		"R": []float64{1.05, 1.20},
	}

	max := 220 - age
	for _, k := range sort_keys {
		if v, ok := inters[k]; !ok {
			fmt.Println("Error: key not exist!", k)
			return
		} else {
			fmt.Printf("%s:%.2f ~ %.2f\n", k, (max-static_rate)*v[0]+static_rate, (max-static_rate)*v[1]+static_rate)
		}
	}
}
