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

func binarySearch(n int, arr []int) int {

	left, right := 0, len(arr)-1 // левая и правая граница поиска
	for left <= right {          //до тех пор пока левая граница меньше или равно правой
		mid := left + (right-left)/2 // находим середину слайса
		switch {
		case arr[mid] == n:
			return mid // возвращаем индекс элемента
		case arr[mid] < n:
			left = mid + 1 // смещаем левую границу на середина+1
		case arr[mid] > n:
			right = mid - 1 // смещаем правую границу на середина+1
		}
	}
	return -1 // возвращаем -1
}
func main() {
	arr := GenArray(100)
	res := quickSort(arr)

	for i := 0; i < 100; i++ {
		index := binarySearch(i, res)
		if index == -1 {
			continue
		}
		fmt.Println(fmt.Sprintf("%d: %d", i, index))

	}

}
func GenArray(n int) []int {
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = rand.Intn(1000)
	}
	return array
}
