package main

import (
	"github.com/mayerkv/go-catalogs/domain"
	"github.com/mayerkv/go-catalogs/grpc-service"
	"github.com/mayerkv/go-catalogs/repository"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	catalogRepository := repository.NewInMemoryCatalogRepository()
	catalogService := domain.NewCatalogService(catalogRepository)
	srv := grpc_service.NewCatalogsServiceServerImpl(catalogService)

	if err := runGrpcServer(srv); err != nil {
		log.Fatalln(err)
	}
}

func runGrpcServer(srv grpc_service.CatalogsServiceServer) error {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	grpc_service.RegisterCatalogsServiceServer(grpcServer, srv)

	return grpcServer.Serve(lis)
}
