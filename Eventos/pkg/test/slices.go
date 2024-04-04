package main

import "fmt"

// Revisitando slices
func main() {
	evento := []string{"teste", "teste1", "teste2", "teste3"}
	fmt.Println(evento[:2])
	fmt.Println(append(evento[:1], evento[2:]...))

}
