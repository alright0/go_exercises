package main

import (
	"fmt"
	"sync"
)

/*
Сделать структуру:

type SafeCounter struct { ... }

Требования:
Inc(key string)
Get(key string) int
конкурентно безопасная
использует Mutex
*/

type SafeCounterRw struct {
	mutex sync.RWMutex
	m     map[string]int
}

func (s *SafeCounterRw) Inc(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.m[key]++
}

func (s *SafeCounterRw) Get(key string) int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.m[key]
}

func WriteToMapRw(key string, sm *SafeCounterRw, wg *sync.WaitGroup) {
	defer wg.Done()

	sm.Inc(key)
}

func ReadFromMapRw(key string, sm *SafeCounterRw, wg *sync.WaitGroup) {
	defer wg.Done()

	value := sm.Get(key)
	fmt.Println(value)
}

func main() {
	wg := sync.WaitGroup{}

	sm := SafeCounterRw{
		m: make(map[string]int),
	}

	for i := 0; i < 1000; i++ {
		wg.Add(2)
		go WriteToMapRw("first", &sm, &wg)
		go ReadFromMapRw("first", &sm, &wg)
	}

	wg.Wait()

}
