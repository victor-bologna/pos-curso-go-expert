package main

func main() {
	a := 10
	println(a)
	var b *int = &a
	println(b)
	*b = 20
	println(a, b)
	ponteiro := &b
	println(ponteiro)
	println(&a, &b, &ponteiro)
	// Variavel -> Memória -> Valor
}
