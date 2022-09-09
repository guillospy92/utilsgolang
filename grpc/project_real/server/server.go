package main

import (
	"context"
	"fmt"
	"github.com/guillospy92/utilsgolang/grpc/mongo"
	"github.com/guillospy92/utilsgolang/grpc/project_real/protos"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
)

type Product struct {
	clientMongo *mongo.Collection
	product_pb.UnimplementedProductServiceServer
}

type product struct {
	ID    string  `bson:"_id,omitempty"`
	Name  string  `bson:"name"`
	Price float64 `bson:"price"`
}

func (pro *Product) CreateProduct(c context.Context, p *product_pb.CreateProductRequest) (*product_pb.CreateProductResponse, error) {

	product := product{
		Name:  p.Product.Name,
		Price: p.Product.Price,
	}

	register, err := pro.clientMongo.InsertOne(c, product)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error %v", err))
	}

	product.ID = register.InsertedID.(string)

	return &product_pb.CreateProductResponse{
		Product: &product_pb.PerProduct{
			Id:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		},
	}, nil
}

func main() {
	// set logs
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("check product server grpc")

	certFile := "./certs/server.crt"
	keyFile := "./certs/server.pem"

	cred, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)

	if sslErr != nil {
		log.Println("error reading certificate %w", sslErr)
		return
	}

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Println("error server, %w", err)
		return
	}

	opts := grpc.Creds(cred)
	s := grpc.NewServer(opts)

	clientMongo, err := mongo_project.NewConnectMongoClient()

	if err != nil {
		log.Println("error connect mongo db , %w", err)
		return
	}

	product_pb.RegisterProductServiceServer(s, &Product{clientMongo: clientMongo})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Printf("error connecting server %v", err)
		}
	}()

	// wait ctrl + x to exit
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)

	// block until
	<-ch

	fmt.Println("stop server")

	s.Stop()

}
