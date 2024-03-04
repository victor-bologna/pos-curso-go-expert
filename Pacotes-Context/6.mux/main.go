package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HttpServe)
	mux.Handle("/blog", blog{})

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Victor!"))
	})

	go http.ListenAndServe(":8080", mux)
	go http.ListenAndServe(":8082", mux2)

	select {}
}

func HttpServe(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(b.title))
}
