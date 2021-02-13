package main

import (
	"google.golang.org/grpc"
	"grpc/pkg/adder"
	"grpc/pkg/api"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.AdderServer{}
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
