package main

import (
	"fmt" //for println
	"sync" //for mutex like timers
)
//amd ryzen 3700U, governor mode "ondemand"
//time with on worker:  99% cpu 4:26.98 total
//time with 4 workers: 241% cpu 1:54.50 total
//time with 8 workers: 312% cpu 2:08.18 total
func main(){
	jobs := make(chan int, 100) //buffered chanel of ints
	result := make(chan int, 100)
	var wg sync.WaitGroup //the mutex
	wg.Add(1)

	fibo := 50 //how many numbers we want to calculate

	go worker(jobs, result) //parallel workers
	go worker(jobs, result) //parallel workers
	go worker(jobs, result) //parallel workers
	go worker(jobs, result) //parallel workers

	for i:=1; i<=fibo; i++ {
		jobs <- i
	}
	close(jobs)

	go func() {
		for i:=1; i<=fibo; i++ {
			fmt.Printf("Number is %v and the result: %+v\n", i, <-result)
		}
		wg.Done()
	}()

	wg.Wait()
	return
}

func worker( jobs <-chan int, result chan<- int ){
	for n := range jobs {
		result <-fib(n)
	}
	fmt.Println("oooohwheee")
}

func fib(num int) int{
	if num <= 1 {
		return num
	}

	return fib(num-1) + fib(num-2)
}
