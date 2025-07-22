package main

import "fmt"

func main() {
	Test()
}

func Test() {
	for i := -20; i < 20; i++ {
		for j := -20; j <= 20; j++ {
			a, b := Change(i, j)
			fmt.Println(fmt.Sprintf("(%d;%d) = (%d;%d)\t%t", i, j, a, b, i == b && j == a))
			/*
				(6;11) = (11;6)	true
				(6;12) = (12;6)	true
				(6;13) = (13;6)	true
				....
				(18;-1) = (-1;18)	true
				(18;0) = (0;18)	true
			*/
		}
	}
}
func Change(a, b int) (int, int) {
	b = a - b
	a = a - b
	b = a + b
	return a, b
}
