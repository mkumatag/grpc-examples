package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/mkumatag/grpc-examples/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type calculator struct{}

func (c *calculator) Sum(ctx context.Context, cr *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Ops: Add, operands are: %f and %f", cr.A, cr.B)
	return &pb.CalculatorResponse{Result: cr.A + cr.B}, nil
}

func (c *calculator) Div(ctx context.Context, cr *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Ops: Div, operands are: %f and %f", cr.A, cr.B)
	if cr.B == 0 {
		return nil, errors.New("Divide by zero")
	}
	return &pb.CalculatorResponse{Result: cr.A / cr.B}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &calculator{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
