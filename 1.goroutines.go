package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Параллельная печать.

Программа запускает 3 goroutine
каждая goroutine печатает своё имя (worker 1, worker 2, worker 3) по 5 раз
*/

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 5; i++ {
		fmt.Printf("worker %d iteration %d\n", id, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	go func() {
		wg.Wait()
	}()
}
