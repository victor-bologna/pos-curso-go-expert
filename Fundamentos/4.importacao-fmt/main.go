package main

import "fmt"

type ID int

var decimal float64 = 1.2
var number ID = 2

func main() {
	fmt.Printf("O tipo de decimal é %T", decimal)
	fmt.Printf("O tipo de ID é %T", number)
}
