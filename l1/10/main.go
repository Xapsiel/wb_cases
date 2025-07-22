package main

import (
	"fmt"
	"math"
)

func Group(temps []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, temp := range temps {
		key := int(math.Floor(temp/10)) * 10
		if key < 0 {
			key += 10
		}
		groups[key] = append(groups[key], temp)
	}

	return groups
}

func main() {
	temps := []float64{-254.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := Group(temps)

	for key := range groups {
		if temps, exists := groups[key]; exists {
			fmt.Printf("%d:%v\n", key, temps)
		}
	}
}
