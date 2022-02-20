package main

import (
	"context"
	"fmt"
	"github.com/shubhamdixit863/gogrpc/livescore"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	livescore.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *livescore.GreetRequest) (*livescore.GreetResponse, error) {
	fmt.Println("GReet Function Invoked")
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello" + firstName
	res := &livescore.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50000")

	if err != nil {
		log.Fatal("Failed To Listen %v", err)
	}

	s := grpc.NewServer()
	livescore.RegisterGreetServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
