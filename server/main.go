package main

import (
	"fmt"
	"log"
	"net"

	"github.com/albdewilde/spying_echo/grpc/spyingechopb"
	"google.golang.org/grpc"
)

const (
	HOST = "0.0.0.0"
	PORT = 10000
)

func main() {
	// Initialise listening connection
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", HOST, PORT))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Server start on " + lis.Addr().String())

	// Create and start the gRPC server
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	spyingechopb.RegisterSpyingEchoServer(grpcServer, NewSpyingechoServer())
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
