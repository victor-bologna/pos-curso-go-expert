package main

func main() {
	a := 1
	b := 2
	c := 3

	if a > b && b > c {
		println("a > b && b > c")
	} else {
		println("else")
	}

	if a == b || b == c {
		println("a == b || b == c")
	}

	switch a {
	case 1:
		println("1")
		break
	case 2:
		println("2")
		break
	case 3:
		println("3")
		break
	default:
		println("N/A")
	}
}
