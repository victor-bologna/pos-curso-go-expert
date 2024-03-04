package main

type MyNumber int

type Number interface { //Isso é uma constraint, elas podem ser usdas como variável para generics
	~int | ~float64 // ~ serve para considerar o tipo da variável (no caso MyNumber)
}

func soma[T Number](m map[string]T) T {
	var somaTotal T
	for _, v := range m {
		somaTotal += v
	}
	return somaTotal
}

func main() {
	m := map[string]int{"Victor": 100, "Maria": 200, "João": 300}
	m2 := map[string]float64{"Victor": 100.4, "Maria": 200.3, "João": 300.2}
	m3 := map[string]MyNumber{"Victor": 100, "Maria": 200, "João": 300}

	println(soma(m))
	println(soma(m2))
	println(soma(m3))
}
