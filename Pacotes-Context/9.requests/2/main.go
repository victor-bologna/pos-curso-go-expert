package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	httpRequest := http.Client{Timeout: time.Second}
	jsonBody := bytes.NewBuffer([]byte(`{"nome": "Victor", "age": 25}`)) //Buffer é parte do io.Read
	response, err := httpRequest.Post("https://google.com/", "application/json", jsonBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	io.CopyBuffer(os.Stdout, response.Body, nil)
}
