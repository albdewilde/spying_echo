package main

import (
	"context"
	"fmt"
	"log"

	"github.com/albdewilde/spying_echo/grpc/spyingechopb"
)

// SypingEchoServer implements the  spyingechopb.SpyingechoServer interface
type SpyingechoServer struct {
	spyingechopb.UnimplementedSpyingEchoServer

	// ss regroup all spys that we'll send messages to
	ss []spy

	// ers is a channel where we put all messages to send to spys
	ers chan *spyingechopb.EchoRequest
}

func NewSpyingechoServer() *SpyingechoServer {
	ses := &SpyingechoServer{
		ss:  make([]spy, 0),
		ers: make(chan *spyingechopb.EchoRequest),
	}

	go ses.dispatch()
	return ses
}

func (ses *SpyingechoServer) Echo(ctx context.Context, echoRequest *spyingechopb.EchoRequest) (*spyingechopb.EchoReply, error) {
	ses.ers <- echoRequest

	log.Printf("%s said: %s", echoRequest.GetName(), echoRequest.GetMsg())

	return &spyingechopb.EchoReply{Msg: fmt.Sprintf("You said: %s", echoRequest.GetMsg())}, nil
}

func (ses *SpyingechoServer) Spy(_ *spyingechopb.Empty, s spyingechopb.SpyingEcho_SpyServer) error {
	// Add client to our spys list
	ses.ss = append(ses.ss, spy{stream: s})
	log.Println("A new spy is connected")

	// Infinite loop to keep stream alive
	// Sending is manage by ses.dispatch method
	for {
	}
}

func (ses *SpyingechoServer) dispatch() {
	for {
		msg := <-ses.ers

		// Send the message to all spy
		for _, s := range ses.ss {
			err := s.stream.Send(
				&spyingechopb.EchoReply{
					Msg: fmt.Sprintf(
						"%s said: %s",
						msg.GetName(),
						msg.GetMsg(),
					),
				},
			)

			if err != nil {
				log.Println(err)
				// We may need to delete spy when they are disconnected
			}

		}
	}
}
