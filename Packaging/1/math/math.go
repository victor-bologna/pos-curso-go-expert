package math

type math struct {
	number1 int
	number2 int
}

func NewMath(number1, number2 int) *math {
	return &math{number1: 1, number2: 2}
}

func (math math) Add() int {
	return math.number1 + math.number2
}
