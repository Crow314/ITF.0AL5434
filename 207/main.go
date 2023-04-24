package main

import (
	"fmt"
	"sync"
)

const BUFFER_SIZE int = 8

func main() {
	var wg sync.WaitGroup

	b := make(chan int, BUFFER_SIZE)

	wg.Add(1)
	go thread1(b, &wg)

	wg.Add(1)
	go thread2(b, &wg)

	wg.Add(1)
	go thread3(b, &wg)

	wg.Wait()
}

func put(b chan<- int, x []int) {
	for _, v := range x {
		b <- v
	}
}

func get(b <-chan int, n int) []int {
	r := make([]int, n)

	for i := 0; i < n; i++ {
		x := <-b
		r[i] = x
	}

	return r
}

func thread1(b chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	x := []int{1, 3, 5}

	fmt.Println("Thread 1: put(3,{1,3,5})")
	put(b, x)
}

func thread2(b chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	x := []int{2, 4, 6}

	fmt.Println("Thread 2: put(3,{2,4,6})")
	put(b, x)
}

func thread3(b <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		x := get(b, 3)
		fmt.Printf("thread3(): get() %v\n", x)
	}
}
