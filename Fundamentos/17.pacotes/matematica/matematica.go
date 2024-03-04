package matematica

import "fmt"

func Soma[T int | float64](a, b T) T { // Se a primeira letra do nome da função for maiuscula,
	// ela será exportada para fora do pacote, se for minuscula será local.
	return a + b
}

var A int = 10

type Carro struct {
	Marca string
}

func (carro Carro) Andar() {
	fmt.Printf("O carro da marca : %v andou\n", carro.Marca)
}
