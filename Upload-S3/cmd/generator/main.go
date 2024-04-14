package main

import (
	"fmt"
	"os"
)

func main() {
	for v := 0; v < 10; v++ {
		f, err := os.Create(fmt.Sprintf("./tmp/file%d.txt", v))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		_, err = f.Write([]byte("Hello world"))
		if err != nil {
			panic(err)
		}
	}
}
