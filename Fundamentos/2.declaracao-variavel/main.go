package main

const helloWorld = "Hello World"

var (
	isTrue  bool    = true
	number  int     = 10
	word    string  = "Victor"
	decimal float64 = 1.2
)

func main() {
	isTrue = true
	println(isTrue)
	println(number)
	println(word)
	println(decimal)

	a := "string"
	println(a)
	a = "string2"
	println(a)
}
