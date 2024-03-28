package main

import (
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		println("%d: A task %s is running", i, name)
		time.Sleep(time.Second)
		wg.Done()
	}
}

// Thread 1
func main() {
	waitGroup := sync.WaitGroup{} // Segura a apliação até que o número especificado seja zerado.
	waitGroup.Add(24)             // 25 operações wg.Done()
	//Thread 2 - Roda em backgound
	go task("A", &waitGroup)
	//Thread 3 - Roda em background
	go task("B", &waitGroup)
	//Se não colocar algo ele sai antes de terminar as tasks.
	go func() { //Função anônima
		for i := 0; i < 4; i++ {
			println("%d: A task annonymous is running", i)
			time.Sleep(time.Second)
			waitGroup.Done()
		}
	}()
	waitGroup.Wait() // Aguarda a execução dos recursos e continua o programa.
}
