package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string
	Logradouro  string
	Complemento string
	Bairro      string
	Localidade  string
	Uf          string
	Ibge        string
	Gia         string
	Ddd         string
	Siafi       string
}

func main() {
	for _, cep := range os.Args[1:] {
		reader, err := getCEP(cep)
		cep := jsonToStruct(reader, err)
		saveFile(cep)
	}
}

func getCEP(cep string) ([]byte, error) {
	response, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		println("Cannot access url CEP:", cep)
		panic(err)
	}
	defer response.Body.Close()

	return getResponseBody(response, err)
}

func getResponseBody(response *http.Response, err error) ([]byte, error) {
	reader, err := io.ReadAll(response.Body)
	if err != nil {
		println("Error while reading request response")
		panic(err)
	}
	return reader, err
}

func jsonToStruct(reader []byte, err error) ViaCEP {
	var cep ViaCEP
	err = json.Unmarshal(reader, &cep)
	if err != nil {
		println("Error while parsing payload to Struct", reader)
		panic(err)
	}
	return cep
}

func saveFile(cep ViaCEP) {
	cepFile, err := os.Create("cep.txt")
	if err != nil {
		println("Could not create file")
		panic(err)
	}

	_, err = cepFile.WriteString(fmt.Sprintf(
		"CEP: %s\nLogradouro: %s\nComplemento: %s\nBairro: %s\nCidade: %s\nUF: %s",
		cep.Cep, cep.Logradouro, cep.Complemento, cep.Bairro, cep.Localidade, cep.Uf))
	if err != nil {
		println("Error while writing on file")
		panic(err)
	}

	println("File created!")
}
