package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"number"`
	Saldo  int `json:"value"`
}

func main() {
	conta := Conta{100, 200}
	res, err := json.Marshal(conta) // Marshal -> Transforma em JSON
	if err != nil {
		panic(err)
	}

	println(string(res)) //Uso do string pois a mensagem fica em Bytes

	err = json.NewEncoder(os.Stdout).Encode(conta) //Prepara a mensagem para Json
	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"number":100,"value":200}`)
	var contaX Conta

	err = json.Unmarshal(jsonPuro, &contaX) // UnMarshal transforma o json em conta
	if err != nil {
		panic(err)
	}

	fmt.Println(contaX)
}
