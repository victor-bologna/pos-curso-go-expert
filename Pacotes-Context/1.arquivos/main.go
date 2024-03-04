package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// criar arquivo

	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//escrever no arquivo

	//tamanho, err := f.WriteString("Hello, World!")
	tamanho, err := f.Write([]byte("Hello, World!"))
	if err != nil {
		panic(err)
	}

	fmt.Println("Arquivo criado!", tamanho, "bytes.")

	//leitura

	//arquivo := os.Open("arquivo.txt")
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Lendo arquivo: ", string(arquivo))

	f.Close()

	//Leitura de pouco em pouco via Buffer

	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 5)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		// Lendo a partir da nova sessão do buffer
		fmt.Printf("length = %d capacity = %d %v\n", len(buffer[:n]), cap(buffer[:n]), string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
