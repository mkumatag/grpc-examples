package main

import (
	"context"
	"log"
	"time"

	pb "github.com/mkumatag/grpc-examples/calculator"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Sum(ctx, &pb.CalculatorRequest{A: 10, B: 20})
	if err != nil {
		log.Fatalf("Could not calculate! error: %v", err)
	}
	log.Printf("Sum: %f", r.Result)

	r1, err := c.Div(ctx, &pb.CalculatorRequest{A: 10, B: 0})
	if err != nil {
		log.Fatalf("Could not calculate! error: %v", err)
	}
	log.Printf("Div: %f", r1.Result)
}
