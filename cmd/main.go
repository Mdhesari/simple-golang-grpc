package main

import (
	"context"
	"flag"
	"log"
	"mdhesari/simple-golang-grpc/contract/hello"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var defaultName = "Greeting"

var (
	addr = flag.String("addr", "localhost:5001", "the address to connect")
	name = flag.String("name", defaultName, "name to greet")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error connecting: %v", err)
	}
	defer conn.Close()

	c := hello.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &hello.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not context: %v", err)
	}

	log.Printf("here is your message: %v",r.GetMessage())
}
