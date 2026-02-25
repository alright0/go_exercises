package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

/*
Нужно сделать пайплайн из 3 этапов:

Генератор чисел
Фильтр (оставляет только четные)
Воркер, который умножает число на 10

Каждый этап работает в своей горутине.
Данные передаются через каналы.
*/

func generator(ch1 chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	for {
		value := rand.IntN(100)
		select {
		case <-ctx.Done():
			close(ch1)
			return
		case ch1 <- value:
			time.Sleep(time.Millisecond * 50)
		}
	}
}

func filter(ch1 chan int, ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch2)

	for val := range ch1 {
		if val%2 == 0 {
			ch2 <- val
		}
	}
}

func multiplier(ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range ch2 {
		result := val * 10
		fmt.Println(result)
	}
}

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan int)
	ch2 := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(3)
	go generator(ch1, &wg, ctx)
	go filter(ch1, ch2, &wg)
	go multiplier(ch2, &wg)

	select {
	case <-time.After(time.Second * 5):
		cancel()
	}

	wg.Wait()

}
