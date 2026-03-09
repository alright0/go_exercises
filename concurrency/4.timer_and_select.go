package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Таймер и select

Создай goroutine, которая делает "работу" 2 секунды.
Жди результат через канал, либо через 1 секунду завершай ожидание (timeout)
*/

func doWork(ch chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	goroutineWorkTime := time.Second * 2

	select {
	case <-ctx.Done():
		fmt.Println("goroutine cancelled in select 1")
		return
	case <-time.After(goroutineWorkTime):
		fmt.Println("goroutine works...")
	}

	select {
	case <-ctx.Done():
		fmt.Println("goroutine cancelled in select 2")
		return
	case ch <- 1:
		fmt.Println("goroutine done")
	}

}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	timeout := time.Second * 1

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	wg.Add(1)
	go doWork(ch, &wg, ctx)

	select {
	case res := <-ch:
		fmt.Printf("Task Done %d", res)
	case <-ctx.Done():
		fmt.Println("Timed out:", ctx.Err())
	}

	wg.Wait()
	fmt.Println("All goroutines finished")

}
