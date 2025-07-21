package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func Generator(ctx context.Context) chan int {
	in := make(chan int)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(in)
				return
			default:
				n := rand.Intn(100)
				in <- n
			}
		}

	}()
	return in
}

func main() {
	delay := flag.Duration("delay", 10*time.Second, "delay between task generations")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), *delay)
	defer cancel()
	result := Generator(ctx)
	for n := range result {
		fmt.Println(n)
	}
	fmt.Println("Время отправки в канал истекло")

}
