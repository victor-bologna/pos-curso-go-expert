package main

import (
	"errors"
	"fmt"
)

func main() {
	println(sumNormal(1, 2))
	println(sumSimplified(1, 2))
	println(sumDoubleReturn(1, 2))
	println(sumDoubleReturn(3, 2))

	num, error := sumError(1, 2)
	if error != nil {
		fmt.Println(error)
	}
	println(num)

	num2, error2 := sumError(5, 6)
	if error2 != nil {
		fmt.Println(error2)
	}
	println(num2)

	num, error = sumError(1, 4)
	fmt.Println(num, error)

	num, test, error := sumTripleReturn(1, 3)
	fmt.Println(num, error, test)

	fmt.Println(sumTripleReturn(2, 3))

}

func sumNormal(a int, b int) int {
	return a + b
}

func sumSimplified(a, b int) int {
	return a + b
}

func sumDoubleReturn(a, b int) (int, bool) {
	if a+b > 3 {
		return a + b, true
	}
	return 0, false
}

func sumError(a, b int) (int, error) {
	if a+b >= 5 {
		return a + b, nil
	}
	return 0, errors.New("O valor não pode ser menor que 5.")
}

func sumTripleReturn(a, b int) (int, bool, error) {
	return a + b, true, nil
}
