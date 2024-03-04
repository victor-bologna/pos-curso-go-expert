package main

func sum(ca, cb int) int {
	return ca + cb
}

func sumA(ca, cb int) int {
	ca = 20
	return ca + cb
}

func sumAddress(ca, cb *int) int {
	return *ca + *cb
}

func sumAddressB(ca, cb *int) int {
	*ca = 20
	return *ca + *cb
}

func main() {
	ca, cb := 10, 20

	println("sum ", sum(ca, cb))
	println("sumA ", sumA(ca, cb))
	println("ca ", ca)

	println("sumAddress ", sumAddress(&ca, &cb))
	println("sumAddressB ", sumAddressB(&ca, &cb))
	println("ca ", ca)
}
