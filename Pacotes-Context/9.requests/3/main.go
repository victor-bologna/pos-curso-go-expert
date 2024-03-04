package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	httpRequest := http.Client{} // Cria o objeto
	request, err := http.NewRequest("GET", "https://google.com/", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Accept", "application/json")
	response, err := httpRequest.Do(request) // Executa a request
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
