package main

import (
	"fmt"
	"sync"
)

type SyncMap[K comparable, V any] struct {
	mu sync.Mutex
	m  map[K]V
}

func (s *SyncMap[K, V]) Set(k K, v V) {
	s.mu.Lock()
	s.m[k] = v
	s.mu.Unlock()
}
func (s *SyncMap[K, V]) Get(k K) (V, bool) {
	s.mu.Lock()
	v, ok := s.m[k]
	s.mu.Unlock()

	return v, ok
}
func (s *SyncMap[K, V]) Delete(k K) bool {
	s.mu.Lock()
	delete(s.m, k)
	s.mu.Unlock()
	return true
}

func New[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		m: make(map[K]V),
	}
}

func main() {
	syncMap := New[int, int]()
	wg := &sync.WaitGroup{}
	wg.Add(22)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				syncMap.Set(i, i)
				fmt.Println(fmt.Sprintf("Вставка %v: %v", i, i))
			}(wg)
		}
	}(wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				n, ok := syncMap.Get(i)
				if !ok {
					fmt.Println(fmt.Sprintf("Не найден %v", i))
					return
				}
				fmt.Println(fmt.Sprintf("Найден %v: %v", i, n))
			}(wg)
		}
	}(wg)

	wg.Wait()
	for i := 0; i < 10; i++ {
		syncMap.Delete(i)
	}

}
