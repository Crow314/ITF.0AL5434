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
	go threadA(b, &wg)

	wg.Add(1)
	go threadB(b, &wg)

	wg.Wait()
}

func put(b chan<- int, x int) {
	b <- x
}

func get(b <-chan int) int {
	return <-b
}

func threadA(b chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		x := i
		fmt.Printf("threadA(): put( %d )\n", x)
		put(b, x)
	}
}

func threadB(b <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		x := get(b)
		fmt.Printf("threadB(): get() %d\n", x)
	}
}
