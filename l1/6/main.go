package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func ConditionWay() {
	fmt.Println("Способ завершения горутины с помощью условий")
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			n := rand.Intn(10)
			time.Sleep(100 * time.Millisecond)
			if n > 7 {
				fmt.Println("Работа горутины завершена")
				return
			}
			fmt.Println("n =", n)
		}
	}(wg)

	wg.Wait()
}

func NotifyWay() {
	fmt.Println("Способ завершения горутины с помощью канала уведомления")
	done := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup, c chan struct{}) {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			default:
				n := rand.Intn(10)
				time.Sleep(100 * time.Millisecond)
				fmt.Println("n =", n)
			}

		}
	}(wg, done)
	time.Sleep(3 * time.Second)
	done <- struct{}{}
	fmt.Println("Работа горутины завершена")
	wg.Wait()
}

func ContextWay() {
	fmt.Println("Способ завершения горутины с помощью контекста")
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup, ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Поступил сигнал завершения контекста")
				return
			default:
				n := rand.Intn(10)
				time.Sleep(100 * time.Millisecond)
				fmt.Println("n =", n)
			}
		}
	}(wg, ctx)

	time.Sleep(3 * time.Second)
	cancel()
	wg.Wait()
}
func GoExitWay() {
	fmt.Println("Способ завершения горутины с помощью GoExit")
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			n := rand.Intn(10)
			time.Sleep(100 * time.Millisecond)
			if n > 7 {
				fmt.Println("используется runtime.GoExit")
				runtime.Goexit()
			}
			fmt.Println("n =", n)
		}
	}(wg)
	wg.Wait()
}
func TimeAfterWay() {
	fmt.Println("Способ завершения горутины с помощью time.After")
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		timeout := time.After(5 * time.Second)
		for {
			select {
			case <-timeout:
				fmt.Println("Завершение по таймауту")
				return
			default:
				n := rand.Intn(10)
				time.Sleep(100 * time.Millisecond)
				fmt.Println("n =", n)

			}
		}
	}(wg)
	wg.Wait()
}
func main() {
	ConditionWay()
	NotifyWay()
	ContextWay()
	GoExitWay()
	TimeAfterWay()
}
