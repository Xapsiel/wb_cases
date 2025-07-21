package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func Generator(ctx context.Context) chan int {
	in := make(chan int) // канал для генерации
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(in) // закрытие канала
				return
			default:
				n := rand.Intn(100) // имитация полезной нагрузки
				in <- n
			}
		}

	}()
	return in
}

func main() {
	delay := flag.Duration("delay", 10*time.Second, "delay for stopping program") //продолжительность работы программы
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), *delay) //контекст с таймаутом
	defer cancel()
	result := Generator(ctx)
	for n := range result {
		fmt.Println(n) //чтение из канала. При закрытии - цикл завершится
	}
	fmt.Println("Время отправки в канал истекло")

}
