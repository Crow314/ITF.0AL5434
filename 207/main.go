package main

import (
	"fmt"
	"sync"
)

const BUFFER_SIZE int = 8

type buffer struct {
	buf  chan int
	rmux sync.Mutex
	wmux sync.Mutex
}

func main() {
	var wg sync.WaitGroup

	b := buffer{buf: make(chan int, BUFFER_SIZE)}

	wg.Add(1)
	go thread1(&b, &wg)

	wg.Add(1)
	go thread2(&b, &wg)

	wg.Add(1)
	go thread3(&b, &wg)

	wg.Wait()
}

func put(b *buffer, x []int) {
	defer b.wmux.Unlock()
	b.wmux.Lock()

	for _, v := range x {
		b.buf <- v
	}
}

func get(b *buffer, n int) []int {
	defer b.rmux.Unlock()

	r := make([]int, n)

	b.rmux.Lock()

	for i := 0; i < n; i++ {
		x := <-b.buf
		r[i] = x
	}

	return r
}

func thread1(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()

	x := []int{1, 3, 5}

	fmt.Println("Thread 1: put(3,{1,3,5})")
	put(b, x)
}

func thread2(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()

	x := []int{2, 4, 6}

	fmt.Println("Thread 2: put(3,{2,4,6})")
	put(b, x)
}

func thread3(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		x := get(b, 3)
		fmt.Printf("thread3(): get() %v\n", x)
	}
}
