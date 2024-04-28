package main

import "fmt"

func myPanic() {
	panic("Something went wrong...")
}

func myPanic2() {
	panic("Something went wrong 2...")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if r == "Something went wrong..." {
				fmt.Println("myPanic recovered.")
			} else {
				fmt.Println("myPanic2 recovered.")
			}
		}
	}()
	myPanic()
}
