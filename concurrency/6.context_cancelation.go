package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

/*
Запусти 3 goroutine, которые работают в бесконечном цикле:

Через 3 секунды отменяй context.
*/

func infinityTask(worker int, ch chan string, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d canceled 1\n", worker)
			return
		default:
			randomValue := rand.IntN(300) + 1
			select {
			case <-ctx.Done():
				fmt.Printf("worker %d canceled 2\n", worker)
				return
			case ch <- fmt.Sprintf("worker %d send info: %d", worker, randomValue):
			}
		}
		time.Sleep(time.Millisecond * time.Duration(300))

	}
}

func main() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cancelAfter := time.Second * 3
	ch := make(chan string)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go infinityTask(i, ch, ctx, &wg)
	}

	go func() {
		time.Sleep(cancelAfter)
		cancel()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}
}
