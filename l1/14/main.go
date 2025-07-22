package main

import (
	"fmt"
)

func main() {
	res := GenArray()
	for _, v := range res {
		switch v.(type) {
		case int:
			fmt.Println(fmt.Sprintf("Переменная типа int"))
		case string:
			fmt.Println(fmt.Sprintf("Переменная типа string"))
		case bool:
			fmt.Println(fmt.Sprintf("Переменная типа bool"))
		case chan int:
			fmt.Println(fmt.Sprintf("Переменная типа chan int"))
		}
	}
}

func GenArray() []interface{} {
	res := make([]interface{}, 4)
	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			res[i] = i
		case 1:
			res[i] = fmt.Sprintf("%d", i)
		case 2:
			res[i] = (i % 2) == 0
		case 3:
			res[i] = make(chan int)
		}
	}
	return res
}
