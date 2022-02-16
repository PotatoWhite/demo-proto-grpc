package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "potato/demo-protocolbuffer-grpc/grpc/proto"
	"potato/demo-protocolbuffer-grpc/logic"
)

func main() {
	s := grpc.NewServer()
	if listen, err := net.Listen("tcp", ":8000"); err != nil {
		log.Fatalf("fail to listen: %v", err)
	} else {
		pb.RegisterSampleServiceServer(s, logic.Service{})
		err := s.Serve(listen)
		if err != nil {
			log.Fatalf("Fail to start service")
		}
	}
}
