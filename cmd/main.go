package main

import (
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/tinchourteaga/go-grpc-product-svc/internal/pb"
	"github.com/tinchourteaga/go-grpc-product-svc/internal/product"
	"github.com/tinchourteaga/go-grpc-product-svc/pkg/config"
	"github.com/tinchourteaga/go-grpc-product-svc/pkg/db"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting server...")
	err := config.Load()
	if err != nil {
		log.Error().Msg("config loading: " + err.Error())
	}

	listener, err := net.Listen("tcp", viper.GetString("PORT"))
	if err != nil {
		log.Fatal().Msg("port listening failed: " + err.Error())
	}

	grpcServer := grpc.NewServer()
	connector := db.NewDatabaseConnection()
	repo := product.NewRepository(connector)
	service := product.NewService(repo)
	handler := product.NewProduct(service)

	fmt.Println("Auth service listening on: " + viper.GetString("PORT"))

	pb.RegisterProductServiceServer(grpcServer, handler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal().Msg("failed to serve: " + err.Error())
	}
}
