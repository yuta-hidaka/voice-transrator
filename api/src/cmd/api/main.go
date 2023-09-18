package main

import (
	"log"
	"net"
	pb "voice-translator/internal/proto"
	"voice-translator/internal/server/translator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:8080"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTranslatorServiceServer(s, &translator.Server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("listen on: %v", addr)
}
