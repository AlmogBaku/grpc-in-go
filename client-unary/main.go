package main

import (
	"google.golang.org/grpc"
	"github.com/AlmogBaku/gRPC-demo/calculator"
	"context"
	"github.com/davecgh/go-spew/spew"
	"flag"
	"strconv"
)

func main() {
	flag.Parse()

	//create the connection
	conn, _ := grpc.Dial(":5897", grpc.WithInsecure())
	defer conn.Close()

	//create a client
	client := calculator.NewMathClient(conn)

	//Create a request
	a, _ := strconv.Atoi(flag.Arg(0))
	b, _ := strconv.Atoi(flag.Arg(1))

	resp, _ := client.Sum(context.Background(), &calculator.SumRequest{ int32(a), int32(b)})

	spew.Dump(resp)
}
