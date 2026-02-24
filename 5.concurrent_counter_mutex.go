package main

import (
	"fmt"
	"sync"
)

/*
Конкурентный счётчик (гонка данных)

Создай переменную counter.
Запусти 10 горутин.
Каждая увеличивает counter на 1.

В конце выведи результат
реализация через mutex
*/

func incrementMu(counter *int32, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	defer mu.Unlock()
	mu.Lock()
	*counter++
}
func main() {
	var counter int32
	maxGoroutines := 10
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < maxGoroutines; i++ {
		wg.Add(1)
		go incrementMu(&counter, &wg, &mu)
	}
	wg.Wait()
	fmt.Printf("final counter = %d\n", counter)
}
