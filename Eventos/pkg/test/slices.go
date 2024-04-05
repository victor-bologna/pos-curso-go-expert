package main

import (
	"fmt"

	secret "github.com/victor-bologna/pos-curso-go-expert-secret-module/pkg/events"
)

// Revisitando slices
func main() {
	evento := []string{"teste", "teste1", "teste2", "teste3"}
	fmt.Println(evento[:2])
	fmt.Println(append(evento[:1], evento[2:]...))
	a := secret.EventDispatcher{}
	fmt.Println(a)
}
