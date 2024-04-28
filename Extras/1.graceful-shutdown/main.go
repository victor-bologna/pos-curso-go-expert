package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second)
		w.Write([]byte("Hello World!"))
	})

	go func() {
		if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1) // Sinal do OS
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	// Quando notificar o canal será esvaziado e o programa acaba
	<-stop // Bloqueia a thread

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Shutting down gracefully server...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown server: %v\n", err)
	}
	fmt.Println("Server stopped.")
}
