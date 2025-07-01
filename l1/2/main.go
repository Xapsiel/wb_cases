package main

import (
	"fmt"
	"sync"
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
		go w.worker(i)
	}
}

func (w *WorkerPool) worker(id int) {
	defer w.wg.Done()
	for task := range w.tasks {
		result := task * task
		w.results <- result
	}
}

func main() {
	ar := []int{2, 4, 6, 8, 10}
	workerCount := len(ar)
	pool := NewWorkerPool(workerCount)

	// Запускаем пул
	pool.Start()

	// Отправляем в канал данные задач
	for _, num := range ar {
		pool.tasks <- num
	}
	close(pool.tasks)

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
