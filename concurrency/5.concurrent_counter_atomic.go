package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Конкурентный счётчик (гонка данных)

Создай переменную counter
Запусти 10 горутин
Каждая увеличивает counter на 1

В конце выведи результат
реализация через atomic
*/

func increment(counter *int32, wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt32(counter, 1)
}

func main() {
	var counter int32
	maxGoroutines := 10
	var wg sync.WaitGroup

	for i := 0; i < maxGoroutines; i++ {
		wg.Add(1)
		go increment(&counter, &wg)
	}

	wg.Wait()
	fmt.Printf("final counter = %d\n", counter)
}
