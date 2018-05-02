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

	//Stream requests
	stream, _ := client.SumAll(context.Background())
	for _, num := range flag.Args() {
		n, _ := strconv.Atoi(num)
		stream.Send(&calculator.SumNumberRequest{Number: int32(n)})
	}

	resp, _ := stream.CloseAndRecv()
	spew.Dump(resp)
}
