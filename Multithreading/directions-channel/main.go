package main

import "fmt"

func recebe(nome string, ch chan<- string) { // Direita <- siginifica que o canal somente
	//vai receber/preencher
	ch <- nome // preenche o canal com string nome
}

func ler(dataCh <-chan string) { // Esquerda <- Esvaziar/entregar o resultado do canal
	fmt.Println(<-dataCh) //esvazia o canal lendo o string
}

func main() {
	ch := make(chan string) // cria o canal
	go recebe("Hello", ch)  // Thread 2 preenche o canal
	ler(ch)                 // Thread 1 le o canal
}
