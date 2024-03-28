package main

import "fmt"

//Thread 1
func main() { // Usado para multithreading
	ch := make(chan string) // Canal

	//Thread 2
	go func() {
		ch <- "Hello world!" // Canal preenchido (Não consegue atualizar ou por mais dados)
		ch <- "test"
	}()

	//Thread 1
	msg := <-ch // Canal esvazia (Permite que possa ser inserido novos dados)
	fmt.Println(msg)
}
