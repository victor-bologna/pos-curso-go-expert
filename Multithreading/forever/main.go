package main

//Thread 1
func main() {
	forever := make(chan bool) // Channel vazio
	// Segura o processo de pé no main
	//Thread 2 enche o forever preenchendo true
	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true
	}()
	<-forever // Esperando ficar cheio para esvaziar

	// Se tentar preencher o forever na mesma thread ele da deadlock

	/*
		test := make(chan bool)

		test <- true // deadlock pq ele tem que rodar em outra go routine

		<-test
	*/
}
