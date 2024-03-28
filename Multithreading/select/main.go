package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Msg string
	Id  int64
}

// Função com 2 canais, onde cada canal demora um tempo para encher (ter um valor).
func main() {
	kafka := make(chan Message)    //Simulando como se recebesse mensagem do Kafka
	rabbitMQ := make(chan Message) //Simulando como se recebesse mensagem do rabbitMQ
	var i int64 = 0

	go func() {
		for { //Loop infinito
			atomic.AddInt64(&i, 1) // Atomic para evitar concorrência (ter o mesmo valor duplicado).
			time.Sleep(time.Millisecond * 200)
			kafka <- Message{"Hello Kafka", i} // Ele carrega o canal Kafka antes do default.
		}
	}()
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Millisecond * 250)
			rabbitMQ <- Message{"Hello RabbitMQ", i}
		}
	}()

	// Select atua como um switch e cada case possui um Forever aguardando o canal Kafka ou o RabbitMQ
	// printar na tela.
	// Um exemplo real é enviar duas requests e dependendo de qual for mais rápida podemos continuar
	// com o código.
	for i := 0; i < 100; i++ {
		select {
		case msg1 := <-kafka:
			fmt.Printf("Received: %d - %s\n", msg1.Id, msg1.Msg)
		case msg2 := <-rabbitMQ:
			fmt.Printf("Received: %d - %s\n", msg2.Id, msg2.Msg)
		case <-time.After(time.Millisecond * 300):
			println("timeout")
			// default:
			// 	println("default")
		}
	}
}
