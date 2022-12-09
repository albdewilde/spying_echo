package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/albdewilde/spying_echo/grpc/spyingechopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	HOST = "0.0.0.0"
	PORT = 10000
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", HOST, PORT), opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := spyingechopb.NewSpyingEchoClient(conn)

	spy(client)
}

func spy(c spyingechopb.SpyingEchoClient) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := c.Spy(ctx, new(spyingechopb.Empty))
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error when receiving from stream: %s", err.Error())
		}

		fmt.Println(msg.GetMsg())
	}
}
