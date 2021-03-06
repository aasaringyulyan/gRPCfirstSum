package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"testR/pkg/adder"
	"testR/pkg/api"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatalln(err)
	}
}