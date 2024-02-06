package main

import (
	"context"
	"fmt"
	"grpc-exam/calculatorpb"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.CalculatorServiceServer
}

func (sv *server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Println("Calculating sum...")
	res := &calculatorpb.SumResponse{
		Result: req.GetNumber1() + req.GetNumber2(),
	}
	return res, nil
}

func (sv *server) Multi(ctx context.Context, req *calculatorpb.MultiRequest) (*calculatorpb.MultiResponse, error) {
	log.Println("Calculating Multi...")
	time.Sleep(time.Second*5)
	res := &calculatorpb.MultiResponse{
		Result: req.GetNumber1() * req.GetNumber2(),
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	fmt.Println("Server connecting...")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalln("error while serve", err)

	}
}
