package main

import "fmt"

func reverse(runes []rune, left, right int) {
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
}

func reverseSeq(s string) string {
	runes := []rune(s)
	n := len(runes)
	reverse(runes, 0, n-1) // переворачиваем строку всю
	index := 0
	for i := 0; i < n; i++ {
		if runes[i] == ' ' {
			reverse(runes, index, i-1) //переворачиваем отдельное слово
			index = i + 1
		}
	}
	reverse(runes, index, n-1) // переворачиваем последнее слова
	return string(runes)

}

func main() {
	a := "snow dog sun"
	fmt.Println(reverseSeq(a))
}
