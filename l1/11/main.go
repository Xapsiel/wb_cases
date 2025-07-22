package main

import "fmt"

func main() {
	a := GenArray(100, 2)
	b := GenArray(100, 1)
	res := Intersection(a, b)
	fmt.Println("Пересечение:")
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
func GenArray(n int, k int) []int {
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = i * k
	}
	return array
}
func Intersection(f, s []int) []int {
	m := make(map[int]int)
	res := make([]int, 0)
	for _, v := range f {
		m[v]++
	}
	for _, v := range s {
		if m[v] > 0 {
			res = append(res, v)
		}
	}
	return res
}
