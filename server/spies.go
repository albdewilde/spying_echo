package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/albdewilde/spying_echo/grpc/spyingechopb"
)

type Spies struct {
	mtx   sync.Mutex
	spies []spy
}

func newSpies() Spies {
	return Spies{
		spies: make([]spy, 0),
	}
}

func (s *Spies) Add(spy spy) {
	s.mtx.Lock()

	s.spies = append(s.spies, spy)

	s.mtx.Unlock()
}

func (s *Spies) Dispatch(req *spyingechopb.EchoRequest) {
	// Send the message to all spy
	for _, s := range s.spies {
		err := s.stream.Send(
			&spyingechopb.EchoReply{
				Msg: fmt.Sprintf(
					"%s said: %s",
					req.GetName(),
					req.GetMsg(),
				),
			},
		)

		if err != nil {
			log.Println(err)
			// We may need to delete spy when they are disconnected
		}

	}
}
