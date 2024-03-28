package main

//Thread 1
func main() {
	ch := make(chan int) //Cria o canal
	//Thread 2
	go publish(ch) //Thread 2 atribui valores para o ch, deixando ele cheio com cada valor do for.
	//Thread 1
	reader(ch) // Thread 1 fica lendo o ch infinitamente e esvaziando-o com o println
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) //Thread 2 fecha o canal ao terminar o loop, evitando o deadlock.
}

func reader(ch chan int) {
	for v := range ch {
		println(v) //Thread 1 esvazia o canal printando o valor do ch.
	}
}
