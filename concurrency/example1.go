package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id  int
	num int
}

type Result struct {
	job Job
	sum int
}

func sumDigit(n int) int {
	sum := 0
	for n > 0 {
		sum = sum + n%10
		n = n / 10
	}
	time.Sleep(3 * time.Second) // To show some delay
	return sum
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func createJobs(jobSize int) {
	for i := 0; i < jobSize; i++ {
		jobs <- Job{
			id:  i,
			num: rand.Intn(999),
		}
	}
	close(jobs)
}

func createWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		res := Result{
			job: job,
			sum: sumDigit(job.num),
		}
		results <- res
	}
}

func createWorkerPool(size int) {
	var wg sync.WaitGroup
	for i := 0; i < size; i++ {
		wg.Add(1)
		go createWorker(&wg)
	}
	wg.Wait()
	close(results)
}

func showResult(done chan<- bool) {
	for result := range results {
		fmt.Printf("Sum of the digits of the num: %d is %d \n", result.job.num, result.sum)
	}
	done <- true
}

func main() {
	startTime := time.Now()

	jobSize := 100
	go createJobs(jobSize)

	done := make(chan bool)
	go showResult(done)

	poolSize := 20
	go createWorkerPool(poolSize)

	<-done

	endTime := time.Now()
	timeTaken := endTime.Sub(startTime)
	fmt.Println("Time taken to complete the job:", timeTaken)
}
