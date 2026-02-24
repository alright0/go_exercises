package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Есть список из 20 задач (например, просто числа 1–20).

Нужно:
Запустить пул из 3 воркеров
Каждая задача «обрабатывается» 500ms
Воркеры читают задачи из канала jobs
Результаты отправляются в канал results
После обработки всех задач программа корректно завершается
*/

func initJobs(maxJobs int) []int {
	jobs := make([]int, maxJobs)
	for i := 0; i < maxJobs; i++ {
		jobs[i] = i + 1
	}
	return jobs
}

func doJob(worker int, jobsCh chan int, resultCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobsCh {
		time.Sleep(time.Millisecond * 500)
		resultCh <- fmt.Sprintf("worker %d done job %d", worker, job)
	}
}

func sendJobs(jobs []int, jobsCh chan int) {
	defer close(jobsCh)

	for _, job := range jobs {
		jobsCh <- job
	}
}

func main() {
	maxJobs := 20
	poolSize := 3
	jobs := initJobs(maxJobs)
	jobsCh := make(chan int)
	resultCh := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go doJob(i, jobsCh, resultCh, &wg)
	}

	go sendJobs(jobs, jobsCh)

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for v := range resultCh {
		fmt.Println(v)
	}
}
