package main

import "fmt"

type Pessoa struct {
	Name string
}

func (pessoa Pessoa) andou() {
	pessoa.Name = "Victor Bologna"
	fmt.Printf("A pessoa %v andou.\n", pessoa.Name)
}

type Conta struct {
	Saldo int
}

func NewConta() *Conta { // Cria uma struct com endereço apontado para a conta, assim consegue alterar valores através de functions
	return &Conta{Saldo: 0}
}

func (conta Conta) simular(valor int) int { // Passa uma cópia da struct
	conta.Saldo += valor
	return conta.Saldo
}

func (conta *Conta) alterar(valor int) int { // Passa o endereço de memória da struct
	conta.Saldo += valor
	return conta.Saldo
}

func main() {
	victor := Pessoa{
		Name: "Victor",
	}
	victor.andou()
	fmt.Println(victor.Name)

	conta := Conta{
		Saldo: 50,
	}
	println(conta.simular(100))
	println(conta.Saldo)
	println(conta.alterar(100))
	println(conta.Saldo)
}
