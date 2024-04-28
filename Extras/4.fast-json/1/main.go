package main

import (
	"fmt"
	"log"

	"github.com/valyala/fastjson"
)

func main() {
	var p fastjson.Parser
	jsonData := `{"foo":"bar", "num": 123, "bool": true, "arr": [1,2,3]}`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("foo=%s\n", v.GetStringBytes("foo"))
	fmt.Printf("num=%d\n", v.GetInt("num"))
	fmt.Printf("bool=%v\n", v.GetBool("bool"))

	a := v.GetArray("arr")
	for i, v := range a {
		fmt.Printf("indice=%d valor=%s\n", i, v)
	}

	var sc fastjson.Scanner

	sc.Init(jsonData)

	for sc.Next() {
		fmt.Printf("%s\n", sc.Value())
	}
	if err := sc.Error(); err != nil {
		log.Fatalf("unexpected Error: %s", err)
	}
}
