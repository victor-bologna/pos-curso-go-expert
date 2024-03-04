package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "Key", "Value")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	token := ctx.Value("Key")
	fmt.Println(token)
}
