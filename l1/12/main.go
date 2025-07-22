package main

import "fmt"

func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}
	res := MakeSet(a)
	fmt.Println("Множество:")
	for _, v := range res {
		fmt.Println(v)
	}

}

func MakeSet(a []string) []string {
	r := make(map[string]bool, 0)
	res := make([]string, 0)
	for _, v := range a {
		if _, ok := r[v]; !ok {
			r[v] = true
			res = append(res, v)
		}
	}
	return res

}
