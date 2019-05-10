// Binary lc counts the number of lines in a file and also determines the
// lenght of the shortest and the longest lines.
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/szamcsi/golang-elte-2019-public/grpc-testing/ex1/linesservice"
	"github.com/szamcsi/golang-elte-2019-public/grpc-testing/ex2/lines"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/szamcsi/golang-elte-2019-public/grpc-testing/proto/lines"
)

const port = ":54321"

func server() {
	lis, err := net.Listen("tcp", port) // HLserver
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()                                // HLserver
	pb.RegisterLinesServiceServer(s, linesservice.New()) // HLserver
	if err := s.Serve(lis); err != nil {                 // HLserver
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	flag.Parse()

	// Starting the server locally to simplify the demo:
	go server()

	// Connecting the client:
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure()) // HLclient
	if err != nil {                                               // HLclient
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewLinesServiceClient(conn) // HLclient

	for _, path := range flag.Args() {
		mmc, err := lines.Count(context.Background(), client, path)
		if err != nil {
			fmt.Printf("ERROR: %s", err)
			continue
		}
		fmt.Printf("%+v\t%s\n", mmc, path)
	}
}
