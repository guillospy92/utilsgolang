package main

import (
	"context"
	"github.com/guillospy92/utilsgolang/grpc/project_real/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {

	certFile := "../server/certs/ca.crt"

	cred, sslErr := credentials.NewClientTLSFromFile(certFile, "")

	if sslErr != nil {
		log.Printf("error get ssl connection %v", sslErr)
		return
	}

	opt := grpc.WithTransportCredentials(cred)

	cc, err := grpc.Dial("localhost:50051", opt)

	if err != nil {
		log.Printf("error failed connect %v", err)
		return
	}

	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {
			log.Printf("error close %v", err)
		}
	}(cc)

	clientService := product_pb.NewProductServiceClient(cc)

	product, err := clientService.CreateProduct(context.Background(), &product_pb.CreateProductRequest{
		Product: &product_pb.PerProduct{
			Name:  "Banana",
			Price: 20.09,
		},
	})

	if err != nil {
		log.Fatalf("error create product, %v %v", product, err)
	}
}
