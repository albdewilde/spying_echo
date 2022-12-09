package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/albdewilde/spying_echo/grpc/spyingechopb"
)

// SypingEchoServer implements the  spyingechopb.SpyingechoServer interface
type SpyingechoServer struct {
	spyingechopb.UnimplementedSpyingEchoServer

	// spies regroup all spies that we'll send messages to
	spies Spies

	// ers is a channel where we put all messages to send to spys
	ers chan *spyingechopb.EchoRequest
}

func NewSpyingechoServer() *SpyingechoServer {
	server := &SpyingechoServer{
		spies: newSpies(),
		ers:   make(chan *spyingechopb.EchoRequest),
	}

	go server.dispatch()
	return server
}

func (server *SpyingechoServer) Echo(ctx context.Context, echoRequest *spyingechopb.EchoRequest) (*spyingechopb.EchoReply, error) {
	server.ers <- echoRequest

	log.Printf("%s said: %s", echoRequest.GetName(), echoRequest.GetMsg())

	return &spyingechopb.EchoReply{Msg: fmt.Sprintf("You said: %s", echoRequest.GetMsg())}, nil
}

func (server *SpyingechoServer) Spy(_ *spyingechopb.Empty, stream spyingechopb.SpyingEcho_SpyServer) error {
	// Add client to our spys list
	server.spies.Add(spy{stream: stream})
	log.Println("A new spy is connected")

	// Infinite loop to keep stream alive
	// Sending is manage by ses.dispatch method
	for {
		time.Sleep(time.Minute * 5)
	}
}

func (server *SpyingechoServer) dispatch() {
	for msg := range server.ers {
		server.spies.Dispatch(msg)
	}
}
