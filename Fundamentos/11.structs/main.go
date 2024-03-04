package main

import "fmt"

type Endereco struct {
	Logradouro string
	Bairro     string
	Numero     int
	Cidade     string
	UF         string
}

type Cliente struct {
	Name     string
	Idade    int
	Ativo    bool
	Endereco // Composição (preferivel "has a" relationship)
	// Address Endereco // Criando propriedade do tipo Endereco
}

type Empresa struct {
	Name  string
	Ativo bool
}

type TipoPessoa interface {
	changeAtivo()
}

func ativarDesativar(tipoPessoa TipoPessoa) {
	tipoPessoa.changeAtivo()
}

func (c Cliente) changeAtivo() {
	c.Ativo = !c.Ativo
	fmt.Println(c.Ativo)
	fmt.Println(fmt.Sprintf("O status do cliente foi alterado para %t", c.Ativo))
}

func (e Empresa) changeAtivo() {
	e.Ativo = !e.Ativo
	fmt.Println("Status da empresa alterado")
}

func main() {
	victor := Cliente{
		"Victor",
		26,
		true,
		Endereco{ // Composição é diferente de criar properiedade.
			"a",
			"b",
			1,
			"c",
			"d",
		},
	}
	victor.Ativo = false
	fmt.Println(victor)
	fmt.Println(victor.Ativo)
	fmt.Printf("O cliente é %s e tem %d anos de idade e esta ativo? %t\n", victor.Name, victor.Idade, victor.Ativo)
	fmt.Println(victor.Cidade)
	fmt.Println(victor.Endereco.Bairro)
	victor.changeAtivo()

	empresa := Empresa{
		Name:  "Empresa",
		Ativo: false,
	}

	ativarDesativar(empresa)
	ativarDesativar(victor)
}
