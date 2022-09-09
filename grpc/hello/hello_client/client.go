package main

import (
	"context"
	"fmt"
	v1 "github.com/guillospy92/utilsgolang/grpc/hello/hellopb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	fmt.Println("run grpc client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Printf("error failed connect %v", err)
	}

	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {
			log.Printf("error close %v", err)
		}
	}(cc)

	c := v1.NewHelloServiceClient(cc)

	// response streaming patter unary
	helloUnaryClientAndServer(c)

	// response streaming grpc server
	helloServerStreaming(c)

	// response streaming client
	helloClientStreaming(c)

	// response bidirectional stream
	HelloServiceGoodByeBidirectional(c)

}

func helloUnaryClientAndServer(service v1.HelloServiceClient) {
	fmt.Println("start patter unary")
	req := &v1.HelloRequest{
		Hello: &v1.Hello{
			FirstName: "Guillermo rom",
			Prefix:    "sr",
		},
	}

	res, err := service.Hello(context.Background(), req)

	if err != nil {
		log.Fatalf("error, calling hello Rpc %v", err)
	}

	log.Printf("response grpc server %v", res)
	fmt.Println("finish patter unary")
}

func helloClientStreaming(service v1.HelloServiceClient) {
	fmt.Println("client streaming")

	request := []*v1.HelloGodByeRequest{
		{
			Hello: &v1.Hello{
				FirstName: "Guillermo",
				Prefix:    "Sr",
			},
		},
		{
			Hello: &v1.Hello{
				FirstName: "Linda",
				Prefix:    "Sra",
			},
		},
		{
			Hello: &v1.Hello{
				FirstName: "Dara",
				Prefix:    "bebe",
			},
		},
	}

	stream, err := service.HelloGoodBye(context.Background())

	for _, r := range request {
		err := stream.Send(r)

		if err != nil {
			log.Printf("error send message server %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	goodBye, err := stream.CloseAndRecv()

	if err != nil {
		log.Printf("error send server %v", err)
	}

	log.Println("message deciweber", goodBye)

}

func helloServerStreaming(service v1.HelloServiceClient) {
	fmt.Println("start streaming grpc client")

	req := v1.HelloManyLanguagesRequest{
		Hello: &v1.Hello{
			FirstName: "Guillermo",
			Prefix:    "sr",
		},
	}

	restStream, err := service.HelloManyLanguages(context.Background(), &req)

	if err != nil {
		log.Printf("error comunicate grpc streaming %v", err)
		return
	}

	for {
		msg, err := restStream.Recv()

		if err == io.EOF {
			fmt.Println("terminate streaming reading")
			break
		}

		if err != nil {
			log.Printf("error reading streaming grpc %v", err)
		}

		log.Println("message", msg.HelloLanguage)
	}
}

func HelloServiceGoodByeBidirectional(service v1.HelloServiceClient) {

	request := []*v1.HelloGodByeRequest{
		{
			Hello: &v1.Hello{
				FirstName: "Guillermo",
				Prefix:    "SR",
			},
		},
		{
			Hello: &v1.Hello{
				FirstName: "Linda",
				Prefix:    "SRA",
			},
		},
		{
			Hello: &v1.Hello{
				FirstName: "Dara",
				Prefix:    "SRA",
			},
		},
		{
			Hello: &v1.Hello{
				FirstName: "Jesus",
				Prefix:    "SR",
			},
		},
	}

	stream, err := service.GoodBye(context.Background())

	if err != nil {
		log.Printf("error get client %v", err)
	}

	wait := make(chan struct{})

	go func() {
		for _, req := range request {
			err := stream.Send(req)

			if err != nil {
				log.Printf("error send message to server")
				return
			}
		}

		err := stream.CloseSend()

		if err != nil {
			log.Printf("error close message send, %v", err)
		}

	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("error recibied message %v", err)
			}

			fmt.Printf("response server %s \n", res.GetGoodbye())
		}

		close(wait)
	}()

	<-wait

	fmt.Println("Guillermo r")
}
