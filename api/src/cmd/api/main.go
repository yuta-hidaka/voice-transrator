package main

import (
	"log"
	"net"
	pb "voice-translator/internal/proto"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:8080"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on: %v", addr) 

	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
