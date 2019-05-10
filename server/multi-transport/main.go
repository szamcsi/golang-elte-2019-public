package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

func main() {
	client := flag.Bool("client", false, "Use as a client")
	flag.Parse()

	if *client {
		clientMode()
		return
	}

	exitCh := make(chan struct{})

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "root path")
	})

	mux.HandleFunc("/kill", func(w http.ResponseWriter, r *http.Request) {
		panic("whoops")
	})

	mux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		close(exitCh)
		//exitCh <- struct{}{}
	})
	srv := &http.Server{Addr: ":8081", Handler: mux	}
	go func() {
		srv.ListenAndServe()
		fmt.Println("close exit channel")
		close(exitCh)
		//exitCh <- struct{}{}
	}()

	grpcListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	RegisterMyTestServiceServer(grpcServer, &myTestServerImpl{})
	go func() {
		grpcServer.Serve(grpcListener)
		close(exitCh)
		//exitCh <- struct{}{}
	}()


	// Wait for the servers
	fmt.Println("wait on exit channel")
	<-exitCh
	fmt.Println("exit the program")
}

type myTestServerImpl struct{}

func (s *myTestServerImpl) Function1(ctx context.Context, req *RequestMsg) (*ResponseMsg, error) {
	resp := &ResponseMsg{}
	fmt.Println("client sent", req.GetNum1())
	resp.Numres = req.GetNum1()*2
	return resp, nil
}


func clientMode() {
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