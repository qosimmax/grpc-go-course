package main

import (
	"context"
	"fmt"

	"log"

	"qosimmax/grpc-go-course/calculator/calculatorpb"

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

	doUnary(c)

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
