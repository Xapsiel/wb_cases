package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex // мьютекс для синхронизации
	value int
}

func NewCounter() *Counter {
	return &Counter{
		value: 0,
	}
}
func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}
func (c *Counter) Decrement() {
	c.mu.Lock()
	c.value--
	c.mu.Unlock()
}
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
func main() {
	wg := &sync.WaitGroup{}
	c := NewCounter()
	wg.Add(10)
	for i := 0; i < 100; i++ {
		go func() {
			c.Increment()
			fmt.Println(c.Value())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("итоговое значение")
	fmt.Println(c.Value())
	// Запускал с -race. Гонок не обнаружено
}
