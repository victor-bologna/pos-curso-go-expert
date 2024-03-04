package main

import "fmt"

func main() {
	fmt.Println(sum(1, 3, 9, 243, 92, 3282, 183, 291, 2381, 392318, 182381, 1823, 181, 238231, 1823))
}

func sum(numbers ...int) int {
	total := 0
	for _, numbers := range numbers {
		total += numbers
	}
	return total
}
