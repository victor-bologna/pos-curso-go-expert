package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/victor-bologna/pos-curso-go-expert/matematica"
)

func main() {
	soma := matematica.Soma(10, 20)
	fmt.Println("O valor da soma é: ", soma)
	fmt.Println(matematica.A)

	fiat := matematica.Carro{
		Marca: "Fiat",
	}

	fmt.Println("A marca do carro é :", fiat.Marca)
	fiat.Andar()

	fmt.Println(uuid.New())
}

// letra inicial minuscula = private
// letra inicial maiuscula = public
