package main

import (
	"context"
	"fmt"
	"github.com/shubhamdixit863/gogrpc/livescore"
	"google.golang.org/grpc"
	"log"
)

func main() {

	fmt.Println("Client Running")
	conn, err := grpc.Dial("localhost:50000", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Could Not Connect %v", err)
	}
	defer conn.Close()
	c := livescore.NewGreetServiceClient(conn)
	fmt.Printf("Created Client %f", c)
	doUnary(c)

}

func doUnary(c livescore.GreetServiceClient) {

	req := &livescore.GreetRequest{
		Greeting: &livescore.Greeting{
			FirstName: "John",
			LastName:  "Lenon",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While Calling Greet Rpc %v", err)
	}
	log.Println(res)

}
