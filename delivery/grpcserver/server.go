package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"mdhesari/simple-golang-grpc/contract/hello"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("Port", 5001, "The server port")
)

type server struct {
	hello.UnimplementedGreeterServer
}

func (s server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloRespone, error) {
	log.Printf("received %v", in.GetName())
	return &hello.HelloRespone{Message: "Hellooooo" + in.GetName()}, nil
}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("net could not listen: %v", err)
	}

	s := grpc.NewServer()

	hello.RegisterGreeterServer(s, server{})

	log.Printf("server is listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
