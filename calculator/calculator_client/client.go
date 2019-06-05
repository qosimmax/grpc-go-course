package main

import (
	"context"
	"fmt"

	"log"

	"qosimmax/grpc-go-course/calculator/calculatorpb"

	"io"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	//doUnary(c)

	doServerStreaming(c)

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Unary Sum RPC ...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 10,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Sum: %v", res.SumResult)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a PrimeDecomposition Server Streaming RPC...")

	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 100,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeNumberDecomposition RPC: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}

		fmt.Println(res.GetPrimeFactor())
	}

}
