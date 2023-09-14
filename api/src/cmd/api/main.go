package main

import (
	"log"
	"net"
	pb "voice-translator/internal/proto"
	"voice-translator/internal/server"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:8080"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("listen on1: %v", addr)
	s := grpc.NewServer()
	log.Printf("listen on2: %v", addr)
	pb.RegisterGreetServiceServer(s, &server.Server{})
	log.Printf("listen on3: %v", addr)
	pb.RegisterCalcServiceServer(s, &server.Server{})
	log.Printf("listen on4: %v", addr)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
