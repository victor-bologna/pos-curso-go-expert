package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":8080", nil)
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Requisição iniciada.")
	defer log.Println("Requisição finalizada")
	select {
	case <-ctx.Done():
		log.Println("Requisição cancelada pelo cliente.")
		w.Write([]byte("Requisição cancelada."))
		return
	case <-time.After(5 * time.Second):
		log.Println("Requisição concluida.")
		w.Write([]byte("Requisição finalizada com sucesso."))
	}
}
