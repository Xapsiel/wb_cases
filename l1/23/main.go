package main

import "fmt"

func Delete(ar []int, index int) ([]int, error) {
	if index < 0 || index >= len(ar) {
		return nil, fmt.Errorf("некорректный индекс")
	}
	copy(ar[index:], ar[index+1:])
	result := ar[:len(ar)-1]
	if cap(result) > len(result) {
		ar[len(ar)-1] = 0
	}
	return result, nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(fmt.Sprintf("Слайсы: %v", slice))
	index := 2
	slice, err := Delete(slice, index)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Полученный слайс: %v", slice))
}
