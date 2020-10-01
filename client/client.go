package main

import (
	"context"
	calcpb "github.com/recep/grpc-golang-calculator/proto/proto-gen"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect: %s", err)
	}

	defer conn.Close()

	c := calcpb.NewCalculatorClient(conn)
	num1, num2 := argParser()

	addResp, err := c.Add(context.Background(), &calcpb.AddRequest{Number1: num1, Number2: num2})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%d + %d = %d\n", num1, num2, addResp.GetResult())

	subtractResp, err := c.Subtract(context.Background(), &calcpb.SubtractRequest{Number1: num1, Number2: num2})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%d - %d = %d\n", num1, num2, subtractResp.GetResult())

	multiplyResp, err := c.Multiply(context.Background(), &calcpb.MultiplyRequest{Number1: num1, Number2: num2})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%d * %d = %d\n", num1, num2, multiplyResp.GetResult())

	divideResp, err := c.Divide(context.Background(), &calcpb.DivideRequest{Number1: num1, Number2: num2})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%d / %d = %d\n", num1, num2, divideResp.GetResult())
}

func argParser() (int32, int32) {
	num1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	num2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	return int32(num1), int32(num2)
}
