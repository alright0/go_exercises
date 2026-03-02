package main

import (
	"fmt"
	"sync"
)

/*
Параллельный подсчет суммы массива.

Дан массив из 1 000 000 чисел.

Нужно:
1. разбить его на N частей
2. каждую часть обрабатывать в своей goroutine
3. вернуть общую сумму
*/

func getArray(maxNum int) []int {
	arr := make([]int, maxNum)
	for i := 0; i < maxNum; i++ {
		arr[i] = i + 1
	}
	return arr
}

func runCalc(arr []int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	var result int

	for i := 0; i < len(arr); i++ {
		result += arr[i]
	}

	ch <- result
}

func main() {
	var wg sync.WaitGroup

	var result int
	maxNum := 1_000_000
	arr := getArray(maxNum)

	chunkSize := 300_000
	workers := (maxNum + chunkSize) / chunkSize

	ch := make(chan int)

	for i := 0; i < workers; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if end > len(arr) {
			end = len(arr)
		}
		array := arr[start:end]

		wg.Add(1)
		go runCalc(array, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for partial := range ch {
		result += partial
	}

	fmt.Println("result:", result)
}
