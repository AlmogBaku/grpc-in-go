package main

import (
	"net"
	"google.golang.org/grpc"
	"github.com/AlmogBaku/gRPC-demo/calculator"
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io"
)

func main() {
	//TCP Listener
	grpcListener, _ := net.Listen("tcp", ":5897")

	//Create a gRPC server
	baseServer := grpc.NewServer()
	//Bind implementation to the server
	calculator.RegisterMathServer(baseServer, &server{})

	fmt.Println("Server is running on " + grpcListener.Addr().String())

	//Bind gRPC server to the TCP
	baseServer.Serve(grpcListener)
}

type server struct{}

func (s *server) Sum(ctx context.Context, sumRequest *calculator.SumRequest) (*calculator.Result, error) {
	spew.Dump(sumRequest)
	return &calculator.Result{Result: sumRequest.A + sumRequest.B}, nil
}

func (s *server) SumAll(stream calculator.Math_SumAllServer) error {
	var sum int32 = 0
	for {
		numReq, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calculator.Result{Result: sum})
		}
		if err != nil {
			return err
		}
		spew.Dump(numReq)
		sum += numReq.Number
	}
}