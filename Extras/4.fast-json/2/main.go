package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type User struct {
	Name string
	Age  int
}

func main() {
	var p fastjson.Parser
	jsonData := `{"user": {"name": "John Doe", "age": 36}}`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	user := v.GetObject("user")
	fmt.Printf("User name: %s\n", user.Get("name"))
	fmt.Printf("User age: %s\n", user.Get("age"))

	// ou passar para struct

	userJson := v.GetObject("user").String()
	var userStruct User
	json.Unmarshal([]byte(userJson), &userStruct)
	fmt.Println(userStruct)
}
