package main

import (
	"context"
	"fmt"
	v1 "github.com/guillospy92/utilsgolang/grpc/hello/hellopb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type server struct {
	v1.UnimplementedHelloServiceServer
}

func (*server) Hello(_ context.Context, req *v1.HelloRequest) (*v1.HelloResponse, error) {

	firstName := req.GetHello().GetFirstName()
	prefix := req.Hello.GetPrefix()

	customerHello := "welcome !" + prefix + " " + firstName

	return &v1.HelloResponse{
		CustomHello: customerHello,
	}, nil
}

func (*server) HelloManyLanguages(req *v1.HelloManyLanguagesRequest, stream v1.HelloService_HelloManyLanguagesServer) error {
	fmt.Println("call HelloManyLanguages", req)

	lang := [5]string{"hello chile", "Hello argentina", "Ni hao", "Hello columbia", "hello peru"}
	firstName := req.GetHello().GetFirstName()
	prefix := req.Hello.GetPrefix()

	for _, hi := range lang {
		helloLanguage := hi + " " + prefix + " " + firstName

		res := &v1.HelloManyLanguagesResponse{
			HelloLanguage: helloLanguage,
		}

		err := stream.Send(res)
		if err != nil {
			return err
		}

		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func (*server) HelloGoodBye(stream v1.HelloService_HelloGoodByeServer) error {

	goodBye := "GoodBye guys: "

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&v1.HelloGodByeResponse{
				Goodbye: goodBye,
			})
		}

		if err != nil {
			log.Printf("error streaming grpc server %v", err)
		}

		firstName := req.Hello.GetFirstName()
		prefix := req.Hello.GetPrefix()

		goodBye += prefix + " " + firstName + " "
	}
}

func (*server) GoodBye(stream v1.HelloService_GoodByeServer) error {
	fmt.Println("invoke bidirectional function")
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Printf("error read grpc stream %v", err)
		}

		firstName := req.GetHello().GetFirstName()
		prefix := req.GetHello().GetPrefix()

		bye := "Good Bye " + prefix + " " + firstName + " :("

		err = stream.Send(&v1.HelloGodByeResponse{
			Goodbye: bye,
		})

		if err != nil {
			log.Printf("error send message backend")
		}
	}
}

func main() {
	list, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed serve %v", err)
	}

	s := grpc.NewServer()

	v1.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(list); err != nil {
		log.Fatalf("failed init server grpc %v", err)
	}
}
