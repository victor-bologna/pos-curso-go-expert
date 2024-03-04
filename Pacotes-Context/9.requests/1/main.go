package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	httpRequest := http.Client{Timeout: time.Second}
	response, err := httpRequest.Get("http://google.com/")
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
