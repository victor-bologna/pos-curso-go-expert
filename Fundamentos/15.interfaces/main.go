package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello world!"

	showType(x)
	showType(y)

	println(x.(int))

	res, ok := x.(int)
	res2, ok2 := x.(string)
	res3 := x.(string)
	fmt.Printf("O valor de res é %v é o resultado é %v\n", res, ok)
	fmt.Printf("O valor de res é %v é o resultado é %v\n", res2, ok2)
	fmt.Printf("O valor de res é %v\n", res3)
}

func showType(x interface{}) {
	fmt.Printf("O tipo é de %T tem o valor %v\n", x, x)
}
