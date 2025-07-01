package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

type WorkerPool struct {
	tasks      chan int
	results    chan int
	wg         sync.WaitGroup
	numWorkers int
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		tasks:      make(chan int, numWorkers),
		results:    make(chan int, numWorkers),
		numWorkers: numWorkers,
	}
}

func (w *WorkerPool) Start() {
	w.wg.Add(w.numWorkers)
	for i := 0; i < w.numWorkers; i++ {
		go w.worker()
	}
}

func (w *WorkerPool) worker() {
	defer w.wg.Done()
	for task := range w.tasks {
		result := task * task
		n := time.Duration(rand.Intn(1000)) * time.Millisecond // Имитация работы
		time.Sleep(n)
		w.results <- result
	}
}

func main() {
	numWorkers := flag.Int("n", 1, "number of workers") // получаем значение параметра
	flag.Parse()
	if *numWorkers < 1 {
		*numWorkers = 1
	}
	pool := NewWorkerPool(*numWorkers)

	// Запускаем пул
	pool.Start()

	sigChan := make(chan os.Signal, 1) // обработка сигнала прерывания (ctrl + c)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		count := 1
		for {
			select {
			case <-sigChan:
				fmt.Println("Завершение работы")
				close(pool.tasks)
				return
			default:
				pool.tasks <- count
				count++
			}

		}
	}()

	// Закрываем канал после окончания работы
	go func() {
		pool.wg.Wait()
		close(pool.results)
	}()

	// Получаем результаты и печатаем
	for result := range pool.results {
		fmt.Printf("Received result: %d\n", result)
	}
}
