package main

import (
	"context"
	calcpb "github.com/recep/grpc-golang-calculator/proto/proto-gen"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	protocol = "tcp"
	port     = ":50051"
)

type srv struct{}

func main() {
	lis, err := net.Listen(protocol, port)
	if err != nil {
		log.Fatalf("failed to listen %s\n", err)
	}
	s := grpc.NewServer()
	calcpb.RegisterCalculatorServer(s, &srv{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %s\n", err)
	}
}

func (s *srv) Add(ctx context.Context, in *calcpb.AddRequest) (*calcpb.AddResponse, error) {
	return &calcpb.AddResponse{Result: in.GetNumber1() + in.GetNumber2()}, nil
}

func (s *srv) Subtract(ctx context.Context, in *calcpb.SubtractRequest) (*calcpb.SubtractResponse, error) {
	return &calcpb.SubtractResponse{Result: in.GetNumber1() - in.GetNumber2()}, nil
}

func (s *srv) Multiply(ctx context.Context, in *calcpb.MultiplyRequest) (*calcpb.MultiplyResponse, error) {
	return &calcpb.MultiplyResponse{Result: in.GetNumber1() * in.GetNumber2()}, nil
}

func (s *srv) Divide(ctx context.Context, in *calcpb.DivideRequest) (*calcpb.DivideResponse, error) {
	return &calcpb.DivideResponse{Result: in.GetNumber1() / in.GetNumber2()}, nil
}
