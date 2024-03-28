package main

import (
	"fmt"
	"time"
)

func work(workerId int, dataCh chan int) {
	for v := range dataCh {
		fmt.Printf("Worker id: %d working on data: %d\n", workerId, v)
		time.Sleep(time.Millisecond * 250)
	}
}

func main() {
	data := make(chan int)

	for i := 0; i < 90; i++ { //Inicializa os workers
		go work(i, data)
	}

	for i := 0; i < 100; i++ { //Processa os dados (Cada worker processa um dado)
		data <- i
	}
}
