package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	client := flag.Bool("client", false, "Use as a client")
	flag.Parse()

	if *client {
		conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
		if err != nil {
			fmt.Println("could not connect", err)
			return
		}
		defer conn.Close()

		myClient := NewMyTestServiceClient(conn)

		resp, err := myClient.Function1(context.Background(), &RequestMsg{Num1: 42})
		if err != nil {
			fmt.Println("whoops, error", err)
			return
		}

		fmt.Println("Server answer:", resp.GetNumres())
		return
	}

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	RegisterMyTestServiceServer(grpcServer, &myTestServerImpl{})
	grpcServer.Serve(lis)
}

type myTestServerImpl struct{}

func (s *myTestServerImpl) Function1(ctx context.Context, req *RequestMsg) (*ResponseMsg, error) {
	resp := &ResponseMsg{}
	fmt.Println("client sent", req.GetNum1())
	resp.Numres = req.GetNum1()*2
	return resp, nil
}
