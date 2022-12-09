package main

import "github.com/albdewilde/spying_echo/grpc/spyingechopb"

// spy represent one of people that listen on the server
type spy struct {
	stream spyingechopb.SpyingEcho_SpyServer
}
