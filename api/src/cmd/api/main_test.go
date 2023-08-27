package main_test

import (
	"log"
	"testing"

	"google.golang.org/grpc"
)

var addr string = "localhost:8080"

func TestMain(m *testing.M) {
	conn, err := grpc.Dial(addr)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
}
