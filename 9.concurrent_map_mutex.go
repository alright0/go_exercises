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

type SafeCounter struct {
	mutex sync.Mutex
	m     map[string]int
}

func (s *SafeCounter) Inc(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.m[key]++
}

func (s *SafeCounter) Get(key string) int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.m[key]
}

func WriteToMap(key string, sm *SafeCounter, wg *sync.WaitGroup) {
	defer wg.Done()

	sm.Inc(key)
}

func ReadFromMap(key string, sm *SafeCounter, wg *sync.WaitGroup) {
	defer wg.Done()

	value := sm.Get(key)
	fmt.Println(value)
}

func main() {
	wg := sync.WaitGroup{}

	sm := SafeCounter{
		m: make(map[string]int),
	}

	for i := 0; i < 1000; i++ {
		wg.Add(2)
		go WriteToMap("first", &sm, &wg)
		go ReadFromMap("first", &sm, &wg)
	}

	wg.Wait()

}
