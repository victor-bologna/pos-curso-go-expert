package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

func main() {
	debug.SetMemoryLimit(10 * 1024 * 1024) // 10 MiB

	// Rodar GODEBUG=gctrace=1 go run main.go
	// Rodar GODEBUG=gctrace=1 GOGC=300 go run main.go

	// debug.SetGCPercent(-1) // Desativa GC automático
	// Ajustar o percentual do GC (por exemplo, para 300%)
	// debug.SetGCPercent(300)

	// Função para alocar memória
	allocateMemory := func(size int) []byte {
		return make([]byte, size)
	}

	// Alocando memória para observar o comportamento
	for i := 0; i < 10; i++ {
		allocateMemory(20 * 1024 * 1024)
		time.Sleep(time.Second)
	}

	// Exibindo o uso de memória
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v Mib\n", m.Alloc/1024/1024)
	fmt.Printf("TotalAlloc = %v MiB\n", m.TotalAlloc/1024/1024)
	fmt.Printf("Sys = %v MiB\n", m.Sys/1024/1024)
	fmt.Printf("NumGC = %v\n", m.NumGC)
}
