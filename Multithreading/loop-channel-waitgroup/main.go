package main

import (
	"sync"
)

// Thread 1
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)
	go publish(ch)
	go reader(ch, &wg)
	wg.Wait()
	println("Wait group done.")
}

func reader(ch chan int, waitGroup *sync.WaitGroup) {
	for v := range ch {
		println(v)
		waitGroup.Done()
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
