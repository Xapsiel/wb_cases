package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 { // меньше двух элементов
		return arr
	}
	pivot := arr[0]                               // опорный элемент
	left, right := make([]int, 0), make([]int, 0) // левый и правый подмассивы
	for i := range arr {
		if pivot > arr[i] {
			left = append(left, arr[i]) // меньше опорного - в левый
		} else if pivot < arr[i] {
			right = append(right, arr[i]) //больше  - в правый
		}
	}
	left = append(quickSort(left), pivot) // сортируем рекурсивно левый подмассив
	right = quickSort(right)              // сортируем правый
	return append(left, right...)         // соединяем
}

func main() {
	arr := GenArray(100)
	res := quickSort(arr)
	fmt.Println(res)

}
func GenArray(n int) []int {
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = rand.Intn(1000)
	}
	return array
}
