package main

import "fmt"

func main() {
	numberArray := []int{2, 4, 6, 8, 10}
	fmt.Printf("length = %d capacity = %d %v\n", len(numberArray), cap(numberArray), numberArray)
	fmt.Printf("length = %d capacity = %d %v\n", len(numberArray[:0]), cap(numberArray[:0]), numberArray[:0])
	fmt.Printf("length = %d capacity = %d %v\n", len(numberArray[:3]), cap(numberArray[:3]), numberArray[:3])
	fmt.Printf("length = %d capacity = %d %v\n", len(numberArray[3:]), cap(numberArray[3:]), numberArray[3:])

	numberArray = append(numberArray, 20)

	fmt.Printf("length = %d capacity = %d %v\n", len(numberArray[:3]), cap(numberArray[:3]), numberArray[:3])
	fmt.Printf("length = %d capacity = %d %v\n", len(numberArray[3:]), cap(numberArray[3:]), numberArray[3:])
	fmt.Printf("length = %d capacity = %d %v\n", len(numberArray), cap(numberArray), numberArray)
}
