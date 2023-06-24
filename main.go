package main

import "fmt"

func fibonacci(max int, ch chan int) {

	fibonacciSlice := make([]int, max)
	fibonacciSlice[0] = 0
	fibonacciSlice[1] = 1

	ch <- fibonacciSlice[0]
	ch <- fibonacciSlice[1]

	for i := 2; i < max; i++ {
		fibonacciSlice[i] = fibonacciSlice[i-1] + fibonacciSlice[i-2]
		ch <- fibonacciSlice[i]
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go fibonacci(10, ch)

	for i := range ch {
		fmt.Println(i)
	}

}
