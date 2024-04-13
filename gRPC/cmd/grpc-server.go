package main

import (
	"database/sql"
	"net"

	"github.com/pos-curso-go-expert-grpc/internal/database"
	"github.com/pos-curso-go-expert-grpc/internal/pb"
	"github.com/pos-curso-go-expert-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./database.db")
	if err != nil {
		panic(err)
	}

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
