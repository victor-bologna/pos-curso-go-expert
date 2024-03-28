package main

func main() {
	ch := make(chan string, 2) // Segundo parâmetro mostra que é um buffer
	// (Não recomendado pois ocupa memória)
	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
