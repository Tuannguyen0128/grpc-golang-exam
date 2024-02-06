package main

import (
	"context"
	"fmt"
	"grpc-exam/calculatorpb"
	"log"
	"time"

	"google.golang.org/grpc"
)
func calSum(req *calculatorpb.SumRequest,c calculatorpb.CalculatorServiceClient){
	res,err:=c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalln("err when call sum",err)
	}
	log.Println("calSum response",res.GetResult())
}
func calMulti(req *calculatorpb.MultiRequest,c calculatorpb.CalculatorServiceClient){
	res,err:=c.Multi(context.Background(), req)
	if err != nil {
		log.Fatalln("err when call multi",err)
	}
	log.Println("calMulti response",res.GetResult())
}
func main() {
	cc, err := grpc.Dial("localhost:3000",grpc.WithInsecure())
	if err != nil {
		log.Fatalln("err when dial",err)
	}
	defer cc.Close()
	client:=calculatorpb.NewCalculatorServiceClient(cc)
	//log.Printf("Service client %f",client)
	start:=time.Now()
	calSum(&calculatorpb.SumRequest{Number1: 100,Number2: 233},client)
	start2:=time.Now()
	fmt.Println("execute time:",time.Since(start))
	calMulti(&calculatorpb.MultiRequest{Number1: 100,Number2: 233},client)
	fmt.Println("execute time:",time.Since(start2))

}