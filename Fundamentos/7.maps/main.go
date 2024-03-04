package main

import "fmt"

func main() {
	salario := map[string]int{"Wesley": 10, "Victor": 30, "Joao": 40}
	fmt.Println(salario["Victor"])
	delete(salario, "Joao") // delete registry
	salario["Roberto"] = 50 // add new registry
	fmt.Println(salario["Roberto"])

	for name, salary := range salario {
		fmt.Printf("O %s ganha %d.\n", name, salary)
	}

	for _, salary := range salario { //_ is a blank identifier (ignored)
		fmt.Printf("O salário é: %d\n", salary)
	}

	sal := make(map[string]int)
	sal2 := map[string]int{}

	sal = salario
	fmt.Println(sal)
	fmt.Println(sal2)
}
