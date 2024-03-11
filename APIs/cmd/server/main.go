package main

import "github.com/victor-bologna/pos-curso-go-expert-apis/configs"

func main() {
	config := configs.LoadConfig(".")
	println(config)
}
