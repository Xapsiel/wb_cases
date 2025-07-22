package main

import (
	"fmt"
	"sync"
)

func GenArray(n int) []int {
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = i
	}
	return array
}
func Generator(out chan<- int) {
	a := GenArray(100)
	for i := range a {
		out <- a[i]
	}
	close(out)
}
func Processor(out chan<- int, in <-chan int) {
	for a, ok := <-in; ok; a, ok = <-in {
		out <- a * a
	}
	close(out)
}
func main() {
	in := make(chan int)
	out := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		Generator(in)
	}()
	go func() {
		defer wg.Done()
		Processor(out, in)
	}()
	go func() {
		defer wg.Done()
		for val := range out {
			fmt.Println(val)
		}
	}()
	wg.Wait()
}
