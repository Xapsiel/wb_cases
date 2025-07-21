package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

type WorkerPool struct {
	tasks      chan int
	wg         sync.WaitGroup
	numWorkers int
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		tasks:      make(chan int, numWorkers),
		numWorkers: numWorkers,
	}
}

func (w *WorkerPool) Start(ctx context.Context) {
	w.wg.Add(w.numWorkers)
	for i := 0; i < w.numWorkers; i++ {
		go w.worker(i, ctx)
	}
}

func (w *WorkerPool) worker(id int, ctx context.Context) {
	defer w.wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Завершение по сигналу контекста\n", id)
			return
		case task, ok := <-w.tasks:
			if !ok {
				fmt.Printf("Worker %d: Канал задач закрыт\n", id)
				return
			}
			result := task * task
			fmt.Printf("Worker %d: Received %d, Result %d\n", id, task, result)
		}
	}
}

func main() {
	numWorkers := flag.Int("n", 1, "number of workers")
	delay := flag.Duration("delay", 10*time.Millisecond, "delay between task generations")
	flag.Parse()

	if *numWorkers < 1 {
		fmt.Fprintln(os.Stderr, "Ошибка: количество воркеров должно быть больше 0")
		os.Exit(1)
	}

	// Создаем контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Создаем пул воркеров
	pool := NewWorkerPool(*numWorkers)
	pool.Start(ctx)

	// Настраиваем обработку SIGINT
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	// Запускаем генерацию данных
	go func() {
		count := 1
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Генератор данных: Завершение по сигналу контекста")
				return
			default:
				pool.tasks <- count
				count++
				time.Sleep(*delay)
			}
		}
	}()

	// Обрабатываем сигнал SIGINT
	<-sigChan
	fmt.Println("Получен сигнал прерывания, завершаем работу...")
	cancel()          // Отменяем контекст
	close(pool.tasks) // Закрываем канал задач для надежности

	// Ожидаем завершения всех воркеров
	pool.wg.Wait()
	fmt.Println("Все воркеры завершены.")
}
