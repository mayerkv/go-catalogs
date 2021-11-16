package main

import (
	"context"
	"fmt"
	"github.com/mayerkv/go-catalogs/grpc-service"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:9090", opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := grpc_service.NewCatalogsServiceClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := &grpc_service.CreateCatalogRequest{
		Catalog: &grpc_service.Catalog{
			Title: "321",
			Items: nil,
		},
	}

	res, err := client.CreateCatalog(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Id)
}
