package main

import (
	"fmt"
	"sync"
)

/*
Producer → Consumer

1 producer (генерирует числа от 1 до 10)
2 consumer (выводят числа в консоль)

Producer отправляет числа в канал.
Consumers получают данные из канала, пока он не закрыт.
*/

func sendNumbers(ch chan int) {
	lastNum := 10
	for i := 1; i < lastNum+1; i++ {
		ch <- i
	}
	close(ch)
}

func readNumbers(consumer int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Printf("consumer %d received numbers: %d\n", consumer, v)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	go sendNumbers(ch)

	for i := 1; i < 3; i++ {
		wg.Add(1)
		go readNumbers(i, &wg, ch)
	}

	wg.Wait()

}
