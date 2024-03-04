package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//Pegando o request
	response, err := http.Get("https://google.com/")
	if err != nil {
		panic(err)
	}
	fmt.Println("Executando test()")
	test()
	defer response.Body.Close() // Defer atrasa a execução para ultima linha da função.
	fmt.Println("Executado test()")
	//Lendo a resposta do request
	reader, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	defer test2()
	fmt.Println(reader)
	//fmt.Println(string(reader))
}

func test() {
	defer fmt.Println("Executado defer no Body Close")
	fmt.Println("Teste")
}

func test2() {
	fmt.Println("Teste defer na func")
}
